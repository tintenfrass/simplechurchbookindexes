package search

import (
	"fmt"
	"testing"
	"time"

	"github.com/alextanhongpin/stringdist"
	"github.com/antzucaro/matchr"
	"github.com/eskriett/strmet"
	"github.com/hbollon/go-edlib"
	"github.com/jamesturk/go-jellyfish"
	fast "github.com/ka-weihe/fast-levenshtein"
	tdl "github.com/lmas/Damerau-Levenshtein"
	"github.com/masatana/go-textdistance"
)

const s1 = "Johann Georg Friedrich"
const s2 = "Johann Gotthelf Benjamin"
const rotations = 100000

func TestDamerauLevenshtein(t *testing.T) {
	s := time.Now()
	d := 0

	s = time.Now()
	for i := 0; i < rotations; i++ {
		d = matchr.DamerauLevenshtein(s1, s2)
	}
	fmt.Println("matchr.DamerauLevenshtein", time.Since(s), d)

	s = time.Now()
	for i := 0; i < rotations; i++ {
		d = tdl.Distance(s1, s2)
	}
	fmt.Println("tdl.Distance", time.Since(s), d)

	s = time.Now()
	true := tdl.New(100)
	for i := 0; i < rotations; i++ {
		d = true.Distance(s1, s2)
	}
	fmt.Println("tdl.Distance", time.Since(s), d)

	s = time.Now()
	for i := 0; i < rotations; i++ {
		d = edlib.DamerauLevenshteinDistance(s1, s2)
	}
	fmt.Println("edlib.DamerauLevenshteinDistance", time.Since(s), d)

	s = time.Now()
	for i := 0; i < rotations; i++ {
		d = jellyfish.DamerauLevenshtein(s1, s2)
	}
	fmt.Println("jellyfish.DamerauLevenshtein", time.Since(s), d)
}

func TestLevenshtein(t *testing.T) {
	s := time.Now()
	d := 0

	s = time.Now()
	for i := 0; i < rotations; i++ {
		d = matchr.Levenshtein(s1, s2)
	}
	fmt.Println("matchr.Levenshtein", time.Since(s), d)

	s = time.Now()
	for i := 0; i < rotations; i++ {
		d = edlib.LevenshteinDistance(s1, s2)
	}
	fmt.Println("edlib.LevenshteinDistance", time.Since(s), d)

	s = time.Now()
	dl := stringdist.NewDamerauLevenshtein(100) //algorithm is wrong, this is Levenshtein
	for i := 0; i < rotations; i++ {
		d = dl.Calculate(s1, s2)
	}
	fmt.Println("stringdist.NewDamerauLevenshtein", time.Since(s), d)

	s = time.Now()
	l := stringdist.NewLevenshtein(100)
	for i := 0; i < rotations; i++ {
		d = l.Calculate(s1, s2)
	}
	fmt.Println("stringdist.NewLevenshtein", time.Since(s), d)

	s = time.Now()
	for i := 0; i < rotations; i++ {
		d = strmet.DamerauLevenshtein(s1, s2, 100) //algorithm is wrong, this is Levenshtein
	}
	fmt.Println("strmet.DamerauLevenshtein", time.Since(s), d)

	s = time.Now()
	for i := 0; i < rotations; i++ {
		d = strmet.Levenshtein(s1, s2, 100)
	}
	fmt.Println("strmet.Levenshtein", time.Since(s), d)

	s = time.Now()
	for i := 0; i < rotations; i++ {
		d = textdistance.DamerauLevenshteinDistance(s1, s2) //algorithm is wrong, this is Levenshtein
	}
	fmt.Println("textdistance.DamerauLevenshteinDistance", time.Since(s), d)

	s = time.Now()
	for i := 0; i < rotations; i++ {
		d = jellyfish.Levenshtein(s1, s2)
	}
	fmt.Println("jellyfish.Levenshtein", time.Since(s), d)

	s = time.Now()
	for i := 0; i < rotations; i++ {
		d = fast.Distance(s1, s2)
	}
	fmt.Println("ka-weihe.Distance", time.Since(s), d)
}

func TestOSA(t *testing.T) {
	s := time.Now()
	d := 0

	s = time.Now()
	for i := 0; i < rotations; i++ {
		d = matchr.OSA(s1, s2)
	}
	fmt.Println("matchr.OSA", time.Since(s), d)

	s = time.Now()
	for i := 0; i < rotations; i++ {
		d = edlib.OSADamerauLevenshteinDistance(s1, s2)
	}
	fmt.Println("edlib.OSADamerauLevenshteinDistance", time.Since(s), d)
}

func TestJaro(t *testing.T) {
	d := 0
	f64 := float64(0)
	f32 := float32(0)
	s := time.Now()

	s = time.Now()
	for i := 0; i < rotations; i++ {
		f64 = matchr.JaroWinkler(s1, s2, true)
	}
	fmt.Println("matchr.JaroWinkler", time.Since(s), f64)

	s = time.Now()
	for i := 0; i < rotations; i++ {
		f64 = matchr.JaroWinkler(s1, s2, false)
	}
	fmt.Println("matchr.JaroWinkler", time.Since(s), f64)

	s = time.Now()
	for i := 0; i < rotations; i++ {
		f64 = matchr.Jaro(s1, s2)
	}
	fmt.Println("matchr.Jaro", time.Since(s), f64)

	s = time.Now()
	for i := 0; i < rotations; i++ {
		f32 = edlib.JaroWinklerSimilarity(s1, s2)
	}
	fmt.Println("edlib.JaroWinklerSimilarity", time.Since(s), f32)

	s = time.Now()
	for i := 0; i < rotations; i++ {
		f32 = edlib.JaroSimilarity(s1, s2)
	}
	fmt.Println("edlib.JaroSimilarity", time.Since(s), f32)

	s = time.Now()
	for i := 0; i < rotations; i++ {
		f64 = stringdist.JaroWinkler(s1, s2)
	}
	fmt.Println("stringdist.JaroWinkler", time.Since(s), f64)

	s = time.Now()
	for i := 0; i < rotations; i++ {
		f64 = stringdist.Jaro(s1, s2)
	}
	fmt.Println("stringdist.Jaro", time.Since(s), f64)

	s = time.Now()
	for i := 0; i < rotations; i++ {
		f64, d = textdistance.JaroDistance(s1, s2)
	}
	fmt.Println("textdistance.JaroDistance", time.Since(s), f64, d)
}
