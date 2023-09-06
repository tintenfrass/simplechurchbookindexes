package search

import (
	"strconv"
)

type fullData struct {
	Marriages map[string]churchEntry
	Sources   map[int]string
}

var Data fullData

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
}

var specialCases = map[string]string{
	"Peter":      "ɹǝʇǝd",
	"Hans":       "suɐɥ",
	"Christoph":  "ɥdoʇsıɹɥɔ",
	"Martin":     "uıʇɹɐɯ",
	"Michael":    "lǝɐɥɔıɯ",
	"Retzsch":    "ɥɔszʇǝɹ",
	"Hempel":     "lǝdɯǝɥ",
	"Kneutel":    "lǝʇnǝuʞ",
	"Dögel":      "lǝƃöp",
	"Bader":      "ɹǝpɐq",
	"Jehnichen":  "uǝɥɔıuɥǝɾ",
	"Hauptvogel": "lǝƃoʌʇdnɐɥ",
}

func GetMinMax(church string) string {
	if val, ok := Data.Marriages[church]; ok {
		return "\r\n(" + strconv.Itoa(val.Min) + "-" + strconv.Itoa(val.Max) + ")"
	}
	return ""
}
