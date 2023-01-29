package search

import (
	"encoding/json"
	"errors"
	"fmt"
	"indexfuzzysearch/config"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
)

type githubFile struct {
	Name string `json:"name"`
	Url  string `json:"download_url"`
	Size int    `json:"size"`
}

var github []githubFile

func ImportFromGit() {
	//Von der Github-Api die Liste der Dateien holen
	list, err := http.Get("https://api.github.com/repos/tintenfrass/simplechurchbookindexes/contents/sachsen")
	if err != nil {
		fmt.Println(err)
	}
	defer list.Body.Close()
	c, err := ioutil.ReadAll(list.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(c, &github)
	if err != nil {
		fmt.Println(err)
	}

	//Die Namen aller Dateien mit Trauungen raussuchen
	remoteFiles := []githubFile{}
	for _, file := range github {
		//Wir wollen nur Trauungen
		if !strings.Contains(file.Name, "Trauungen") {
			continue
		}
		remoteFiles = append(remoteFiles, file)
	}

	count := 0
	if data.marriages == nil {
		data.marriages = make(map[string]churchEntry)
	}
	for _, remoteFile := range remoteFiles {
		//Prüfen, ob Datei lokal schon existiert und size übereinstimmt
		fileInfo, err := os.Stat("sachsen/" + path.Base(remoteFile.Name))
		if !errors.Is(err, os.ErrNotExist) {
			if fileInfo.Size() == int64(remoteFile.Size) {
				count += ImportFromLocal("sachsen/" + path.Base(remoteFile.Name))
				continue
			}
		}

		resp, err := http.Get(remoteFile.Url)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		subcount := importMarriage(strings.Split(string(content), "\n"), path.Base(remoteFile.Name))
		count += subcount
		fmt.Println(fmt.Sprintf("online geladen \"%s\" mit %d Einträgen", path.Base(remoteFile.Name), subcount))

		//Datei local abspeichern
		if config.Config.OnlineOnly {
			continue
		}
		os.Mkdir("sachsen", os.ModePerm)
		localFile, err := os.Create("sachsen/" + path.Base(remoteFile.Name))
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer localFile.Close()
		localFile.WriteString(string(content))
		localFile.Sync()
	}
	fmt.Println(fmt.Sprintf("%d Trauungen aus %d Quellen geladen", count, len(data.marriages)))
}

func ImportFromLocal(fileName string) (count int) {
	fileData, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	count = importMarriage(strings.Split(strings.Replace(string(fileData), "\r", "", -1), "\n"), fileName)
	fmt.Println(fmt.Sprintf("lokal geladen \"%s\" mit %d Einträgen", path.Base(fileName), count))
	return
}

//Importiert eine Datei
func importMarriage(lines []string, fileName string) int {
	fileMarriages := []marriageEntry{}
	year := -1
	minYear := 0
	maxYear := 0
	source := ""
	for nr, line := range lines {
		//In der ersten Zeile steht die Kirchengemeinde
		if nr == 0 {
			parts := strings.Split(line, "(")
			source = strings.TrimSpace(parts[0])
		}

		if len(line) == 0 || line[0:1] == "x" || strings.HasPrefix(line, "...") {
			continue
		}

		//? in Jahreszahlen ersetzen
		_, err := strconv.Atoi(line[0:1])
		if err == nil {
			line = strings.Replace(line, "?", "", -1)
		}

		content := strings.Split(strings.TrimSpace(line), " ")
		//Jahreszahl?
		if len(content) == 1 || len(content) > 2 && strings.Contains(content[1], "Teil") {
			//4-Stellig?
			if len(content[0]) >= 4 {
				if content[0][:4] == "0000" {
					year = 0
				} else {
					number, err := strconv.Atoi(content[0][:4])
					if err != nil || number == 0 {
						continue
					}
					year = number
				}
			}
			continue
		}
		if year < 0 {
			continue
		}

		marriage := marriageEntry{
			simple: line,
			source: fileName,
			year:   year,
		}
		if year > maxYear {
			maxYear = year
		}
		if minYear == 0 || year < minYear {
			minYear = year
		}

		groom := line

		//Braut dabei?
		if strings.Contains(line, "|") {
			persons := strings.Split(line, "|")
			groom = persons[0]
		}

		//Vorname / Nachname trennen
		names := strings.Split(strings.TrimSpace(groom), " ")
		if len(names) == 1 {
			marriage.groom = names[0]
		} else {
			marriage.groom = strings.Join(names[0:len(names)-1], " ")
			marriage.groomFN = names[len(names)-1]
			//Special Case
			for _, special := range specialCases {
				if strings.Contains(groom, special) {
					temp := marriage.groom
					marriage.groom = marriage.groomFN
					marriage.groomFN = temp
					break
				}
			}
		}

		//Alle Trauungen in dieser Datei sammeln
		fileMarriages = append(fileMarriages, marriage)
	}

	//speichern
	if _, ok := data.marriages[source]; ok {
		if data.marriages[source].max > maxYear {
			maxYear = data.marriages[source].max
		}
		if data.marriages[source].min < minYear && data.marriages[source].min > 0 {
			minYear = data.marriages[source].min
		}
	}
	if _, ok := data.marriages[source]; ok && data.marriages[source].max > maxYear {
		maxYear = data.marriages[source].max
	}

	data.marriages[source] = churchEntry{
		min:  minYear,
		max:  maxYear,
		data: append(data.marriages[source].data, fileMarriages...),
	}

	return len(fileMarriages)
}
