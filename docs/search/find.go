package search

import (
	"fmt"
	"math"
	"path"
	"strings"

	"github.com/antzucaro/matchr"
)

func FindMarriage(search string, min, max int, churches map[string]bool, algo int) (output []string) {
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

	results := make(map[int][]marriageEntry)

	//Algo
	searcher := getSearcher(algo)

	for church, sourceMarriages := range Data.Marriages {
		//Prüfen, ob wir in dieser Quelle suchen wollen
		_, exists := churches[church]
		if exists && !churches[church] {
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
				//Jaro Vorfilterung, nur bei Algo = 0
				if algo == 0 {
					if matchr.Jaro(search, entry.V+" "+entry.N) < JaroTreshold {
						continue
					}
				}

				//Simple Search
				distance = searcher.search(search, entry.V+" "+entry.N)
			} else {
				if algo == 0 {
					//Jaro Vorfilterung, nur bei Algo = 0
					jaroGroom := 1.0
					jaroGroomFn := 1.0
					if searchParts[0] != "?" {
						jaroGroom = matchr.Jaro(searchParts[0], entry.V)
						if jaroGroom < JaroTreshold {
							continue
						}
					}
					if searchParts[1] != "?" {
						jaroGroomFn = matchr.Jaro(searchParts[1], entry.N)
						if jaroGroomFn < JaroTreshold {
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
				//? matched auf alles
				if searchParts[0] != "?" {
					distanceGroom = searcher.search(searchParts[0], entry.V)
				}
				if searchParts[1] != "?" {
					distanceGroomFn = searcher.search(searchParts[1], entry.N)
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

			results[distance] = append(results[distance], entry)
		}
	}

	count := 0
	for i := 0; i < 32; i++ {
		if count > 50 {
			break
		}
		for _, match := range results[i] {
			pageId := uint32(0)
			if Data.Offset[match.S] > 0 {
				pageId = Data.Offset[match.S] + uint32(match.P)
			}
			output = append(output, fmt.Sprintf("%d %s#%s#%d#%s#%d", match.Y, match.L, path.Base(Data.Sources[match.S]), i, Data.Links[match.S], pageId))
			count++
		}
	}

	return
}
