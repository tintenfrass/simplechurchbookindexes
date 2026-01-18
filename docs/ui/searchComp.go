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
	searchValue   string
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

	resultList, debug := search.FindMarriage(h.searchValue, h.slideValueMin, h.slideValueMax, h.checked[h.activeTab], h.algo)
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

	for r := 0; r < 42; r++ {
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

func isValid(tab int, prefix string) bool {
	switch {
	case strings.HasPrefix(prefix, "auerbach/"):
		if tab == 0 {
			return true
		}
	case strings.HasPrefix(prefix, "torgau-delitzsch/"):
		fallthrough
	case strings.HasPrefix(prefix, "bad-liebenwerda/"):
		if tab == 1 {
			return true
		}
	case strings.HasPrefix(prefix, "dresden/"):
		fallthrough
	case strings.HasPrefix(prefix, "meissen/"):
		fallthrough
	case strings.HasPrefix(prefix, "freiberg/"):
		fallthrough
	case strings.HasPrefix(prefix, "dippoldiswalde/"):
		if tab == 2 {
			return true
		}
	case strings.HasPrefix(prefix, "bautzen/"):
		if tab == 3 {
			return true
		}
	}

	return false
}

// Button show all
func (h *searchComp) all(ctx app.Context, e app.Event) {
	for key, _ := range search.Data.Marriages {
		if !isValid(h.activeTab, key) {
			continue
		}
		h.checked[h.activeTab][key] = true
	}
}

// Button show nothing
func (h *searchComp) nothing(ctx app.Context, e app.Event) {
	for key, _ := range search.Data.Marriages {
		if !isValid(h.activeTab, key) {
			continue
		}
		h.checked[h.activeTab][key] = false
	}
}

func (h *searchComp) tab0(ctx app.Context, e app.Event) {
	h.showTab(0)
}

func (h *searchComp) tab1(ctx app.Context, e app.Event) {
	h.showTab(1)
}

func (h *searchComp) tab2(ctx app.Context, e app.Event) {
	h.showTab(2)
}

func (h *searchComp) tab3(ctx app.Context, e app.Event) {
	h.showTab(3)
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
