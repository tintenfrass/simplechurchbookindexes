package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type JsonConfig struct {
	Year          map[string]int  `json:"zeitraum"`
	Churches      map[string]bool `json:"kirchen"`
	InstantSearch bool            `json:"sofortsuche"`
}

const YearMin = 1550
const YearMax = 1800

var Config JsonConfig

func Load() {
	jsonFile, err := os.Open("config.json")
	if err == nil {
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &Config)
	}

	//Defaultwerte setzen
	if Config.Year == nil {
		Config.Year = map[string]int{"min": YearMin, "max": YearMax}
	}
	if Config.Year["min"] < YearMax {
		Config.Year["min"] = YearMin
	}
	if Config.Year["max"] > YearMax {
		Config.Year["max"] = YearMax
	}

	if Config.Churches == nil {
		Config.Churches = map[string]bool{}
	}
}

func Save() {
	jsonData, _ := json.MarshalIndent(Config, "", " ")
	ioutil.WriteFile("config.json", jsonData, 0644)
}
