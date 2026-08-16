package search

import (
	"strconv"

	"onlinefuzzysearch/config"
)

type fullData struct {
	Marriages map[string]churchEntry
	Sources   map[int]string
	Links     map[int]string
	Offset    map[int]uint32
	NamesV    map[int]string
	NamesN    map[int]string
	PlaceV    map[string]map[int]struct{}
	PlaceN    map[string]map[int]struct{}
	DecadeV   map[int]map[int]struct{}
	DecadeN   map[int]map[int]struct{}
}

type churchEntry struct {
	Min  int
	Max  int
	Data []marriageEntry
}

type marriageEntry struct {
	Y int    //Year
	L string //Line
	V int    //Vorname
	N int    //Nachname
	S int    //Source
	P int    //Page
}

func GetMinMax(church string) string {
	if val, ok := Data.Marriages[church]; ok {
		max := val.Max
		if max > config.YearMax {
			max = config.YearMax
		}
		return "\r\n(" + strconv.Itoa(val.Min) + "-" + strconv.Itoa(max) + ")"
	}
	return ""
}

var Data fullData
