package ui

import (
	"fmt"
	"time"

	"github.com/maxence-charriere/go-app/v10/pkg/app"

	"onlinefuzzysearch/config"
	"onlinefuzzysearch/search"
)

// Render the full ui
func (h *searchComp) Render() app.UI {
	catholic := getCatholic()
	intTabs(h)
	initGrid()

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
			if _, exists := catholic[k]; exists {
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
						app.H2().Body(app.Text("Onlinesuche im Trauindex Sachsen")),
						app.Label().Text("Bisher erfasst: Dresden Meißen Freiberg Dippoldiswalde Auerbach (+Teile von Torgau und Bautzen)"),
						app.Br(),
						app.Br(),
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
		createTabs(h, cbs),
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
		app.Text(" => Sucht nach Nachname Schönberg mit beliebigen Vornamen. (Adelige Nachnamen z.B. 'von Schönberg' sind im Index ebenfalls als 'Schönberg' geführt und können (außer in Torgau und Auerbach) nicht von der nicht-adeligen Variante unterschieden werden)"),
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
		app.H3().Body().Text(" v1.12 (März 2026) latest updates:"),
		app.Label().Text("Trauungen aus dem Gebiet Auerbach hinzugefügt"),
		app.Br(),
		app.Label().Text("kleine Fehlerkorrekturen"),
	).Attr("style", "font-family:verdana,sans-serif;font-size:8pt")
}

// after rendering
func (h *searchComp) OnMount(ctx app.Context) {
	start := time.Now()
	search.LoadData()
	dur := time.Since(start)

	//set defaults, slider and checkboxes
	h.slideValueMin = config.YearMin
	h.slideValueMax = config.YearMax
	h.activeTab = defaultTab

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
