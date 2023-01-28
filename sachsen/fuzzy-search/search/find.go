package search

import (
	"fmt"
	"indexfuzzysearch/config"
	"math"
	"strings"

	"github.com/antzucaro/matchr"
)

func FindMarriage(search string) (output string) {
	searchParts := strings.Split(search, " ")
	results := make(map[int][]marriageEntry)
	for source, sourceMarriages := range data.marriages {
		//Prüfen, ob wir in dieser Quelle suchen wollen
		_, exists := config.Config.Churches[source]
		if !exists && !config.Config.Churches["Andere"] {
			continue
		}
		if exists && !config.Config.Churches[source] {
			continue
		}
		for _, entry := range sourceMarriages.data {
			//Prüfen, ob wir in dieser Zeit suchen wollen
			if entry.year != 0 && (entry.year < config.Config.Year["min"] || entry.year > config.Config.Year["max"]) {
				continue
			}
			distance := 0
			if len(searchParts) == 1 {
				//Simple Search
				distance = matchr.DamerauLevenshtein(search, replaceFlippedSpecial(entry.groom)+" "+replaceFlippedSpecial(entry.groomFN))
			} else {
				//Double Search
				distanceGroom := 0
				distanceGroomFn := 0
				//? matched auf alles
				if searchParts[0] != "?" {
					distanceGroom = matchr.DamerauLevenshtein(searchParts[0], replaceFlippedSpecial(entry.groom))
				}
				if searchParts[1] != "?" {
					replaced := replaceFlippedSpecial(entry.groomFN)
					distanceGroomFn = matchr.DamerauLevenshtein(searchParts[1], replaced)
					//Bonuspunkt wenn der erste Buchstabe passt
					if len(replaced) > 0 && searchParts[1][0:1] == replaced[0:1] {
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
			output += fmt.Sprintf("%d %s \t%s \t(Abweichung: ~%d)\r\n", match.year, replaceSpecial(match.groom+" "+match.groomFN), match.source, i)
			count++
		}
	}

	return
}

//Special Cases
func replaceFlippedSpecial(input string) string {
	for normal, special := range specialCases {
		input = strings.Replace(input, special, normal, -1)
		input = strings.Replace(input, "?", "", -1)
		input = strings.Replace(input, "-", "", -1)
	}
	return input
}

//Special Cases
func replaceSpecial(input string) string {
	for normal, special := range specialCases {
		input = strings.Replace(input, special, normal, -1)
	}
	input = strings.Replace(input, "ÿ", "y", -1)
	input = strings.Replace(input, "æ", "ae", -1)
	input = strings.Replace(input, "ë", "e", -1)
	input = strings.Replace(input, "ï", "i", -1)
	return input
}
