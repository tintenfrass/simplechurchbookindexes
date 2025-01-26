package ui

import (
	"fmt"
	"onlinefuzzysearch/search"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
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

	data, debug := search.FindMarriage(h.searchValue, h.slideValueMin, h.slideValueMax, h.checked[h.activeTab], h.algo)
	for _, res := range data {
		parts := strings.Split(res, "#")
		dis, _ := strconv.Atoi(parts[2])
		if dis > 7 {
			break
		}
		full = dis

		if strings.Contains(res, "0 Zu viele Ergebnisse") {
			boxes[dis] = append(boxes[dis], app.Tr().Body(
				app.Td().Body(app.Label().Text("»»»").Style("font-weight", "bold").Attr("style", "color: "+getColor(dis))),
				app.Td().Body(app.Text(parts[0][2:])),
			))
			continue
		}

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

//Button show all
func (h *searchComp) all(ctx app.Context, e app.Event) {
	for key, _ := range search.Data.Marriages {
		h.checked[h.activeTab][key] = true
	}
}

//Button show nothing
func (h *searchComp) nothing(ctx app.Context, e app.Event) {
	for key, _ := range search.Data.Marriages {
		h.checked[h.activeTab][key] = false
	}
}

func (h *searchComp) tab0(ctx app.Context, e app.Event) {
	h.showTab(0)
}

func (h *searchComp) tab1(ctx app.Context, e app.Event) {
	h.showTab(1)
}

//show tab
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
