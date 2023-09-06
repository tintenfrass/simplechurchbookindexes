package search

import (
	"fmt"
	"math"
	"path"
	"strings"

	"github.com/antzucaro/matchr"
)

func FindMarriage(search string, min, max int, churches map[string]bool) (output []string) {
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

	for church, sourceMarriages := range Data.Marriages {
		//Prüfen, ob wir in dieser Quelle suchen wollen
		_, exists := churches[church]
		if exists && !churches[church] {
			continue
		}

		sm := sourceMarriages //Prevent Bug
		for _, entry := range sm.Data {
			//Prüfen, ob wir in dieser Zeit suchen wollen
			if entry.Y != 0 && (entry.Y < min || entry.Y > max) {
				continue
			}
			distance := 0
			if len(searchParts) == 1 {
				//Simple Search
				distance = matchr.DamerauLevenshtein(search, entry.V+" "+entry.N)
			} else {
				//Double Search
				distanceGroom := 0
				distanceGroomFn := 0
				//? matched auf alles
				if searchParts[0] != "?" {
					distanceGroom = matchr.DamerauLevenshtein(searchParts[0], entry.V)
				}
				if searchParts[1] != "?" {
					distanceGroomFn = matchr.DamerauLevenshtein(searchParts[1], entry.N)
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
			output = append(output, fmt.Sprintf("%d %s/%s/%d", match.Y, match.L, path.Base(Data.Sources[match.S]), i))
			count++
		}
	}

	return
}
