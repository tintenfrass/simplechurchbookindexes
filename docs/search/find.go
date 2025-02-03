package search

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/antzucaro/matchr"
	gophonetics "gopkg.in/Regis24GmbH/go-phonetics.v3"
)

type Result struct {
	Year   int
	Line   string
	Source string
	Dis    int
	Link   string
	Page   int
}

const MaxDistance = 7
const MaxResultsDis = 1000

func FindMarriage(search string, min, max int, churches map[string]bool, algo int) (resultList []Result, debug string) {
	defer func() {
		if r := recover(); r != nil {
			debug = fmt.Sprintf(" Es ist ein Fehler aufgetreten. Seite muss neu geladen werden! (F5)")
		}
	}()

	search = strings.TrimSpace(search)
	//Alle bis auf das letzte Leerzeichen ersetzen, damit Vornamen zusammengehangen werden
	for {
		if strings.Count(search, " ") < 2 {
			break
		}
		search = strings.Replace(search, " ", "-", 1)
	}

	//* oder ? als Platzhalter
	search = strings.Replace(search, "*", "?", -1)

	//Nachname abspalten
	searchParts := strings.Split(search, " ")
	searchParts[0] = strings.Replace(searchParts[0], "-", " ", -1)
	if len(searchParts) > 1 {
		searchParts[1] = strings.Replace(searchParts[1], "-", " ", -1)
	}

	results := make(map[int][]marriageEntry)

	//Algo
	searcher := getSearcher(algo)
	jaroTreshold := JaroTreshold
	if algo == 1 {
		//für Soundex brauchen wir einen geringen Grenzwert
		jaroTreshold = JaroTresholdSoundex
	}

	for church, sourceMarriages := range Data.Marriages {
		//Prüfen, ob wir in dieser Quelle suchen wollen
		_, exists := churches[church]
		if !exists || !churches[church] {
			continue
		}

		sm := sourceMarriages //Prevent Bug
		for _, entry := range sm.Data {
			//Prüfen, ob wir in dieser Zeit suchen wollen
			if (entry.Y < min || entry.Y > max) && entry.Y != 0 {
				continue
			}

			distance := 0
			if len(searchParts) == 1 {
				searchName := search
				//Jaro Vorfilterung, nur bei Algo = 0
				if algo == 0 {
					if matchr.Jaro(search, entry.V+" "+entry.N) < jaroTreshold {
						continue
					}
				}
				//Soundex bei Algo = 1
				if algo == 1 {
					searchName = gophonetics.NewPhoneticCode(search)
					entry.V = gophonetics.NewPhoneticCode(entry.V)
					entry.N = gophonetics.NewPhoneticCode(entry.N)
				}

				//Simple Search
				distance = searcher.search(searchName, entry.V+" "+entry.N)
			} else {
				if algo < 2 {
					//Jaro Vorfilterung, nur bei Algo = 0 oder 1
					jaroGroom := 1.0
					jaroGroomFn := 1.0
					if searchParts[0] != "?" {
						jaroGroom = matchr.Jaro(searchParts[0], entry.V)
						if jaroGroom < jaroTreshold {
							continue
						}
					}
					if searchParts[1] != "?" {
						jaroGroomFn = matchr.Jaro(searchParts[1], entry.N)
						if jaroGroomFn < jaroTreshold {
							continue
						}
					}
					//Die Kombination aus Vor- und Nachname muss ca. 50% übereinstimmen, um überhaupt in die nähere Betrachtung zu gelangen
					if jaroGroom+jaroGroomFn < 1 {
						continue
					}
				}

				//Double Search
				distanceGroom := 0
				distanceGroomFn := 0

				s0 := searchParts[0]
				s1 := searchParts[1]
				//Soundex
				if algo == 1 {
					s0 = gophonetics.NewPhoneticCode(searchParts[0])
					s1 = gophonetics.NewPhoneticCode(searchParts[1])
					entry.V = gophonetics.NewPhoneticCode(entry.V)
					entry.N = gophonetics.NewPhoneticCode(entry.N)
				}

				//? matched auf alles
				if searchParts[0] != "?" {
					distanceGroom = searcher.search(s0, entry.V)
				}
				if searchParts[1] != "?" {
					distanceGroomFn = searcher.search(s1, entry.N)
					//Bonuspunkt wenn der erste Buchstabe passt
					if len(entry.N) > 0 && searchParts[1][0:1] == entry.N[0:1] {
						distanceGroomFn--
					}
				}

				//1/3 der Differenz abziehen => damit werden Matches leicht bevorteilt, wo ein Part sehr gut matched
				distance = distanceGroom + distanceGroomFn - int(math.Round(0.3*math.Abs(float64(distanceGroom)-float64(distanceGroomFn))))
				if distance < 0 {
					distance = 0
				}
			}

			if distance > MaxDistance {
				continue
			}

			if len(results[distance]) <= MaxResultsDis {
				results[distance] = append(results[distance], entry)
			}
		}
	}

	count := 0
	for i := 0; i < MaxDistance+1; i++ {
		if count > 50 {
			break
		}

		//too much
		if len(results[i]) > MaxResultsDis {
			resultList = append(resultList, Result{
				Year:   0,
				Line:   "Zu viele Ergebnisse zum Anzeigen (>" + strconv.Itoa(MaxResultsDis) + ")",
				Source: "",
				Dis:    i,
				Link:   "",
				Page:   0,
			})
			return
		}

		for _, match := range results[i] {
			pageId := 0
			if Data.Offset[match.S] > 0 {
				pageId = int(Data.Offset[match.S] + uint32(match.P))
			}
			line := match.L
			if len(line) == 0 {
				line = match.V + " " + match.N
			}
			resultList = append(resultList, Result{
				Year:   match.Y,
				Line:   line,
				Source: Data.Sources[match.S],
				Dis:    i,
				Link:   Data.Links[match.S],
				Page:   pageId,
			})
			count++
		}
	}

	return
}
