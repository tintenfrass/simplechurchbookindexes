package search

import (
	"strings"

	"github.com/antzucaro/matchr"
)

const (
	DamerauLevenshtein        = 0
	SoundexDamerauLevenshtein = 1
	Exact                     = 2
)

type Searcher interface {
	search(s1, s2 string) int
}

func getSearcher(code int) (al Searcher) {
	switch code {
	case SoundexDamerauLevenshtein, DamerauLevenshtein:
		al = damerauLevenshtein{}
	case Exact:
		al = exact{}
	default:
		al = damerauLevenshtein{}
	}

	return
}

type damerauLevenshtein struct{}

func (damerauLevenshtein) search(s1, s2 string) int {
	return matchr.DamerauLevenshtein(s1, s2)
}

type exact struct{}

func (exact) search(s1, s2 string) int {
	if strings.Contains(s2, s1) {
		return 0
	}
	return 100
}
