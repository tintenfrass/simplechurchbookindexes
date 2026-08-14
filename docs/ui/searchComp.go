package ui

import (
	"fmt"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/maxence-charriere/go-app/v10/pkg/app"

	"onlinefuzzysearch/search"
)

type searchComp struct {
	app.Compo
	searchValueV  string
	searchValueN  string
	slideValueMin int
	slideValueMax int
	results       []app.UI
	debug         string
	checked       map[int]map[string]bool
	algo          int
	//Tab stuff
	activeTab  int
	displayTab []string
}

func (h *searchComp) search(ctx app.Context, e app.Event) {
	const linkPrefix = "https://github.com/tintenfrass/simplechurchbookindexes/blob/main/sachsen/"
	start := time.Now()

	h.results = []app.UI{}
	full := -1

	boxes := make(map[int][]app.UI, 10)
	for i := 0; i < 8; i++ {
		boxes[i] = []app.UI{}
	}

	places := make(map[string]struct{})
	exact := h.activeTab != 0

	// Filter places
	for place, val := range h.checked[h.activeTab] {
		if !val {
			continue
		}
		if exact && strings.Contains(place, "*") {
			continue
		}
		if !exact && !strings.Contains(place, "*") {
			continue
		}
		places[place] = struct{}{}
	}

	resultList, debug := search.FindMarriage(h.searchValueV, h.searchValueN, h.slideValueMin, h.slideValueMax, places, h.algo, exact)
	for _, res := range resultList {
		if res.Dis > search.MaxDistance {
			break
		}
		full = res.Dis

		if strings.Contains(res.Line, "Zu viele Ergebnisse") {
			boxes[res.Dis] = append(boxes[res.Dis], app.Tr().Body(
				app.Td().Body(app.Label().Text("»»»").Style("font-weight", "bold").Attr("style", "color: "+getColor(res.Dis))),
				app.Td().Body(app.Text(res.Line)),
			))
			continue
		}

		src := getSource(res.Link)
		if res.Page != 0 && src == "Archion" {
			res.Link += "?pageId=" + strconv.Itoa(res.Page)
		} else if res.Page != 0 && src == "Matricula" {
			res.Link += "?pg=" + strconv.Itoa(res.Page)
		} else if res.Page != 0 && src == "Familysearch" {
			res.Link += "&i=" + strconv.Itoa(res.Page)
		}

		boxes[res.Dis] = append(boxes[res.Dis], app.Tr().Body(
			app.Td().Body(app.Label().Text("»»»").Style("font-weight", "bold").Attr("style", "color: "+getColor(res.Dis))),
			app.Td().Body(app.Text(strconv.Itoa(res.Year)+" "+res.Line)),
			app.Td().Body(app.Label().Style("margin", "16px")),
			app.Td().Body(app.Text(replaceKK(path.Dir(res.Source)))).Style("color", "dimgrey;font-size:7pt"),
			app.Td().Body(app.A().Href(linkPrefix+res.Source).Text(path.Base(res.Source))),
			app.Td().Body(app.Label().Style("margin", "16px")),
			app.Td().Body(app.A().Href(res.Link).Text(src)),
		))
	}

	for i := 0; i < search.MaxDistance+1; i++ {
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

func (h *searchComp) plusminus(value bool, k string) {
	posi, posj := getPos(h.activeTab, k)

	for r := 0; r < plusminusrange; r++ {
		next := true
		for i := 0; i < cols+r; i++ {
			if i >= posi-r && i <= posi+r {
				for j := 0; j < rows+r; j++ {
					if j >= posj-r && j <= posj+r {
						if val, ok := grid[h.activeTab][i][j]; ok {
							if h.checked[h.activeTab][val] != value {
								h.checked[h.activeTab][val] = value
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

// Button show all
func (h *searchComp) all(ctx app.Context, e app.Event) {
	for _, place := range getPlacesList() {
		if !isValid(h.activeTab, place) {
			continue
		}
		h.checked[h.activeTab][place] = true
	}
}

// Button show nothing
func (h *searchComp) nothing(ctx app.Context, e app.Event) {
	for _, place := range getPlacesList() {
		if !isValid(h.activeTab, place) {
			continue
		}
		h.checked[h.activeTab][place] = false
	}
}

// show tab
func (h *searchComp) showTab(nr int) {
	for key, _ := range h.displayTab {
		if nr == key {
			h.displayTab[key] = "block"
		} else {
			h.displayTab[key] = "none"
		}
	}
	h.activeTab = nr
}

// determines if location is valid for a given tab
func isValid(tab int, place string) bool {
	row, ok := grid[tab]
	if !ok {
		return false
	}

	for _, col := range row {
		for _, name := range col {
			if name == place {
				return true
			}
			if tab == 0 {
				prefixGrid, ast := strings.CutSuffix(name, "/*")
				if ast && strings.HasPrefix(place, prefixGrid) {
					return true
				}
			}
		}
	}

	return false
}

func getPlacesList() []string {
	var placesList []string
	for place, _ := range search.Data.Marriages {
		placesList = append(placesList, place)
	}
	placesList = append(placesList, "auerbach/*")
	placesList = append(placesList, "bad-liebenwerda/*")
	placesList = append(placesList, "bautzen/*")
	placesList = append(placesList, "dippoldiswalde/*")
	placesList = append(placesList, "dresden/*")
	placesList = append(placesList, "freiberg/*")
	placesList = append(placesList, "meissen/*")
	placesList = append(placesList, "torgau-delitzsch/*")

	return placesList
}
