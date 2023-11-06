package search

import (
	"strconv"
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
	Data []marriageEntry
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
		return "\r\n(" + strconv.Itoa(val.Min) + "-" + strconv.Itoa(val.Max) + ")"
	}
	return ""
}

var Data fullData
