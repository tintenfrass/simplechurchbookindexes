package search

import (
	"math"

	gophonetics "gopkg.in/Regis24GmbH/go-phonetics.v3"
)

func findDouble(
	searchParts []string,
	min, max int,
	places map[string]struct{},
	exact bool,
	searcher Searcher,
	soundex bool,
) ([]Result, string) {
	var debug string

	if soundex {
		searchParts[0] = gophonetics.NewPhoneticCode(searchParts[0])
		searchParts[1] = gophonetics.NewPhoneticCode(searchParts[1])
	}

	// create search index
	namesPlaceV := make(map[int]struct{})
	namesPlaceN := make(map[int]struct{})
	for place, _ := range Data.Marriages {
		if !isPlaceValid(places, exact, place) {
			continue
		}
		for nameIdV, _ := range Data.PlaceV[place] {
			if _, ok := namesPlaceV[nameIdV]; !ok {
				namesPlaceV[nameIdV] = struct{}{}
			}
		}
		for nameIdN, _ := range Data.PlaceN[place] {
			if _, ok := namesPlaceN[nameIdN]; !ok {
				namesPlaceN[nameIdN] = struct{}{}
			}
		}
	}

	namesDecadeV := make(map[int]struct{})
	namesDecadeN := make(map[int]struct{})
	for decade := min / 10; decade <= max/10; decade++ {
		for nameIdV := range Data.DecadeV[decade] {
			namesDecadeV[nameIdV] = struct{}{}
		}
		for nameIdN := range Data.DecadeN[decade] {
			namesDecadeN[nameIdN] = struct{}{}
		}
	}

	searchTargetsV := make(map[int]int)
	for nameIndexV, val := range Data.NamesV {
		if _, ok := namesPlaceV[nameIndexV]; !ok {
			continue
		}
		if _, ok := namesDecadeV[nameIndexV]; !ok {
			continue
		}
		if soundex {
			val = gophonetics.NewPhoneticCode(val)
		}
		dis := getDistanceV(searchParts[0], val, searcher)
		if dis > MaxDistance {
			continue
		}
		searchTargetsV[nameIndexV] = dis
	}

	searchTargetsN := make(map[int]int)
	for nameIndexN, val := range Data.NamesN {
		if _, ok := namesPlaceN[nameIndexN]; !ok {
			continue
		}
		if _, ok := namesDecadeN[nameIndexN]; !ok {
			continue
		}
		if soundex {
			val = gophonetics.NewPhoneticCode(val)
		}
		dis := getDistanceN(searchParts[1], val, searcher)
		if dis > MaxDistance {
			continue
		}
		searchTargetsN[nameIndexN] = dis
	}

	searchResults := make(map[int][]marriageEntry)

	// Full Data
	for place, sourceMarriages := range Data.Marriages {
		if !isPlaceValid(places, exact, place) {
			continue
		}

		sm := sourceMarriages // reasign to prevent a stupid bug
		for _, entry := range sm.Data {
			//Prüfen, ob wir in dieser Zeit suchen wollen
			if (entry.Y < min || entry.Y > max) && entry.Y != 0 {
				continue
			}

			disV, ok := searchTargetsV[entry.V]
			if !ok {
				continue
			}
			disN, ok := searchTargetsN[entry.N]
			if !ok {
				continue
			}
			//1/3 der Differenz abziehen => damit werden Matches leicht bevorteilt, wo ein Part sehr gut matched
			distance := disV + disN - int(math.Round(0.3*math.Abs(float64(disV)-float64(disN))))
			if distance > MaxDistance {
				continue
			}
			if distance < 0 {
				distance = 0
			}

			if len(searchResults[distance]) <= MaxResultsDis {
				searchResults[distance] = append(searchResults[distance], entry)
			}
		}
	}

	return mapResults(searchResults), debug
}

func getDistanceV(input, target string, searcher Searcher) int {
	distance := 0
	if input == "?" {
		return distance
	}

	return searcher.search(input, target)
}

func getDistanceN(input, target string, searcher Searcher) int {
	distance := 0
	if input == "?" {
		return distance
	}

	dis := searcher.search(input, target)
	//Bonuspunkt wenn der erste Buchstabe passt
	if len(target) > 0 && input[0:1] == target[0:1] {
		dis--
	}

	return dis
}
