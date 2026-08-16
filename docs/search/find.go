package search

import (
	"strconv"
	"strings"
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

func FindMarriage(inputV, inputN string, min, max int, places map[string]struct{}, algo int, exact bool) ([]Result, string) {
	searchParts := cleanInput(inputV, inputN)

	soundex := algo == 1
	searcher := getSearcher(algo)

	return search(searchParts, min, max, places, exact, searcher, soundex)
}

func cleanInput(inputV, inputN string) []string {
	inputV = strings.Replace(inputV, "*", "", -1)
	inputV = strings.Replace(inputV, "?", "", -1)
	inputV = strings.Replace(inputV, "-", "", -1)
	inputV = strings.Replace(inputV, "_", "", -1)
	inputV = strings.TrimSpace(inputV)

	inputN = strings.Replace(inputN, "*", "", -1)
	inputN = strings.Replace(inputN, "?", "", -1)
	inputN = strings.Replace(inputN, "_", "", -1)
	inputN = strings.TrimSpace(inputN)

	return []string{inputV, inputN}
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

func isPlaceValid(places map[string]struct{}, exact bool, place string) bool {
	//Prüfen, ob wir in dieser Quelle suchen wollen
	if exact {
		if _, exists := places[place]; !exists {
			return false
		}
		return true
	}

	for pl, _ := range places {
		if strings.HasPrefix(place, strings.Replace(pl, "*", "", 1)) {
			return true
		}
	}
	return false
}
