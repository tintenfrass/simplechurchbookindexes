package ui

import (
	"fmt"
	"goapptest/config"
	"goapptest/search"
	"strings"
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func (h *searchComp) OnMount(ctx app.Context) {
	dur := search.Import()

	//Defaults
	h.slideValueMin = config.YearMin
	h.slideValueMax = config.YearMax

	count := 0
	h.checked = make(map[string]bool)
	for key, val := range search.Data.Marriages {
		h.checked[key] = true
		count += len(val.Data)
	}

	h.debug = fmt.Sprintf("%v Datensätze geladen in %s", count, dur.Round(time.Millisecond).String())
}

func replace(input string) (output string) {
	output = input
	output = strings.Replace(output, "Dresden", "DD", 1)
	output = strings.Replace(output, "Böhmische", "Böhm.", 1)
	output = strings.Replace(output, "Exulantengemeinde", "Exulanten", 1)
	output = strings.Replace(output, "Meißen", "MEI", 1)

	return
}

func (h *searchComp) Render() app.UI {
	//Checkboxes
	cbs := map[string]app.HTMLDiv{}
	for key, _ := range search.Data.Marriages {
		k := key //prevent Bug
		changeValue := func(ctx app.Context, e app.Event) {
			h.checked[k] = !h.checked[k]
		}
		cb := app.Input().Type("checkbox").Checked(h.checked[k]).OnChange(changeValue).Attr("style", "float:left")
		cbs[k] = app.Div().Body(
			cb,
			app.Label().Text(replace(k)).OnClick(changeValue),
			app.Br(),
			app.Label().Text(search.GetMinMax(k)).OnClick(changeValue),
		)
	}

	return app.Div().Body(
		app.Table().Body(
			app.Tr().Body(
				app.Div().Body(
					app.H2().Body(app.Text("Trau-Index Dresden-Meißen")),
					app.Label().Text("Jahr Min "),
					app.Input().Type("range").Attr("min", config.YearMin).Attr("max", config.YearMax).OnChange(h.ValueTo(&h.slideValueMin)).Value(h.slideValueMin).Attr("style", "width: 350px"),
					app.Label().Text(h.slideValueMin),
					app.Br(),
					app.Label().Text("Jahr Max"),
					app.Input().Type("range").Attr("min", config.YearMin).Attr("max", config.YearMax).OnChange(h.ValueTo(&h.slideValueMax)).Value(h.slideValueMax).Attr("style", "width: 350px"),
					app.Label().Text(h.slideValueMax),
					app.Br(),
					app.Br(),
					app.Br(),
					app.Table().Body(
						app.Tr().Body(
							app.Td().Body(cbs["Dörschnitz"]),
							app.Td(),
							app.Td(),
							app.Td(),
							app.Td().Body(cbs["Großdobritz"]),
						),
						app.Tr().Body(
							app.Td().Body(cbs["Lommatzsch"]),
							app.Td().Body(cbs["Zehren"]),
							app.Td().Body(cbs["Zadel"]),
							app.Td().Body(cbs["Gröbern"]),
							app.Td().Body(cbs["Oberau"]),
						),
						app.Tr().Body(
							app.Td(),
							app.Td(),
							app.Td(),
							app.Td(),
							app.Td().Body(cbs["Niederau"]),
						),
						app.Tr().Body(
							app.Td().Body(cbs["Leuben"]),
							app.Td(),
							app.Td().Body(cbs["Meißen St. Afra"]),
							app.Td().Body(cbs["Meißen Trinitatiskirche"]),
							app.Td().Body(cbs["Weinböhla"]),
						),
						app.Tr().Body(
							app.Td().Body(cbs["Planitz"]),
							app.Td(),
							app.Td().Body(cbs["Meißen Frauenkirche"]),
							app.Td().Body(cbs["Meißen Johanneskirche"]),
							app.Td(),
							app.Td().Body(cbs["Reichenberg"]),
							app.Td().Body(cbs["Wilschdorf"]),
						),
						app.Tr().Body(
							app.Td().Body(cbs["Ziegenhain"]),
							app.Td(),
							app.Td(),
							app.Td().Body(cbs["Brockwitz"]),
							app.Td().Body(cbs["Coswig"]),
							app.Td(),
							app.Td().Body(cbs["Klotzsche"]),
						),
						app.Tr().Body(
							app.Td().Body(cbs["Raußlitz"]),
							app.Td().Body(cbs["Krögis"]),
							app.Td(),
							app.Td().Body(cbs["Naustadt"]),
							app.Td().Body(cbs["Constappel"]),
							app.Td().Body(cbs["Kötzschenbroda"]),
							app.Td().Body(cbs["Dresden Dreikönigskirche"]),
						),
						app.Tr().Body(
							app.Td(),
							app.Td().Body(cbs["Miltitz"]),
							app.Td(),
							app.Td(),
							app.Td(),
							app.Td(),
							app.Td().Body(cbs["Dresden Hofkirche"]),
						),
						app.Tr().Body(
							app.Td().Body(cbs["Wendischbora"]),
							app.Td().Body(cbs["Heynitz"]),
							app.Td().Body(cbs["Taubenheim"]),
							app.Td().Body(cbs["Röhrsdorf"]),
							app.Td().Body(cbs["Weistropp"]),
							app.Td().Body(cbs["Kaditz"]),
							app.Td().Body(cbs["Dresden Sophienkirche"]),
						),
						app.Tr().Body(
							app.Td().Body(cbs["Rothschönberg"]),
							app.Td().Body(cbs["Burkhardswalde"]),
							app.Td().Body(cbs["Sora"]),
							app.Td(),
							app.Td(),
							app.Td(),
							app.Td().Body(cbs["Dresden Kreuzkirche"]),
						),
						app.Tr().Body(
							app.Td().Body(cbs["Deutschenbora"]),
							app.Td().Body(cbs["Tanneberg"]),
							app.Td().Body(cbs["Limbach"]),
							app.Td().Body(cbs["Wilsdruff"]),
							app.Td().Body(cbs["Unkersdorf"]),
							app.Td().Body(cbs["Briesnitz"]),
							app.Td().Body(cbs["Dresden Frauenkirche"]),
						),
						app.Tr().Body(
							app.Td().Body(cbs["Hirschfeld"]),
							app.Td().Body(cbs["Neukirchen"]),
							app.Td().Body(cbs["Blankenstein"]),
							app.Td().Body(cbs["Grumbach"]),
							app.Td().Body(cbs["Kesselsdorf"]),
							app.Td(),
							app.Td().Body(cbs["Dresden Friedrichstadt"]),
						),
						app.Tr().Body(
							app.Td(),
							app.Td(),
							app.Td().Body(cbs["Mohorn"]),
							app.Td().Body(cbs["Herzogswalde"]),
							app.Td(),
							app.Td().Body(cbs["Pesterwitz"]),
							app.Td().Body(cbs["Dresden Annenkirche"]),
						),
						app.Tr().Body(
							app.Td(),
							app.Td(),
							app.Td(),
							app.Td().Body(cbs["Fördergersdorf"]),
							app.Td().Body(cbs["Tharandt"]),
							app.Td().Body(cbs["Döhlen"]),
							app.Td().Body(cbs["Dresden Böhmische Exulantengemeinde"]),
						),
						app.Tr().Body(
							app.Td().Body(app.Button().Text("Alles").OnClick(h.all).Attr("style", "width: 70px")),
							app.Td().Body(app.Button().Text("Nichts").OnClick(h.nothing).Attr("style", "width: 70px")),
							app.Td(),
							app.Td(),
							app.Td().Body(cbs["Somsdorf"]),
							app.Td(),
							app.Td().Body(cbs["Plauen"]),
						),
					),
					app.P().Body(
						app.Input().Type("text").Placeholder("Vorname Nachname").AutoFocus(true).OnChange(h.ValueTo(&h.searchValue)).Attr("style", "width: 200px"),
						app.Text(" "),
						app.Button().Text("Search").OnClick(h.onClick).Attr("style", "width: 100px"),
						app.Text(" "),
						app.Label().Text(h.debug),
					),
					app.H3().Body().Text("Hinweise:"),
					app.Text("Für die meisten Kirchengemeinden wurden die Daten bisher nur bis 1788 erfasst."),
					app.Br(),
					app.Br(),
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
					app.Text("* (oder ?) als Platzhalter:"),
					app.Br(),
					app.Text("Max * => Sucht nach Vornamen Max"),
					app.Br(),
					app.Text("* Mustermann => Sucht nach Nachname Mustermann"),
					app.Br(),
					app.Br(),
					app.Text("Manchmal kommt es vor, dass man einen gefundenen Treffer im Kirchenbuch selber nicht findet."),
					app.Br(),
					app.Text("Dann hilft es in die erfassten Daten selber zu gucken, um die genaue Position im Kirchenbuch zu ermitteln."),
					app.Br(),
					app.A().Href("https://github.com/tintenfrass/simplechurchbookindexes/tree/main/sachsen").Text("https://github.com/tintenfrass/simplechurchbookindexes/tree/main/sachsen"),
					app.Br(),
					app.Text("Manche Kirchenbücher sind nicht chronologisch bzw. die Daten im Buch verstreut."),
					app.Br(),
					app.Text("Es gibt auch immer mal wieder Fehler in den Daten, falsch erfasste Namen oder manchmal ist das Jahr um 1 verrutscht, etc."),
					app.Br(),
					app.Br(),
					app.H3().Body().Text("Mehr Informationen zum Projekt:"),
					app.A().Href("https://github.com/tintenfrass/simplechurchbookindexes").Text("https://github.com/tintenfrass/simplechurchbookindexes"),
					app.Br(),
					app.H3().Body().Text(" v1.1 (August 2023) latest updates:"),
					app.Label().Text("Trauungen 1789-1799"),
				),
				app.Td().Body(app.Textarea().Text(h.result).Attr("readonly", true).Attr("warp", "hard").Attr("cols", 120).Attr("rows", 50)),
			),
		),
	).Attr("style", "font-family:verdana,sans-serif;font-size:8pt")
}

func (h *searchComp) onClick(ctx app.Context, e app.Event) {
	start := time.Now()
	h.result = search.FindMarriage(h.searchValue, h.slideValueMin, h.slideValueMax, h.checked)
	dur := time.Since(start)
	h.debug = fmt.Sprintf("Suchzeit: %s", dur.Round(time.Millisecond).String())
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
