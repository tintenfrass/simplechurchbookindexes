package ui

import (
	"fmt"
	"time"

	"onlinefuzzysearch/config"
	"onlinefuzzysearch/search"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

const tabs = 3
const rows = 21
const cols = 10

var grid = map[int]map[int]map[int]string{} //tab, row, col, name

// Render the full ui
func (h *searchComp) Render() app.UI {

	if len(h.displayTab) == 0 {
		//show only first tab
		h.displayTab = []string{"none", "block", "none"}
	}

	kath := map[string]struct{}{
		"dresden/Dresden Hofkirche":                  {},
		"dresden/Dresden Neustadt":                   {},
		"dresden/Dresden Friedrichstadt St. Michael": {},
		"meissen/Meißen":                             {},
		"freiberg/Freiberg":                          {},
		"bautzen/Bautzen Dom":                        {},
	}

	grid = map[int]map[int]map[int]string{
		0: {
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
		},
		1: {
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
		},
		2: {
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
		},
	}
	grid[0][6][9] = "torgau-delitzsch/Rosenfeld"

	grid[0][6][10] = "torgau-delitzsch/Zwethau"

	grid[0][6][11] = "torgau-delitzsch/Kreischau-Eulenau"
	grid[0][7][11] = "torgau-delitzsch/Beilrode"

	grid[0][5][12] = "torgau-delitzsch/Torgau St. Marien"

	grid[0][5][13] = "torgau-delitzsch/Torgau Garnisonsgemeinde"
	grid[0][6][13] = "torgau-delitzsch/Torgau Schmerzhafte Mutter"

	grid[0][6][14] = "torgau-delitzsch/Loßwig"
	grid[0][7][14] = "torgau-delitzsch/Triestewitz"
	grid[0][8][14] = "torgau-delitzsch/Arzberg"

	grid[0][7][15] = "torgau-delitzsch/Weßnig"
	grid[0][9][15] = "torgau-delitzsch/Blumberg"

	grid[0][5][16] = "torgau-delitzsch/Beckwitz"

	grid[0][6][17] = "torgau-delitzsch/Taura"
	grid[0][8][17] = "torgau-delitzsch/Belgern"

	grid[0][5][18] = "torgau-delitzsch/Sitzenroda"
	grid[0][7][18] = "torgau-delitzsch/Lausa"
	grid[0][8][18] = "torgau-delitzsch/Neußen"
	grid[0][9][18] = "torgau-delitzsch/Staritz"

	grid[0][9][19] = "torgau-delitzsch/Schirmenitz"

	grid[0][9][20] = "torgau-delitzsch/Paußnitz"

	grid[1][0][0] = "meissen/Bloßwitz"
	grid[1][1][0] = "meissen/Staucha"
	grid[1][2][0] = "meissen/Striegnitz"
	grid[1][3][0] = "meissen/Dörschnitz"
	grid[1][5][0] = "meissen/Großdobritz"
	grid[1][8][0] = "dresden/Großdittmannsdorf"

	grid[1][1][1] = "meissen/Lommatzsch"
	grid[1][2][1] = "meissen/Zehren"
	grid[1][3][1] = "meissen/Zadel"
	grid[1][4][1] = "meissen/Gröbern"
	grid[1][5][1] = "meissen/Oberau"
	grid[1][8][1] = "dresden/Medingen"
	grid[1][9][1] = "dresden/Ottendorf"

	grid[1][0][2] = "meissen/Neckanitz"
	grid[1][2][2] = "meissen/Meißen St. Afra"
	grid[1][3][2] = "meissen/Meißen Trinitatiskirche"
	grid[1][4][2] = "meissen/Niederau"
	grid[1][7][2] = "dresden/Grünberg"
	grid[1][8][2] = "dresden/Seifersdorf"
	grid[1][9][2] = "dresden/Wachau"

	grid[1][0][3] = "meissen/Leuben"
	grid[1][1][3] = "meissen/Planitz"
	grid[1][2][3] = "meissen/Meißen"
	grid[1][3][3] = "meissen/Meißen Frauenkirche"
	grid[1][4][3] = "meissen/Meißen Johanneskirche"
	grid[1][5][3] = "meissen/Weinböhla"
	grid[1][8][3] = "dresden/Lausa"
	grid[1][9][3] = "dresden/Schönborn"

	grid[1][1][4] = "meissen/Ziegenhain"
	grid[1][4][4] = "meissen/Brockwitz"
	grid[1][5][4] = "meissen/Coswig"
	grid[1][6][4] = "dresden/Reichenberg"
	grid[1][7][4] = "dresden/Wilschdorf"
	grid[1][8][4] = "dresden/Radeberg"
	grid[1][9][4] = "dresden/Langebrück"

	grid[1][2][5] = "meissen/Krögis"
	grid[1][3][5] = "meissen/Miltitz"
	grid[1][5][5] = "meissen/Naustadt"
	grid[1][6][5] = "dresden/Kötzschenbroda"
	grid[1][7][5] = "dresden/Klotzsche"
	grid[1][8][5] = "dresden/Dresden Neustadt"
	grid[1][9][5] = "dresden/Kleinwolmsdorf"

	grid[1][0][6] = "meissen/Rüsseina"
	grid[1][1][6] = "meissen/Raußlitz"
	grid[1][2][6] = "meissen/Heynitz"
	grid[1][3][6] = "meissen/Taubenheim"
	grid[1][4][6] = "meissen/Röhrsdorf"
	grid[1][5][6] = "meissen/Constappel"
	grid[1][6][6] = "dresden/Dresden Friedrichstadt St. Michael"
	grid[1][7][6] = "dresden/Kaditz"
	grid[1][8][6] = "dresden/Dresden Dreikönigskirche"
	grid[1][9][6] = "dresden/Großerkmannsdorf"

	grid[1][0][7] = "meissen/Wendischbora"
	grid[1][1][7] = "meissen/Rothschönberg"
	grid[1][2][7] = "meissen/Burkhardswalde"
	grid[1][3][7] = "meissen/Sora"
	grid[1][4][7] = "meissen/Weistropp"
	grid[1][5][7] = "dresden/Briesnitz"
	grid[1][6][7] = "dresden/Dresden Friedrichstadt"
	grid[1][7][7] = "dresden/Dresden Sophienkirche"
	grid[1][8][7] = "dresden/Dresden Hofkirche"
	grid[1][9][7] = "dresden/Weißig"

	grid[1][0][8] = "meissen/Nossen"
	grid[1][1][8] = "meissen/Deutschenbora"
	grid[1][2][8] = "meissen/Tanneberg"
	grid[1][3][8] = "meissen/Limbach"
	grid[1][4][8] = "meissen/Wilsdruff"
	grid[1][5][8] = "meissen/Unkersdorf"
	grid[1][6][8] = "dresden/Dresden Kreuzkirche"
	grid[1][7][8] = "dresden/Dresden Frauenkirche"
	grid[1][8][8] = "dresden/Loschwitz"

	grid[1][0][9] = "meissen/Siebenlehn"
	grid[1][1][9] = "meissen/Hirschfeld"
	grid[1][2][9] = "meissen/Neukirchen"
	grid[1][3][9] = "meissen/Blankenstein"
	grid[1][4][9] = "meissen/Grumbach"
	grid[1][5][9] = "meissen/Kesselsdorf"
	grid[1][6][9] = "dippoldiswalde/Pesterwitz"
	grid[1][7][9] = "dresden/Dresden Annenkirche"
	grid[1][8][9] = "dresden/Dresden Böhmische Exulantengemeinde"
	grid[1][9][9] = "dresden/Schönfeld"

	grid[1][0][10] = "meissen/Obergruna"
	grid[1][1][10] = "meissen/Bieberstein"
	grid[1][2][10] = "meissen/Reinsberg"
	grid[1][3][10] = "meissen/Dittmannsdorf"
	grid[1][4][10] = "meissen/Herzogswalde"
	grid[1][5][10] = "dippoldiswalde/Döhlen"
	grid[1][6][10] = "dresden/Plauen"
	grid[1][7][10] = "dresden/Leubnitz"
	grid[1][8][10] = "dresden/Leuben"
	grid[1][9][10] = "dresden/Hosterwitz"

	grid[1][0][11] = "freiberg/Großschirma"
	grid[1][1][11] = "freiberg/Krummenhennersdorf"
	grid[1][2][11] = "freiberg/Niederschöna"
	grid[1][3][11] = "meissen/Mohorn"
	grid[1][4][11] = "dippoldiswalde/Fördergersdorf"
	grid[1][5][11] = "dippoldiswalde/Tharandt"
	grid[1][6][11] = "dippoldiswalde/Deuben"
	grid[1][8][11] = "dresden/Lockwitz"

	grid[1][0][12] = "freiberg/Langhennersdorf"
	grid[1][1][12] = "freiberg/Tuttendorf"
	grid[1][2][12] = "freiberg/Conradsdorf"
	grid[1][3][12] = "freiberg/Naundorf"
	grid[1][4][12] = "dippoldiswalde/Dorfhain"
	grid[1][5][12] = "dippoldiswalde/Somsdorf"
	grid[1][6][12] = "dippoldiswalde/Rabenau"
	grid[1][7][12] = "dippoldiswalde/Possendorf"
	grid[1][8][12] = "dresden/Röhrsdorf"

	grid[1][0][13] = "freiberg/Bräunsdorf"
	grid[1][1][13] = "freiberg/Freiberg Dom St. Marien"
	grid[1][2][13] = "freiberg/Freiberg St. Nikolai"
	grid[1][4][13] = "dippoldiswalde/Klingenberg"
	grid[1][5][13] = "dippoldiswalde/Höckendorf"
	grid[1][6][13] = "dippoldiswalde/Seifersdorf"
	grid[1][8][13] = "dippoldiswalde/Kreischa"

	grid[1][0][14] = "freiberg/Kleinwaltersdorf"
	grid[1][1][14] = "freiberg/Freiberg St. Petri"
	grid[1][2][14] = "freiberg/Freiberg St. Jacobi"
	grid[1][3][14] = "freiberg/Hilbersdorf"
	grid[1][4][14] = "dippoldiswalde/Colmnitz"
	grid[1][5][14] = "dippoldiswalde/Ruppendorf"
	grid[1][6][14] = "dippoldiswalde/Dippoldiswalde"
	grid[1][7][14] = "dippoldiswalde/Reinhardtsgrimma"

	grid[1][0][15] = "freiberg/Kleinschirma"
	grid[1][1][15] = "freiberg/Freiberg St. Johannis"
	grid[1][2][15] = "freiberg/Freiberg"
	grid[1][3][15] = "freiberg/Niederbobritzsch"
	grid[1][6][15] = "dippoldiswalde/Reichstädt"

	grid[1][0][16] = "freiberg/Oberschöna"
	grid[1][1][16] = "freiberg/Erbisdorf"
	grid[1][2][16] = "freiberg/Berthelsdorf"
	grid[1][3][16] = "freiberg/Weißenborn"
	grid[1][4][16] = "freiberg/Oberbobritzsch"
	grid[1][5][16] = "dippoldiswalde/Pretzschendorf"
	grid[1][8][16] = "dippoldiswalde/Glashütte"

	grid[1][0][17] = "freiberg/Langenau"
	grid[1][1][17] = "freiberg/Weigmannsdorf"
	grid[1][2][17] = "freiberg/Lichtenberg"
	grid[1][3][17] = "dippoldiswalde/Burkersdorf"
	grid[1][4][17] = "dippoldiswalde/Hartmannsdorf"
	grid[1][5][17] = "dippoldiswalde/Hennersdorf"
	grid[1][6][17] = "dippoldiswalde/Sadisdorf"
	grid[1][7][17] = "dippoldiswalde/Schmiedeberg"
	grid[1][8][17] = "dippoldiswalde/Johnsbach"
	grid[1][9][17] = "dippoldiswalde/Dittersdorf"

	grid[1][0][18] = "freiberg/Gränitz"
	grid[1][1][18] = "freiberg/Großhartmannsdorf"
	grid[1][2][18] = "freiberg/Helbigsdorf"
	grid[1][3][18] = "freiberg/Mulda"
	grid[1][4][18] = "dippoldiswalde/Frauenstein"
	grid[1][5][18] = "dippoldiswalde/Schönfeld"
	grid[1][7][18] = "dippoldiswalde/Bärenstein"
	grid[1][8][18] = "dippoldiswalde/Lauenstein"
	grid[1][9][18] = "dippoldiswalde/Liebenau"

	grid[1][1][19] = "freiberg/Zethau"
	grid[1][2][19] = "freiberg/Dorfchemnitz"
	grid[1][3][19] = "dippoldiswalde/Dittersbach"
	grid[1][4][19] = "dippoldiswalde/Nassau"
	grid[1][5][19] = "dippoldiswalde/Hermsdorf"
	grid[1][6][19] = "dippoldiswalde/Schellerhau"
	grid[1][7][19] = "dippoldiswalde/Altenberg"
	grid[1][8][19] = "dippoldiswalde/Geising"
	grid[1][9][19] = "dippoldiswalde/Fürstenwalde"

	grid[1][1][20] = "freiberg/Voigtsdorf"
	grid[1][2][20] = "freiberg/Sayda"
	grid[1][3][20] = "freiberg/Clausnitz"
	grid[1][4][20] = "freiberg/Cämmerswalde"
	grid[1][8][20] = "dippoldiswalde/Fürstenau"

	grid[2][2][8] = "bautzen/Bautzen"
	grid[2][3][8] = "bautzen/Bautzen Dom"

	grid[2][2][9] = "bautzen/Bautzen St. Michael"
	grid[2][3][9] = "bautzen/Bautzen St. Petri"

	raw := "https://raw.githubusercontent.com/tintenfrass/simplechurchbookindexes/main/docs/"

	//Checkboxes
	cbs := make(map[int]map[string]app.HTMLDiv)
	for tab := 0; tab < tabs; tab++ {
		cbs[tab] = make(map[string]app.HTMLDiv)
		for key, _ := range search.Data.Marriages {
			if !isValid(tab, key) {
				continue
			}

			k := key //prevent Bug
			changeValue := func(ctx app.Context, e app.Event) {
				h.checked[h.activeTab][k] = !h.checked[h.activeTab][k]
			}

			//Plus
			plus := func(ctx app.Context, e app.Event) {
				h.plusminus(true, k)
			}
			//Minus
			minus := func(ctx app.Context, e app.Event) {
				h.plusminus(false, k)
			}

			cb := app.Input().Type("checkbox").Checked(h.checked[tab][k]).OnChange(changeValue).Attr("style", "float:left;")
			text := replace(k)
			if _, exists := kath[k]; exists {
				cb = app.Input().Type("checkbox").Checked(h.checked[tab][k]).OnChange(changeValue).Attr("style", "float:left;accent-color:darkred")
				text += " (kath)"
			}
			cbs[tab][k] = app.Div().Body(
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
	}

	return app.Div().Body(
		app.Table().Body(
			app.Tr().Body(
				app.Td().Body(
					app.Div().Body(
						app.H2().Body(app.Text("Onlinesuche im Trauindex")),
						app.H3().Body(app.Text("Dresden-Meißen-Freiberg-Dippoldiswalde (+Teile von Torgau und Bautzen)")),
						app.Label().Text("Jahr Min "),
						app.Input().Type("range").Attr("min", config.YearMin).Attr("max", config.YearMax).OnChange(h.ValueTo(&h.slideValueMin)).Value(h.slideValueMin).Attr("style", "width: 600px"),
						app.Label().Text(h.slideValueMin),
						app.Br(),
						app.Label().Text("Jahr Max "),
						app.Input().Type("range").Attr("min", config.YearMin).Attr("max", config.YearMax).OnChange(h.ValueTo(&h.slideValueMax)).Value(h.slideValueMax).Attr("style", "width: 600px"),
						app.Label().Text(h.slideValueMax),
						app.Br(),
						app.Br(),
						app.Br(),
					),
				),
				app.Td().Body(
					app.Div().Style("margin-left", "60px").Body(
						app.Div().Style("font-weight", "bold").Body(app.Text("Suchalgorithmus:")).Title("Bestimmt wie schnell und genau die Ähnlichkeitssuche funktioniert, je nach Algorithmus können die Ergebnisse abweichen."),
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
							app.Text(" ----> einfacher und schneller als DamerauLevenshtein"),
						),
						app.Div().Title("einfachere Version von DamerauLevenshtein, ohne Vorfilterung").Body(
							app.Input().Type("radio").Checked(false).OnChange(func(ctx app.Context, e app.Event) { h.algo = search.Levenshtein }).Name("algo"),
							app.Text("Levenshtein"),
							app.Text(" ----> einfacher und schneller als OSA"),
						),
						app.Div().Title("Der einzige Algorithmus ohne Ähnlichkeitssuche sondern klassisch nach einer festen Buchstabenkombination.").Body(
							app.Input().Type("radio").Checked(false).OnChange(func(ctx app.Context, e app.Event) { h.algo = search.Exact }).Name("algo"),
							app.Text("Klassisch"),
							app.Text(" ----> sucht buchstabengenau nach dem Suchtext innerhalb des Namens."),
						),
					),
				),
			),
		),
		app.Td().Body(app.Button().Disabled(h.displayTab[0] == "block").Text("TO").OnClick(h.tab0).Title("Torgau")),
		app.Td().Body(app.Button().Disabled(h.displayTab[1] == "block").Text("DD-MEI-FG-DW").OnClick(h.tab1).Title("Dresden-Meißen-Freiberg-Dippoldiswalde")),
		app.Td().Body(app.Button().Disabled(h.displayTab[2] == "block").Text("BZ").OnClick(h.tab2).Title("Bautzen")),
		app.Table().Body(
			getTabContent(h, 0, cbs),
			getTabContent(h, 1, cbs),
			getTabContent(h, 2, cbs),
		),
		app.P().Body(
			app.Input().Type("text").Placeholder("Vorname Nachname").AutoFocus(true).OnChange(h.ValueTo(&h.searchValue)).Attr("style", "width: 250px"),
			app.Text(" "),
			app.Button().Text("Search").OnClick(h.search).Attr("style", "width: 100px"),
			app.Text(" "),
			app.Label().Text(h.debug),
		),
		app.Div().Body(h.results...),
		app.H3().Body().Text("Hinweise:"),
		app.Text("Die Suche erfolgt über alle Vornamen, den Nachnamen oder den kompletten Namen des Bräutigams. Eine Suche über Namensteile ist im Suchmodus Klassisch möglich. (siehe Beispiele)"),
		app.Br(),
		app.Text("Dabei wird eine Ähnlichkeitssuche benutzt und die Ergebnisse nach Treffergenauigkeit aufgelistet. (Ausnahme Suchmodus Klassisch)"),
		app.Br(),
		app.Text("Die Suche wird direkt im Browser ausgeführt, die Suchgeschwindigkeit ist damit stark abhängig vom Gerät womit diese Seite aufgerufen wird."),
		app.Br(),
		app.Text("Es wird Groß- und kleinschreibung unterschieden."),
		app.Br(),
		app.Text("Komplexe Suchen (kurze Suchtexte im großen Suchgebiet) können die Suche manchmal abstürzen lassen (Seite bleibt hängen, Button reagiert nicht mehr) => in diesem Fall muss die Seite im Browser neu geladen werden."),
		app.Br(),
		app.Text("Über die Reiter kann das Suchgebiet ausgewählt werden, damit kann z.B: in das Gebiet Bautzen (BZ) gewechselt werden"),
		app.Br(),
		app.Br(),
		app.Text("Bsp: für die Suche:"),
		app.Br(),
		app.B().Text("Max Mustermann"),
		app.Text(" => Sucht nach Vor- und Nachnamen in dieser Kombination"),
		app.Br(),
		app.B().Text("Max Moritz Mustermann"),
		app.Text(" => bei mehreren Vornahmen, dabei ist die Reihenfolge wichtig. Hat diese Person im Index aber noch weitere Vornamen oder diese in anderer Reihenfolge, dann wird sie so wahrscheinlich nicht gefunden."),
		app.Br(),
		app.B().Text("Max *"),
		app.Text(" => Sucht nach Vornamen Max mit beliebigem Nachnamen"),
		app.Br(),
		app.B().Text("Max Moritz *"),
		app.Text(" => Sucht nach der Vornamenskombination Max Moritz mit beliebigem Nachnamen"),
		app.Br(),
		app.B().Text("* Schönberg"),
		app.Text(" => Sucht nach Nachname Schönberg mit beliebigen Vornamen. (Adelige Nachnamen z.B. 'von Schönberg' sind im Index ebenfalls als 'Schönberg' geführt und können (außer in Torgau) nicht von der nicht-adeligen Variante unterschieden werden)"),
		app.Br(),
		app.B().Text("* VitzthumEckstädt"),
		app.Text(" => als Bps. für den Nachnamen 'Vitzthum von Eckstädt' mit beliebigen Vornamen. (Mehrteilige Nachnamen wurden im Index (außer in Torgau) verkürzt und bestehen immer nur aus einem Wort)"),
		app.Br(),
		app.B().Text("* vonSchönberg"),
		app.Text(" => Nur im Gebiet Torgau sind adelige Nachnamen genauer erfasst, müssen aber ohne Leerzeichen geschrieben werden, damit die Suche damit funktioniert."),
		app.Br(),
		app.Br(),
		app.Text("Für spezielle Suchen kann rechts oben der Suchalgorithmus ausgewählt werden."),
		app.Br(),
		app.Text("Im Suchmodus "), app.B().Text("Klassisch"), app.Text(" wird buchstabengenau gesucht, nur damit kann gezielt nach Namensbestandteilen gesucht werden, z.B:"),
		app.Br(),
		app.B().Text("Napoleon *"),
		app.Text(" => Liefert Ergebnisse mit beliebigen Nachnamen, wo in den Vornamen das Wort 'Napoleon' enthalten ist, egal ob es weitere Vornamen davor oder danach gibt."),
		app.Br(),
		app.B().Text("* dorf"),
		app.Text(" => Liefert Ergebnisse mit beliebigen Vornamen, wo im Nachnamen 'dorf' enthalten ist"),
		app.Br(),
		app.B().Text("lieb Qu"),
		app.Text(" => Liefert Ergebnisse, wo im Vornamen 'lieb' und im Nachnamen 'Qu' enthalten ist"),
		app.Br(),
		app.B().Text("Thoma"),
		app.Text(" => Liefert Ergebnisse, wo 'Thoma' im Vor- oder Nachnamen enthalten ist"),
		app.Br(),
		app.Br(),
		app.Text("Weitere Informationen zu den Suchalgorithmen:"),
		app.Br(),
		app.A().Href("https://de.wikipedia.org/wiki/Levenshtein-Distanz").Text(" (Damerau-)Levenshtein-Distanz"),
		app.Br(),
		app.A().Href("https://srinivas-kulkarni.medium.com/jaro-winkler-vs-levenshtein-distance-2eab21832fd6").Text(" Jaro-Winkler vs. Levenshtein"),
		app.Br(),
		app.A().Href("https://de.wikipedia.org/wiki/K%C3%B6lner_Phonetik").Text(" Soundex (Kölner Phonetik)"),
		app.Br(),
		app.Br(),
		app.Text("Ergebnisse:"),
		app.Br(),
		app.Text("In den Ergebnissen führt der erste Link zur Datei mit den Roh-Daten, das kann helfen die genaue Position des Eintrages im Kirchenbuch zu finden."),
		app.Br(),
		app.Text("Manche Kirchenbücher sind auch nicht chronologisch bzw. die Daten im Buch verstreut."),
		app.Br(),
		app.Text("Es gibt auch immer mal wieder Fehler in den Daten, falsch erfasste Namen oder manchmal ist das Jahr um 1 verrutscht, etc."),
		app.Br(),
		app.Br(),
		app.Text("Der zweie Link führt direkt zum Kirchenbuch, meistens auch auf den richtigen Scan, das kann durch Fehler aber evtl. auch etwas abweichen."),
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
		app.H3().Body().Text(" v1.12 (Dezember 2025) latest updates:"),
		app.Label().Text("Einige Trauungen aus dem Gebiet Torgau hinzugefügt"),
		app.Br(),
		app.Label().Text("Verlinkungen zu Familysearch und Matricula verbessert"),
		app.Br(),
		app.Label().Text("kleine Fehlerkorrekturen"),
	).Attr("style", "font-family:verdana,sans-serif;font-size:8pt")
}

func getTabContent(h *searchComp, tab int, cbs map[int]map[string]app.HTMLDiv) app.HTMLDiv {
	return app.Div().Style("border", "1px solid #D3D3D3").Style("display", h.displayTab[tab]).Body(
		func() (row []app.UI) {
			for j := 0; j < rows; j++ {
				row = append(row, app.Tr().Body(
					func() (ele []app.UI) {
						for i := 0; i < cols; i++ {
							if val, ok := grid[tab][i][j]; ok {
								ele = append(ele, app.Td().Body(cbs[tab][val]))
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
	)
}

// after rendering
func (h *searchComp) OnMount(ctx app.Context) {
	start := time.Now()
	search.LoadData()
	dur := time.Since(start)

	//set defaults, slider and checkboxes
	h.slideValueMin = config.YearMin
	h.slideValueMax = config.YearMax
	h.activeTab = 1

	count := 0
	h.checked = make(map[int]map[string]bool)
	for tab := 0; tab < tabs; tab++ {
		h.checked[tab] = make(map[string]bool)
		for key, val := range search.Data.Marriages {
			if tab == 0 {
				count += len(val.Data) //count only once
			}
			if !isValid(tab, key) {
				continue
			}
			h.checked[tab][key] = true
		}
	}

	h.debug = fmt.Sprintf("%d Datensätze geladen in %s", count, dur.Round(time.Millisecond).String())
}
