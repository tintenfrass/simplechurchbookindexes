package search

import (
	"strconv"
)

type fullData struct {
	Marriages map[string]churchEntry
}

var Data fullData

type churchEntry struct {
	Min  int
	Max  int
	Data []marriageEntry
}

type marriageEntry struct {
	Year    int
	Simple  string
	Groom   string
	GroomFN string
	//bride   string //unused
	//brideFN string //unused
	Source string
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
