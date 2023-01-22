package search

import (
	"strconv"
)

type fullData struct {
	marriages map[string]churchEntry
}

var data fullData

type churchEntry struct {
	min  int
	max  int
	data []marriageEntry
}

type marriageEntry struct {
	year    int
	simple  string
	groom   string
	groomFN string
	//bride   string //unused
	//brideFN string //unused
	source string
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
	if val, ok := data.marriages[church]; ok {
		return "\r\n(" + strconv.Itoa(val.min) + "-" + strconv.Itoa(val.max) + ")"
	}
	return ""
}

//TODO Shotcuts: Hans<=>Johann, Balzer<=>Balthasar, Brosius<=>Ambrosius, etc.
