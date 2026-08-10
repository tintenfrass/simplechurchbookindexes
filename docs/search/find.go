package search

import (
	"fmt"
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

func FindMarriage(input string, min, max int, churches map[string]bool, algo int) ([]Result, string) {
	var debug string
	defer func() {
		if r := recover(); r != nil {
			debug = fmt.Sprintf(" Es ist ein Fehler aufgetreten. Seite muss neu geladen werden! (F5)")
		}
	}()
	searchParts := cleanInput(input)

	soundex := algo == 1
	searcher := getSearcher(algo)

	// neue Variante für Double Search
	if len(searchParts) == 2 {
		return findDouble(searchParts, min, max, churches, searcher, soundex)
	}

	// Fallback to Single Search the old way
	results := make(map[int][]marriageEntry)
	jaro := algo < 2

	//Algo
	jaroTreshold := JaroTreshold
	if soundex {
		//für Soundex brauchen wir einen geringen Grenzwert
		jaroTreshold = JaroTresholdSoundex
	}

	// Full Data
	for church, sourceMarriages := range Data.Marriages {
		//Prüfen, ob wir in dieser Quelle suchen wollen
		_, exists := churches[church]
		if !exists || !churches[church] {
			continue
		}

		sm := sourceMarriages //Prevent Bug
		for _, entry := range sm.Data {
			nameV, _ := Data.NamesV[entry.V]
			nameN, _ := Data.NamesN[entry.N]
			//Prüfen, ob wir in dieser Zeit suchen wollen
			if (entry.Y < min || entry.Y > max) && entry.Y != 0 {
				continue
			}

			distance := 0
			// Vorname und Nachname als kombinierter Input
			searchName := searchParts[0]
			//Jaro Vorfilterung
			if jaro {
				if matchr.Jaro(searchName, nameV+" "+nameN) < jaroTreshold {
					continue
				}
			}
			if soundex {
				searchName = gophonetics.NewPhoneticCode(searchName)
				nameV = gophonetics.NewPhoneticCode(nameV)
				nameN = gophonetics.NewPhoneticCode(nameN)
			}

			//Simple Search
			distance = searcher.search(searchName, nameV+" "+nameN)

			if distance > MaxDistance {
				continue
			}

			if len(results[distance]) <= MaxResultsDis {
				results[distance] = append(results[distance], entry)
			}
		}
	}

	return mapResults(results), debug
}

func cleanInput(search string) []string {
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

	return searchParts
}

func mapResults(searchResults map[int][]marriageEntry) []Result {
	var mappedResults []Result
	count := 0
	for i := 0; i < MaxDistance+1; i++ {
		if count > 50 {
			break
		}

		//too much
		if len(searchResults[i]) > MaxResultsDis {
			mappedResults = append(mappedResults, Result{
				Year:   0,
				Line:   "Zu viele Ergebnisse zum Anzeigen (>" + strconv.Itoa(MaxResultsDis) + ")",
				Source: "",
				Dis:    i,
				Link:   "",
				Page:   0,
			})
			return mappedResults
		}

		for _, match := range searchResults[i] {
			pageId := 0
			if Data.Offset[match.S] > 0 {
				pageId = int(Data.Offset[match.S] + uint32(match.P))
			} else {
				pageId = match.P
			}
			line := match.L
			if len(line) == 0 {
				line = Data.NamesV[match.V] + " " + Data.NamesN[match.N]
			}
			mappedResults = append(mappedResults, Result{
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

	return mappedResults
}
