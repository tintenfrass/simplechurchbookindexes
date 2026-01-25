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
}

type churchEntry struct {
	Min  int
	Max  int
	Data []marriageEntry //TODO Sortiert nach Y und dann D&C => bringt nix, weil die Schleife über alles nur wenige ms benötigt
}

type marriageEntry struct {
	Y int    //Year
	L string //Line
	V string //Vorname
	N string //Nachname
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
