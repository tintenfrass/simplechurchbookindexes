package ui

import (
	"bytes"
	"unicode/utf8"
)

//Umlaute aus dem UI in UTF-8 umwandeln
func ui2utf8(val string) (text string) {

	for i := 0; i < len(val); i++ {
		x := val[i : i+1]

		if bytes.Equal([]byte(x), []byte{228}) {
			text += "ä"
		} else if bytes.Equal([]byte(x), []byte{246}) {
			text += "ö"
		} else if bytes.Equal([]byte(x), []byte{252}) {
			text += "ü"
		} else if bytes.Equal([]byte(x), []byte{223}) {
			text += "ß"
		} else if bytes.Equal([]byte(x), []byte{196}) {
			text += "Ä"
		} else if bytes.Equal([]byte(x), []byte{214}) {
			text += "Ö"
		} else if bytes.Equal([]byte(x), []byte{220}) {
			text += "Ü"
		} else {
			text += x
		}
	}
	return
}

//UTF-8 für die Anzeige im UI umwandeln
func utf82ui(input string) (uitext string) {
	for i, w := 0, 0; i < len(input); i += w {
		_, w = utf8.DecodeRuneInString(input[i:])

		x := input[i : i+w]

		switch x {
		case "ä":
			uitext += string([]byte{228})
		case "ö":
			uitext += string([]byte{246})
		case "ü":
			uitext += string([]byte{252})
		case "ß":
			uitext += string([]byte{223})
		case "Ä":
			uitext += string([]byte{196})
		case "Ö":
			uitext += string([]byte{214})
		case "Ü":
			uitext += string([]byte{220})
		default:
			uitext += x
		}
	}

	return
}
