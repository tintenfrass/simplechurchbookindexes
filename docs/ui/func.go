package ui

import "strings"

func replace(input string) (output string) {
	output = input
	output = strings.Replace(output, "dresden/", "", 1)
	output = strings.Replace(output, "meissen/", "", 1)
	output = strings.Replace(output, "freiberg/", "", 1)
	output = strings.Replace(output, "dippoldiswalde/", "", 1)
	output = strings.Replace(output, "auerbach/", "", 1)
	output = strings.Replace(output, "torgau-delitzsch/", "", 1)
	output = strings.Replace(output, "bautzen/", "", 1)
	output = strings.Replace(output, "bad-liebenwerda/", "", 1)

	output = strings.Replace(output, "Dresden", "DD", 1)
	output = strings.Replace(output, "Böhmische", "Böhm.", 1)
	output = strings.Replace(output, "Exulantengemeinde", "Exulanten", 1)
	output = strings.Replace(output, "Meißen ", "MEI ", 1)
	output = strings.Replace(output, "Freiberg ", "FG ", 1)
	output = strings.Replace(output, "Friedrichstadt St. Michael", "Friedrichstadt", 1)
	output = strings.Replace(output, "Garnisonsgemeinde", "Garnison", 1)
	output = strings.Replace(output, " bei Neiden", "", 1)
	return
}

func replaceKK(input string) (output string) {
	output = input
	output = strings.Replace(output, "dresden", "DD", 1)
	output = strings.Replace(output, "meissen", "MEI", 1)
	output = strings.Replace(output, "freiberg", "FG", 1)
	output = strings.Replace(output, "dippoldiswalde", "DW", 1)
	output = strings.Replace(output, "auerbach", "AU", 1)
	output = strings.Replace(output, "torgau-delitzsch", "TO", 1)
	output = strings.Replace(output, "bautzen", "BZ", 1)
	output = strings.Replace(output, "bad-liebenwerda", "BL", 1)

	return
}

func getColor(distant int) (color string) {
	switch distant {
	case 0:
		color = "green"
	case 1:
		color = "#4CBB17"
	case 2:
		color = "#C4B454"
	case 3:
		color = "#FFC300"
	case 4:
		color = "orange"
	case 5:
		color = "#FF7518"
	default:
		color = "red"
	}

	return
}

func getPos(tab int, key string) (posi, posj int) {
	posi = -1
	posj = -1
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			if grid[tab][i][j] == key {
				posi = i
				posj = j
				break
			}
		}
	}
	return
}

func getSource(link string) string {
	switch {
	case strings.Contains(link, "archion.de"):
		return "Archion"
	case strings.Contains(link, "matricula-online.eu"):
		return "Matricula"
	case strings.Contains(link, "familysearch.org"):
		return "Familysearch"
	default:
		return "Online-Link"
	}
}
