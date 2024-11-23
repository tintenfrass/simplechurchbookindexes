package ui

import (
	"fmt"
	"onlinefuzzysearch/config"
	"onlinefuzzysearch/search"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

var grid = map[int]map[int]string{}

func (h *searchComp) OnMount(ctx app.Context) {
	start := time.Now()
	search.LoadData()
	dur := time.Since(start)

	//Defaults
	h.slideValueMin = config.YearMin
	h.slideValueMax = config.YearMax

	count := 0
	h.checked = make(map[string]bool)
	for key, val := range search.Data.Marriages {
		h.checked[key] = true
		count += len(val.Data)
	}

	h.debug = fmt.Sprintf("%d Datensätze geladen in %s", count, dur.Round(time.Millisecond).String())
}

func replace(input string) (output string) {
	output = input
	output = strings.Replace(output, "dresden/", "", 1)
	output = strings.Replace(output, "meissen/", "", 1)
	output = strings.Replace(output, "freiberg/", "", 1)
	output = strings.Replace(output, "dippoldiswalde/", "", 1)
	output = strings.Replace(output, "Dresden", "DD", 1)
	output = strings.Replace(output, "Böhmische", "Böhm.", 1)
	output = strings.Replace(output, "Exulantengemeinde", "Exulanten", 1)
	output = strings.Replace(output, "Meißen ", "MEI ", 1)
	output = strings.Replace(output, "Freiberg ", "FG ", 1)
	output = strings.Replace(output, "Friedrichstadt St. Michael", "Friedrichstadt", 1)
	return
}

func replaceKK(input string) (output string) {
	output = input
	output = strings.Replace(output, "dresden", "DD", 1)
	output = strings.Replace(output, "meissen", "MEI", 1)
	output = strings.Replace(output, "freiberg", "FG", 1)
	output = strings.Replace(output, "dippoldiswalde", "DW", 1)

	return
}

var rows = 21
var cols = 10

func (h *searchComp) Render() app.UI {

	kath := map[string]struct{}{
		"dresden/Dresden Hofkirche":                  {},
		"dresden/Dresden Neustadt":                   {},
		"dresden/Dresden Friedrichstadt St. Michael": {},
		"meissen/Meißen":                             {},
		"freiberg/Freiberg":                          {},
	}

	grid = map[int]map[int]string{
		0: make(map[int]string),
		1: make(map[int]string),
		2: make(map[int]string),
		3: make(map[int]string),
		4: make(map[int]string),
		5: make(map[int]string),
		6: make(map[int]string),
		7: make(map[int]string),
		8: make(map[int]string),
		9: make(map[int]string),
	}
	grid[0][0] = "meissen/Bloßwitz"
	grid[1][0] = "meissen/Staucha"
	grid[2][0] = "meissen/Striegnitz"
	grid[3][0] = "meissen/Dörschnitz"
	grid[5][0] = "meissen/Großdobritz"
	grid[8][0] = "dresden/Großdittmannsdorf"

	grid[1][1] = "meissen/Lommatzsch"
	grid[2][1] = "meissen/Zehren"
	grid[3][1] = "meissen/Zadel"
	grid[4][1] = "meissen/Gröbern"
	grid[5][1] = "meissen/Oberau"
	grid[8][1] = "dresden/Medingen"
	grid[9][1] = "dresden/Ottendorf"

	grid[0][2] = "meissen/Neckanitz"
	grid[2][2] = "meissen/Meißen St. Afra"
	grid[3][2] = "meissen/Meißen Trinitatiskirche"
	grid[4][2] = "meissen/Niederau"
	grid[7][2] = "dresden/Grünberg"
	grid[8][2] = "dresden/Seifersdorf"
	grid[9][2] = "dresden/Wachau"

	grid[0][3] = "meissen/Leuben"
	grid[1][3] = "meissen/Planitz"
	//grid[2][3] = "meissen/Meißen"
	grid[3][3] = "meissen/Meißen Frauenkirche"
	grid[4][3] = "meissen/Meißen Johanneskirche"
	grid[5][3] = "meissen/Weinböhla"
	grid[8][3] = "dresden/Lausa"
	grid[9][3] = "dresden/Schönborn"

	grid[1][4] = "meissen/Ziegenhain"
	grid[4][4] = "meissen/Brockwitz"
	grid[5][4] = "meissen/Coswig"
	grid[6][4] = "dresden/Reichenberg"
	grid[7][4] = "dresden/Wilschdorf"
	grid[8][4] = "dresden/Radeberg"
	grid[9][4] = "dresden/Langebrück"

	grid[2][5] = "meissen/Krögis"
	grid[3][5] = "meissen/Miltitz"
	grid[5][5] = "meissen/Naustadt"
	grid[6][5] = "dresden/Kötzschenbroda"
	grid[7][5] = "dresden/Klotzsche"
	grid[8][5] = "dresden/Dresden Neustadt"
	grid[9][5] = "dresden/Kleinwolmsdorf"

	grid[0][6] = "meissen/Rüsseina"
	grid[1][6] = "meissen/Raußlitz"
	grid[2][6] = "meissen/Heynitz"
	grid[3][6] = "meissen/Taubenheim"
	grid[4][6] = "meissen/Röhrsdorf"
	grid[5][6] = "meissen/Constappel"
	grid[6][6] = "dresden/Dresden Friedrichstadt St. Michael"
	grid[7][6] = "dresden/Kaditz"
	grid[8][6] = "dresden/Dresden Dreikönigskirche"
	grid[9][6] = "dresden/Großerkmannsdorf"

	grid[0][7] = "meissen/Wendischbora"
	grid[1][7] = "meissen/Rothschönberg"
	grid[2][7] = "meissen/Burkhardswalde"
	grid[3][7] = "meissen/Sora"
	grid[4][7] = "meissen/Weistropp"
	grid[5][7] = "dresden/Briesnitz"
	grid[6][7] = "dresden/Dresden Friedrichstadt"
	grid[7][7] = "dresden/Dresden Sophienkirche"
	grid[8][7] = "dresden/Dresden Hofkirche"
	grid[9][7] = "dresden/Weißig"

	grid[0][8] = "meissen/Nossen"
	grid[1][8] = "meissen/Deutschenbora"
	grid[2][8] = "meissen/Tanneberg"
	grid[3][8] = "meissen/Limbach"
	grid[4][8] = "meissen/Wilsdruff"
	grid[5][8] = "meissen/Unkersdorf"
	grid[6][8] = "dresden/Dresden Kreuzkirche"
	grid[7][8] = "dresden/Dresden Frauenkirche"
	grid[8][8] = "dresden/Loschwitz"

	grid[0][9] = "meissen/Siebenlehn"
	grid[1][9] = "meissen/Hirschfeld"
	grid[2][9] = "meissen/Neukirchen"
	grid[3][9] = "meissen/Blankenstein"
	grid[4][9] = "meissen/Grumbach"
	grid[5][9] = "meissen/Kesselsdorf"
	grid[6][9] = "dippoldiswalde/Pesterwitz"
	grid[7][9] = "dresden/Dresden Annenkirche"
	grid[8][9] = "dresden/Dresden Böhmische Exulantengemeinde"
	grid[9][9] = "dresden/Schönfeld"

	grid[0][10] = "meissen/Obergruna"
	grid[1][10] = "meissen/Bieberstein"
	grid[2][10] = "meissen/Reinsberg"
	grid[3][10] = "meissen/Dittmannsdorf"
	grid[4][10] = "meissen/Herzogswalde"
	grid[5][10] = "dippoldiswalde/Döhlen"
	grid[6][10] = "dresden/Plauen"
	grid[7][10] = "dresden/Leubnitz"
	grid[8][10] = "dresden/Leuben"
	grid[9][10] = "dresden/Hosterwitz"

	grid[0][11] = "freiberg/Großschirma"
	grid[1][11] = "freiberg/Krummenhennersdorf"
	grid[2][11] = "freiberg/Niederschöna"
	grid[3][11] = "meissen/Mohorn"
	grid[4][11] = "dippoldiswalde/Fördergersdorf"
	grid[5][11] = "dippoldiswalde/Tharandt"
	grid[8][11] = "dresden/Lockwitz"

	grid[0][12] = "freiberg/Langhennersdorf"
	grid[1][12] = "freiberg/Tuttendorf"
	grid[2][12] = "freiberg/Conradsdorf"
	grid[3][12] = "freiberg/Naundorf"
	grid[4][12] = "dippoldiswalde/Dorfhain"
	grid[5][12] = "dippoldiswalde/Somsdorf"
	grid[6][12] = "dippoldiswalde/Rabenau"
	grid[7][12] = "dippoldiswalde/Possendorf"
	grid[8][12] = "dresden/Röhrsdorf"

	grid[0][13] = "freiberg/Bräunsdorf"
	grid[1][13] = "freiberg/Freiberg Dom St. Marien"
	grid[2][13] = "freiberg/Freiberg St. Nikolai"
	grid[4][13] = "dippoldiswalde/Klingenberg"
	grid[5][13] = "dippoldiswalde/Höckendorf"
	grid[6][13] = "dippoldiswalde/Seifersdorf"
	grid[8][13] = "dippoldiswalde/Kreischa"

	grid[0][14] = "freiberg/Kleinwaltersdorf"
	grid[1][14] = "freiberg/Freiberg St. Petri"
	grid[2][14] = "freiberg/Freiberg St. Jacobi"
	grid[3][14] = "freiberg/Hilbersdorf"
	grid[4][14] = "dippoldiswalde/Colmnitz"
	grid[5][14] = "dippoldiswalde/Ruppendorf"
	grid[6][14] = "dippoldiswalde/Dippoldiswalde"
	grid[7][14] = "dippoldiswalde/Reinhardtsgrimma"

	grid[0][15] = "freiberg/Kleinschirma"
	grid[1][15] = "freiberg/Freiberg St. Johannis"
	grid[2][15] = "freiberg/Freiberg"
	grid[3][15] = "freiberg/Niederbobritzsch"
	grid[6][15] = "dippoldiswalde/Reichstädt"

	grid[0][16] = "freiberg/Oberschöna"
	grid[1][16] = "freiberg/Erbisdorf"
	grid[2][16] = "freiberg/Berthelsdorf"
	grid[3][16] = "freiberg/Weißenborn"
	grid[4][16] = "freiberg/Oberbobritzsch"
	grid[5][16] = "dippoldiswalde/Pretzschendorf"
	grid[8][16] = "dippoldiswalde/Glashütte"

	grid[0][17] = "freiberg/Langenau"
	grid[1][17] = "freiberg/Weigmannsdorf"
	grid[2][17] = "freiberg/Lichtenberg"
	grid[3][17] = "dippoldiswalde/Burkersdorf"
	grid[4][17] = "dippoldiswalde/Hartmannsdorf"
	grid[5][17] = "dippoldiswalde/Hennersdorf"
	grid[6][17] = "dippoldiswalde/Sadisdorf"
	grid[7][17] = "dippoldiswalde/Schmiedeberg"
	grid[8][17] = "dippoldiswalde/Johnsbach"
	grid[9][17] = "dippoldiswalde/Dittersdorf"

	grid[0][18] = "freiberg/Gränitz"
	grid[1][18] = "freiberg/Großhartmannsdorf"
	grid[2][18] = "freiberg/Helbigsdorf"
	grid[3][18] = "freiberg/Mulda"
	grid[4][18] = "dippoldiswalde/Frauenstein"
	grid[7][18] = "dippoldiswalde/Bärenstein"
	grid[8][18] = "dippoldiswalde/Lauenstein"
	grid[9][18] = "dippoldiswalde/Liebenau"

	grid[1][19] = "freiberg/Zethau"
	grid[2][19] = "freiberg/Dorfchemnitz"
	grid[3][19] = "dippoldiswalde/Dittersbach"
	grid[4][19] = "dippoldiswalde/Nassau"
	grid[5][19] = "dippoldiswalde/Hermsdorf"
	grid[6][19] = "dippoldiswalde/Schellerhau"
	grid[7][19] = "dippoldiswalde/Altenberg"
	grid[8][19] = "dippoldiswalde/Geising"
	grid[9][19] = "dippoldiswalde/Fürstenwalde"

	grid[1][20] = "freiberg/Voigtsdorf"
	grid[2][20] = "freiberg/Sayda"
	grid[3][20] = "freiberg/Clausnitz"
	grid[4][20] = "freiberg/Cämmerswalde"
	grid[8][20] = "dippoldiswalde/Fürstenau"

	raw := "https://raw.githubusercontent.com/tintenfrass/simplechurchbookindexes/main/docs/"

	//Checkboxes
	cbs := map[string]app.HTMLDiv{}
	for key, _ := range search.Data.Marriages {
		k := key //prevent Bug
		changeValue := func(ctx app.Context, e app.Event) {
			h.checked[k] = !h.checked[k]
		}

		//Plus
		plus := func(ctx app.Context, e app.Event) {
			h.plusminus(true, k)
		}
		//Minus
		minus := func(ctx app.Context, e app.Event) {
			h.plusminus(false, k)
		}

		cb := app.Input().Type("checkbox").Checked(h.checked[k]).OnChange(changeValue).Attr("style", "float:left;")
		text := replace(k)
		if _, exists := kath[k]; exists {
			cb = app.Input().Type("checkbox").Checked(h.checked[k]).OnChange(changeValue).Attr("style", "float:left;accent-color:darkred")
			text += " (kath)"
		}
		cbs[k] = app.Div().Body(
			app.Div().Style("float", "left").Body(
				app.Img().Src(raw+"plus.jpg").Style("display", "flex").Style("margin-top", "1px").OnClick(plus).Title("mehrfach Klicken um dieses Suchgebiet zu vergrößern"),
				app.Img().Src(raw+"minus.jpg").Style("display", "flex").Style("margin-top", "0px").OnClick(minus).Title("mehrfach Klicken um dieses Suchgebiet zu verkleiner"),
			),
			cb,
			app.Div().Style("float", "left").Body(
				app.Label().Text(text).OnClick(changeValue),
				app.Br(),
				app.Label().Text(search.GetMinMax(k)).OnClick(changeValue).Attr("style", "color:dimgrey;font-size:7pt"),
			),
		)
	}

	return app.Div().Body(
		app.Table().Body(
			app.Tr().Body(
				app.Td().Body(
					app.Div().Body(
						app.H2().Body(app.Text("Onlinsuche im Trau-Schnell-Index")),
						app.H3().Body(app.Text("Dresden-Meißen-Freiberg-Dippoldiswalde")),
						app.Label().Text("Jahr Min "),
						app.Input().Type("range").Attr("min", config.YearMin).Attr("max", config.YearMax).OnChange(h.ValueTo(&h.slideValueMin)).Value(h.slideValueMin).Attr("style", "width: 600px"),
						app.Label().Text(h.slideValueMin),
						app.Br(),
						app.Label().Text("Jahr Max"),
						app.Input().Type("range").Attr("min", config.YearMin).Attr("max", config.YearMax).OnChange(h.ValueTo(&h.slideValueMax)).Value(h.slideValueMax).Attr("style", "width: 600px"),
						app.Label().Text(h.slideValueMax),
						app.Br(),
						app.Br(),
						app.Br(),
					),
				),
				app.Td().Body(
					app.Div().Style("margin-left", "60px").Style("color", "dimgrey").Body(
						app.Div().Style("font-weight", "bold").Body(app.Text("Erweiterte Einstellung zum Suchalgorithmus (experimentell):")).Title("Bestimmt wie schnell und genau die Ähnlichkeitssuche funktioniert, je nach Algorithmus können die Ergebnisse leicht abweichen. Im Zweifelsfall einfach so lassen."),
						app.Br(),
						app.Div().Title("Ergebnisse werden mit Jaro vorgefiltern und nur die Ergebnisse mit einem Wert größer als 0.4 werden danach mit DamerauLevenshtein genau berechnet. Dieses zweistufe Vorgehen macht die Suche schnell und trotzdem genau.").Body(
							app.Input().Type("radio").Checked(true).OnChange(func(ctx app.Context, e app.Event) { h.algo = search.JaroDamerauLevenshtein }).Name("algo"),
							app.Text("Jaro + DamerauLevenshtein"),
							app.Text(" ----> best Kombination aus Geschwindigkeit und Komplexität"),
						),
						app.Div().Title("Ergebnisse werden mit Jaro vorgefiltern und nur die Ergebnisse mit einem Wert größer als 0.3 werden danach mit Soundex umgewandelt und diese dann mit DamerauLevenshtein berechnet. Das ist sehr langsam aber findet viel mehr Treffer. Große Suchen brechen manchmal ab.").Body(
							app.Input().Type("radio").Checked(false).OnChange(func(ctx app.Context, e app.Event) { h.algo = search.DamerauLevenshtein }).Name("algo"),
							app.Text("Jaro + Soundex + DamerauLevenshtein"),
							app.Text(" ----> mehr Treffer, sehr langsam, nicht für große Suchen geeignet"),
						),
						app.Div().Title("Die OSA-Variante von DamerauLevenshtein, ohne Vorfilterung").Body(
							app.Input().Type("radio").Checked(false).OnChange(func(ctx app.Context, e app.Event) { h.algo = search.Osa }).Name("algo"),
							app.Text("OSA"),
							app.Text(" ----> einfacher und schneller als DamerauLevenshtein, für schwächere Geräte"),
						),
						app.Div().Title("einfachere Version von DamerauLevenshtein, ohne Vorfilterung").Body(
							app.Input().Type("radio").Title("einfachere Version von DamerauLevenshtein").Checked(false).OnChange(func(ctx app.Context, e app.Event) { h.algo = search.Levenshtein }).Name("algo"),
							app.Text("Levenshtein"),
							app.Text(" ----> einfacher und schneller als OSA, für schwache Geräte"),
						),
						app.Br(),
						app.A().Href("https://de.wikipedia.org/wiki/Levenshtein-Distanz").Text(" (Damerau-)Levenshtein-Distanz").Style("color", "grey"),
						app.Br(),
						app.A().Href("https://srinivas-kulkarni.medium.com/jaro-winkler-vs-levenshtein-distance-2eab21832fd6").Text(" Jaro-Winkler vs. Levenshtein").Style("color", "grey"),
						app.Br(),
						app.A().Href("https://de.wikipedia.org/wiki/K%C3%B6lner_Phonetik").Text(" Soundex (Kölner Phonetik)").Style("color", "grey"),
					),
				),
			),
		),
		app.Table().Body(
			app.Div().Style("border", "1px solid #D3D3D3").Body(
				func() (row []app.UI) {
					for j := 0; j < rows; j++ {
						row = append(row, app.Tr().Body(
							func() (ele []app.UI) {
								for i := 0; i < cols; i++ {
									if val, ok := grid[i][j]; ok {
										ele = append(ele, app.Td().Body(cbs[val]))
									} else {
										ele = append(ele, app.Td())
									}
								}
								return
							}()...,
						))
					}

					//Extra Buttons
					row = append(row, app.Tr().Body(
						app.Td().Body(app.Button().Text("Alles").OnClick(h.all).Attr("style", "width: 70px")),
						app.Td().Body(app.Button().Text("Nichts").OnClick(h.nothing).Attr("style", "width: 70px")),
					))

					return
				}()...,
			),
		),
		app.P().Body(
			app.Input().Type("text").Placeholder("Vorname Nachname").AutoFocus(true).OnChange(h.ValueTo(&h.searchValue)).Attr("style", "width: 250px"),
			app.Text(" "),
			app.Button().Text("Search").OnClick(h.onClick).Attr("style", "width: 100px"),
			app.Text(" "),
			app.Label().Text(h.debug),
		),
		app.Div().Body(h.results...),
		app.H3().Body().Text("Hinweise:"),
		app.Text("Die Suche erfolgt über den Namen des Bräutigams."),
		app.Br(),
		app.Text("Dabei wird eine Ähnlichkeitssuche benutzt und die Ergebnisse nach Treffergenauigkeit aufgelistet."),
		app.Br(),
		app.Text("Es wird Groß- und kleinschreibung unterschieden."),
		app.Br(),
		app.Br(),
		app.Text("Bsp: für die Suche:"),
		app.Br(),
		app.Text("Max Mustermann => Sucht nach Vor- und Nachnamen in dieser Kombination"),
		app.Br(),
		app.Text("Man kann auch mit mehrere Vornamen suchen z.B: Max Moritz Mustermann"),
		app.Br(),
		app.Br(),
		app.Text("* als Platzhalter:"),
		app.Br(),
		app.Text("Max * => Sucht nach Vornamen Max"),
		app.Br(),
		app.Text("* Mustermann => Sucht nach Nachname Mustermann"),
		app.Br(),
		app.Br(),
		app.Text("In den Ergebnissen führt der erste Link zur Datei mit den Roh-Daten, das kann helfen die genaue Position des Eintrages im Kirchenbuch zu finden."),
		app.Br(),
		app.Text("Manche Kirchenbücher sind auch nicht chronologisch bzw. die Daten im Buch verstreut."),
		app.Br(),
		app.Text("Es gibt auch immer mal wieder Fehler in den Daten, falsch erfasste Namen oder manchmal ist das Jahr um 1 verrutscht, etc."),
		app.Br(),
		app.Br(),
		app.Text("Der zweie Link führt direkt zum Kirchenbuch, teilweise auch auf den richtigen Scan, das kann durch Fehler aber evtl. auch etwas abweichen."),
		app.Br(),
		app.Text("Die Archion-Links bringen nur was, wenn man bei Archion einen Pass hat."),
		app.Br(),
		app.Br(),
		app.Text("Die Kirchengemeinden sind grob geographisch angeordnet, siehe auch: "),
		app.A().Href("https://www.google.com/maps/d/viewer?mid=1FYfIGUV4g66wImeIqkkr8lcs8kzaAx4s&ll=50.96592383350824%2C13.63222561152344&z=10").Text("Karte Kirchenbücher Sachsen"),
		app.Br(),
		app.Br(),
		app.H3().Body().Text("Mehr Informationen zum Projekt:"),
		app.A().Href("https://github.com/tintenfrass/simplechurchbookindexes").Text("https://github.com/tintenfrass/simplechurchbookindexes"),
		app.Br(),
		app.Br(),
		app.H3().Body().Text(" v1.10 (November 2024) latest updates:"),
		app.Label().Text("Trauungen 1820-1839 hinzugefügt"),
		app.Br(),
		app.Label().Text("Trauungen Freiberg (katholisch) hinzugefügt"),
		//app.Br(),
		//app.Label().Text("Trauungen Meißen (katholisch) hinzugefügt"),
		//app.Br(),
		//app.Label().Text("kleinere Fehlerkorrekturen"),
	).Attr("style", "font-family:verdana,sans-serif;font-size:8pt")
}

const linkPrefix = "https://github.com/tintenfrass/simplechurchbookindexes/blob/main/sachsen/"

func (h *searchComp) onClick(ctx app.Context, e app.Event) {
	start := time.Now()

	h.results = []app.UI{}
	full := -1

	boxes := make(map[int][]app.UI, 10)
	for i := 0; i < 8; i++ {
		boxes[i] = []app.UI{}
	}

	data, debug := search.FindMarriage(h.searchValue, h.slideValueMin, h.slideValueMax, h.checked, h.algo)
	for _, res := range data {
		parts := strings.Split(res, "#")
		dis, _ := strconv.Atoi(parts[2])
		if dis > 7 {
			break
		}
		full = dis
		src := getSource(parts[3])
		if parts[4] != "0" && src == "Archion" {
			parts[3] += "?pageId=" + parts[4]
		}

		boxes[dis] = append(boxes[dis], app.Tr().Body(
			app.Td().Body(app.Label().Text("»»»").Style("font-weight", "bold").Attr("style", "color: "+getColor(dis))),
			app.Td().Body(app.Text(parts[0])),
			app.Td().Body(app.Label().Style("margin", "16px")),
			app.Td().Body(app.Text(replaceKK(path.Dir(parts[1])))).Style("color", "dimgrey;font-size:7pt"),
			app.Td().Body(app.A().Href(linkPrefix+parts[1]).Text(path.Base(parts[1]))),
			app.Td().Body(app.Label().Style("margin", "16px")),
			app.Td().Body(app.A().Href(parts[3]).Text(src)),
		))
	}

	for i := 0; i < 8; i++ {
		if len(boxes[i]) == 0 {
			boxes[i] = append(boxes[i], app.Text("--------"))
			if i > full {
				if full == -1 {
					h.results = append(h.results, app.H4().Body().Text("Keine Ergebnisse gefunden").Style("color", "red"))
				}
				break
			}
		}

		//result table
		rs := []app.UI{}
		for _, b := range boxes[i] {
			rs = append(rs, b)
		}
		tbl := app.Table().Body(
			rs...,
		)

		h.results = append(h.results, app.H4().Body().Text(fmt.Sprintf("Ergebnisse mit Abweichung von ~%d-%d:", i, i+1)), tbl)
	}

	dur := time.Since(start)
	h.debug = fmt.Sprintf("Suchzeit: %s%s", dur.Round(time.Millisecond).String(), debug)
}

func (h *searchComp) all(ctx app.Context, e app.Event) {
	for key, _ := range search.Data.Marriages {
		h.checked[key] = true
	}
}

func (h *searchComp) nothing(ctx app.Context, e app.Event) {
	for key, _ := range search.Data.Marriages {
		h.checked[key] = false
	}
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

func getPos(key string) (posi, posj int) {
	posi = -1
	posj = -1
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			if grid[i][j] == key {
				posi = i
				posj = j
				break
			}
		}
	}
	return
}

func (h *searchComp) plusminus(value bool, k string) {
	posi, posj := getPos(k)

	for r := 0; r < 42; r++ {
		next := true
		for i := 0; i < cols+r; i++ {
			if i >= posi-r && i <= posi+r {
				for j := 0; j < rows+r; j++ {
					if j >= posj-r && j <= posj+r {
						if val, ok := grid[i][j]; ok {
							if h.checked[val] != value {
								h.checked[val] = value
								if r > 0 {
									next = false
								}
							}
						}
					}
				}
			}
		}
		if !next {
			break
		}
	}
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
