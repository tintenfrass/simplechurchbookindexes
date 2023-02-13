package ui

import (
	"indexfuzzysearch/config"
	"indexfuzzysearch/search"

	"github.com/gen2brain/iup-go/iup"
)

const Version = "v1.2"

var boxes map[string]iup.Ihandle

func BuildAndRun() {
	//Control
	iup.Open()
	defer iup.Close()

	//Config
	minSlider := iup.Val("HORIZONTAL").SetAttribute("SIZE", "200x15").SetAttribute("TITLE", "min").SetAttribute("MIN", config.YearMin).SetAttribute("MAX", config.YearMax).SetAttribute("VALUE", config.Config.Year["min"])
	minLabel := iup.Label("").SetHandle("min").SetAttribute("TITLE", config.Config.Year["min"])
	maxSlider := iup.Val("HORIZONTAL").SetAttribute("SIZE", "200x15").SetAttribute("TITLE", "max").SetAttribute("MIN", config.YearMin).SetAttribute("MAX", config.YearMax).SetAttribute("VALUE", config.Config.Year["max"])
	maxLabel := iup.Label("").SetHandle("max").SetAttribute("TITLE", config.Config.Year["max"])
	//instantSearch := iup.Toggle("Suchen beim Tippen").SetAttribute("VALUE", getOnOffInstantSearch("instantSearch")).SetAttribute("key", "instantSearch")

	//Buttons
	exitButton := iup.Button("Exit").SetAttribute("SIZE", "50x15")
	searchButton := iup.Button("Suche").SetAttribute("SIZE", "50x15")
	selectAllButton := iup.Button(utf82ui("Alles")).SetAttribute("SIZE", "50x10")
	selectNoneButton := iup.Button(utf82ui("Nichts")).SetAttribute("SIZE", "50x10")

	//Search and Find
	searchField := iup.Text().SetAttribute("SIZE", "150x").SetHandle("searchField")
	results := iup.MultiLine().SetAttribute("SIZE", "400x440").SetAttribute("READONLY", "YES").SetHandle("output")

	configs := iup.GridBox(
		iup.Label("Min:").SetAttribute("SIZE", "30x15"),
		minSlider,
		minLabel,
		iup.Label("Max:").SetAttribute("SIZE", "30x15"),
		maxSlider,
		maxLabel,
		//instantSearch.SetAttribute("SIZE", "x15"),
	).SetAttribute("NUMDIV", 3)
	//Map
	boxes = map[string]iup.Ihandle{
		"Dörschnitz":  iup.Toggle(utf82ui("Dörschnitz")),
		"Großdobritz": iup.Toggle(utf82ui("Großdobritz")),

		"Lommatzsch": iup.Toggle("Lommatzsch"),
		"Zehren":     iup.Toggle("Zehren"),
		"Zadel":      iup.Toggle("Zedel"),
		"Gröbern":    iup.Toggle(utf82ui("Gröbern")),
		"Oberau":     iup.Toggle("Oberau"),

		"Niederau": iup.Toggle("Niederau"),

		"Leuben":                  iup.Toggle("Leuben"),
		"Meißen St. Afra":         iup.Toggle(utf82ui("Meißen St. Afra")),
		"Meißen Trinitatiskirche": iup.Toggle(utf82ui("Meißen Trinitatiskirche")),
		"Weinböhla":               iup.Toggle(utf82ui("Weinböhla")),

		"Planitz":               iup.Toggle("Planitz"),
		"Meißen Frauenkirche":   iup.Toggle(utf82ui("Meißen Frauenkirche")),
		"Meißen Johanneskirche": iup.Toggle(utf82ui("Meißen Johanneskirche")),
		"Reichenberg":           iup.Toggle("Reichenberg"),
		"Wilschdorf":            iup.Toggle("Wilschdorf"),

		"Ziegenhain": iup.Toggle("Ziegenhain"),
		"Klotzsche":  iup.Toggle("Klotzsche"),

		"Raußlitz":       iup.Toggle(utf82ui("Raußlitz")),
		"Krögis":         iup.Toggle(utf82ui("Krögis")),
		"Naustadt":       iup.Toggle("Naustadt"),
		"Constappel":     iup.Toggle("Constappel"),
		"Kötzschenbroda": iup.Toggle(utf82ui("Kötzschenbroda")),

		"Miltitz": iup.Toggle("Miltitz"),

		"Wendischbora": iup.Toggle("Wendischbora"),
		"Heynitz":      iup.Toggle("Heynitz"),
		"Taubenheim":   iup.Toggle("Taubenheim"),
		"Röhrsdorf":    iup.Toggle(utf82ui("Röhrsdorf")),
		"Weistropp":    iup.Toggle("Weistropp"),
		"Kaditz":       iup.Toggle("Kaditz"),

		"Rothschönberg":            iup.Toggle(utf82ui("Rothschönberg")),
		"Burkhardswalde":           iup.Toggle("Burkhardswalde"),
		"Dresden Dreikönigskirche": iup.Toggle(utf82ui("Dresden Dreikönigskirche")),

		"Deutschenbora":         iup.Toggle("Deutschenbora"),
		"Tanneberg":             iup.Toggle("Tanneberg"),
		"Limbach":               iup.Toggle("Limbach"),
		"Wilsdruff":             iup.Toggle("Wilsdruff"),
		"Unkersdorf":            iup.Toggle("Unkersdorf"),
		"Briesnitz":             iup.Toggle("Briesnitz"),
		"Dresden Sophienkirche": iup.Toggle("Dresden Sophienkirche"),

		"Hirschfeld":          iup.Toggle("Hirschfeld"),
		"Neukirchen":          iup.Toggle("Neukirchen"),
		"Blankenstein":        iup.Toggle("Blankenstein"),
		"Grumbach":            iup.Toggle("Grumbach"),
		"Kesselsdorf":         iup.Toggle("Kesselsdorf"),
		"Dresden Annenkirche": iup.Toggle("Dresden Annenkirche"),

		"Mohorn":                              iup.Toggle("Mohorn"),
		"Herzogswalde":                        iup.Toggle("Herzogswalde"),
		"Pesterwitz":                          iup.Toggle("Pesterwitz"),
		"Dresden Böhmische Exulantengemeinde": iup.Toggle("Dresden Böhmische Exulantengemeinde"),

		"Fördergersdorf": iup.Toggle(utf82ui("Fördergersdorf")),
		"Tharandt":       iup.Toggle("Tharandt"),
		"Döhlen":         iup.Toggle(utf82ui("Döhlen")),
		"Plauen":         iup.Toggle("Plauen"),

		"Somsdorf": iup.Toggle("Somsdorf"),

		"Andere": iup.Toggle("Andere"),
	}

	for key, box := range boxes {
		_, ex := config.Config.Churches[key]
		if !ex {
			config.Config.Churches[key] = true
		}

		box.SetAttribute("VALUE", getOnOffChurches(key)).SetAttribute("key", key)
		box.SetAttribute("TITLE", utf82ui(key)+search.GetMinMax(key))
	}

	grid := iup.GridBox(
		boxes["Dörschnitz"].SetAttributes("SIZE=70x15"),
		iup.Space().SetAttributes("SIZE=70x0"),
		iup.Space().SetAttributes("SIZE=90x0"),
		iup.Space().SetAttributes("SIZE=97x0"),
		boxes["Großdobritz"],
		iup.Space().SetAttributes("SIZE=70x0"),
		iup.Space().SetAttributes("SIZE=150x0"),

		boxes["Lommatzsch"],
		boxes["Zehren"],
		boxes["Zadel"],
		boxes["Gröbern"],
		boxes["Oberau"],
		iup.Space(),
		iup.Space(),

		iup.Space().SetAttributes("SIZE=0x15"),
		iup.Space(),
		iup.Space(),
		iup.Space(),
		boxes["Niederau"],
		iup.Space(),
		iup.Space(),

		boxes["Leuben"],
		iup.Space(),
		boxes["Meißen St. Afra"],
		boxes["Meißen Trinitatiskirche"],
		boxes["Weinböhla"],
		iup.Space(),
		iup.Space(),

		boxes["Planitz"],
		iup.Space(),
		boxes["Meißen Frauenkirche"],
		boxes["Meißen Johanneskirche"],
		iup.Space(),
		boxes["Reichenberg"],
		boxes["Wilschdorf"],

		boxes["Ziegenhain"],
		iup.Space(),
		iup.Space(),
		iup.Space(),
		iup.Space(),
		iup.Space(),
		boxes["Klotzsche"],

		boxes["Raußlitz"],
		boxes["Krögis"],
		iup.Space(),
		boxes["Naustadt"],
		boxes["Constappel"],
		boxes["Kötzschenbroda"],
		iup.Space(),

		iup.Space().SetAttributes("SIZE=0x15"),
		boxes["Miltitz"],
		iup.Space(),
		iup.Space(),
		iup.Space(),
		iup.Space(),
		iup.Space(),

		boxes["Wendischbora"],
		boxes["Heynitz"],
		boxes["Taubenheim"],
		boxes["Röhrsdorf"],
		boxes["Weistropp"],
		boxes["Kaditz"],
		iup.Space(),

		boxes["Rothschönberg"],
		boxes["Burkhardswalde"],
		iup.Space(),
		iup.Space(),
		iup.Space(),
		iup.Space(),
		boxes["Dresden Dreikönigskirche"],

		boxes["Deutschenbora"],
		boxes["Tanneberg"],
		boxes["Limbach"],
		boxes["Wilsdruff"],
		boxes["Unkersdorf"],
		boxes["Briesnitz"],
		boxes["Dresden Sophienkirche"],

		boxes["Hirschfeld"],
		boxes["Neukirchen"],
		boxes["Blankenstein"],
		boxes["Grumbach"],
		boxes["Kesselsdorf"],
		iup.Space(),
		boxes["Dresden Annenkirche"],

		iup.Space().SetAttributes("SIZE=0x15"),
		iup.Space(),
		boxes["Mohorn"],
		boxes["Herzogswalde"],
		iup.Space(),
		boxes["Pesterwitz"],
		boxes["Dresden Böhmische Exulantengemeinde"],

		iup.Space().SetAttributes("SIZE=0x15"),
		iup.Space(),
		iup.Space(),
		boxes["Fördergersdorf"],
		boxes["Tharandt"],
		boxes["Döhlen"],
		boxes["Plauen"],

		boxes["Andere"].SetAttributes("SIZE=90x15"),
		iup.Space(),
		iup.Space(),
		iup.Space(),
		boxes["Somsdorf"],
		iup.Space(),
		iup.Space(),

		selectAllButton,
		selectNoneButton,
	).SetAttribute("NUMDIV", 7)

	configFrame := iup.Frame(configs).SetAttribute("TITLE", "Einstellungen").SetAttributes("SIZE=500x70")
	mapFrame := iup.Frame(grid).SetAttribute("TITLE", "Kirchengemeinden")

	//Infotext
	infotext := iup.Label("").SetAttribute("TITLE", utf82ui("Bsp:\n"+
		"Max Mustermann  ->  Suche nach Vor- und Nachnamen\r\n"+
		"Max ?  ->  Suche nach Personen mit Vornamen Max\r\n"+
		"? Mustermann  ->  Suche nach Personen mit Nachnamen Mustermann\r\n"+
		"Max  ->  Suche unabhängig von Vor- oder Nachname\r\n"+
		"\r\n"+
		"Es wird Groß- und kleinschreibung unterschieden.\r\n",
	))

	content := iup.Hbox(
		iup.Vbox(
			configFrame,
			mapFrame,
			iup.Hbox(
				searchField,
				searchButton,
			),
			infotext,
			exitButton,
		).SetAttributes("MARGIN=10x10, GAP=8"),
		iup.Vbox(
			results,
		),
	)

	dlg := iup.Dialog(content).SetAttributes(`TITLE="Fuzzy Search ` + Version + `"`)
	dlg.SetHandle("dlg").SetAttributes("SIZE=1080x480")

	//Callbacks
	iup.SetCallback(exitButton, "ACTION", iup.ActionFunc(exit))
	iup.SetCallback(searchButton, "ACTION", iup.ActionFunc(searchName))
	iup.SetCallback(searchField, "VALUECHANGED_CB", iup.ValueChangedFunc(searchInstant))
	iup.SetCallback(selectAllButton, "ACTION", iup.ActionFunc(selectAll))
	iup.SetCallback(selectNoneButton, "ACTION", iup.ActionFunc(selectNone))
	for _, box := range boxes {
		iup.SetCallback(box, "ACTION", iup.ActionFunc(toogleChurchUI))
	}
	iup.SetCallback(minSlider, "VALUECHANGED_CB", iup.ValueChangedFunc(valueChanged))
	iup.SetCallback(maxSlider, "VALUECHANGED_CB", iup.ValueChangedFunc(valueChanged))
	//iup.SetCallback(instantSearch, "ACTION", iup.ActionFunc(toogleInstantSearch))

	//Run
	iup.Map(dlg)
	iup.Show(dlg)
	iup.MainLoop()
}

func getOnOffChurches(key string) string {
	if config.Config.Churches[key] {
		return "ON"
	}
	return "OFF"
}

func getOnOffInstantSearch(key string) string {
	if config.Config.InstantSearch {
		return "ON"
	}
	return "OFF"
}
