package search

import (
	"fmt"
	"testing"
	"time"

	"github.com/alextanhongpin/stringdist"
	"github.com/antzucaro/matchr"
	"github.com/hbollon/go-edlib"
	tdl "github.com/lmas/Damerau-Levenshtein"
)

func TestFind(t *testing.T) {
	s := time.Now()
	for i := 0; i < 100000; i++ {
		matchr.Levenshtein("Johann Gottfried", "Joseph Wilhelm")
	}
	fmt.Println(time.Since(s))

	for i := 0; i < 100000; i++ {
		matchr.OSA("Johann Gottfried", "Joseph Wilhelm")
	}
	fmt.Println(time.Since(s))

	s = time.Now()
	for i := 0; i < 100000; i++ {
		matchr.DamerauLevenshtein("Johann Gottfried", "Joseph Wilhelm")
	}
	fmt.Println(time.Since(s))

	s = time.Now()
	for i := 0; i < 100000; i++ {
		matchr.JaroWinkler("Johann Gottfried", "Joseph Wilhelm", true)
	}
	fmt.Println(time.Since(s))

	s = time.Now()
	for i := 0; i < 100000; i++ {
		matchr.JaroWinkler("Johann Gottfried", "Joseph Wilhelm", false)
	}
	fmt.Println(time.Since(s))

	s = time.Now()
	for i := 0; i < 100000; i++ {
		matchr.Jaro("Johann Gottfried", "Joseph Wilhelm")
	}
	fmt.Println(time.Since(s))

	s = time.Now()
	for i := 0; i < 100000; i++ {
		tdl.Distance("Johann Gottfried", "Joseph Wilhelm")
	}
	fmt.Println(time.Since(s))

	s = time.Now()
	for i := 0; i < 100000; i++ {
		edlib.LevenshteinDistance("Johann Gottfried", "Joseph Wilhelm")
	}
	fmt.Println(time.Since(s))

	s = time.Now()
	for i := 0; i < 100000; i++ {
		edlib.DamerauLevenshteinDistance("Johann Gottfried", "Joseph Wilhelm")
	}
	fmt.Println(time.Since(s))

	s = time.Now()
	for i := 0; i < 100000; i++ {
		edlib.OSADamerauLevenshteinDistance("Johann Gottfried", "Joseph Wilhelm")
	}
	fmt.Println(time.Since(s))

	s = time.Now()
	for i := 0; i < 100000; i++ {
		edlib.JaroWinklerSimilarity("Johann Gottfried", "Joseph Wilhelm")
	}
	fmt.Println(time.Since(s))

	s = time.Now()
	for i := 0; i < 100000; i++ {
		edlib.JaroSimilarity("Johann Gottfried", "Joseph Wilhelm")
	}
	fmt.Println(time.Since(s))

	s = time.Now()
	for i := 0; i < 100000; i++ {
		stringdist.NewDamerauLevenshtein(100).Calculate("Johann Gottfried", "Joseph Wilhelm")
	}
	fmt.Println(time.Since(s))

	s = time.Now()
	for i := 0; i < 100000; i++ {
		stringdist.NewLevenshtein(100).Calculate("Johann Gottfried", "Joseph Wilhelm")
	}
	fmt.Println(time.Since(s))

	s = time.Now()
	for i := 0; i < 100000; i++ {
		stringdist.JaroWinkler("Johann Gottfried", "Joseph Wilhelm")
	}
	fmt.Println(time.Since(s))

	s = time.Now()
	for i := 0; i < 100000; i++ {
		stringdist.Jaro("Johann Gottfried", "Joseph Wilhelm")
	}
	fmt.Println(time.Since(s))
}
