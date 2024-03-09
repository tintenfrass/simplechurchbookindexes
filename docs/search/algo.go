package search

import "github.com/antzucaro/matchr"

const (
	JaroDamerauLevenshtein = 0
	DamerauLevenshtein     = 1
	Osa                    = 2
	Levenshtein            = 3
	JaroTreshold           = 0.4
	JaroTresholdSoundex    = 0.3
)

type algoInterface interface {
	search(s1, s2 string) int
}

func getSearcher(code int) (al algoInterface) {
	switch code {
	case JaroDamerauLevenshtein, DamerauLevenshtein:
		al = damerauLevenshtein{}
	case Osa:
		al = osa{}
	case Levenshtein:
		al = levenshtein{}
	default:
		al = damerauLevenshtein{}
	}

	return
}

type damerauLevenshtein struct{}

func (damerauLevenshtein) search(s1, s2 string) int {
	return matchr.DamerauLevenshtein(s1, s2)
}

type osa struct{}

func (osa) search(s1, s2 string) int {
	return matchr.OSA(s1, s2)
}

type levenshtein struct{}

func (levenshtein) search(s1, s2 string) int {
	return matchr.Levenshtein(s1, s2)
}
