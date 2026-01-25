package ui

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

const (
	tabs       = 4
	defaultTab = 2
)

var tabConfig = []struct {
	Short string
	Title string
}{
	{Short: "AU", Title: "Auerbach"},
	{Short: "TO", Title: "Torgau"},
	{Short: "DD-MEI-FG-DW", Title: "Dresden-Meißen-Freiberg-Dippoldiswalde"},
	{Short: "BZ", Title: "Bautzen"},
}

func intTabs(h *searchComp) {
	if len(h.displayTab) == 0 {
		for i := 0; i < tabs; i++ {
			if i == defaultTab {
				h.displayTab = append(h.displayTab, "block")
			} else {
				h.displayTab = append(h.displayTab, "none")
			}
		}
	}
}

func createTabs(h *searchComp, cbs map[int]map[string]app.HTMLDiv) app.UI {
	var elems []app.UI
	contents := make([]app.UI, 0, tabs)
	for i := 0; i < tabs; i++ {
		elems = append(elems,
			app.Td().Body(app.Button().
				Disabled(h.displayTab[i] == "block").
				Text(tabConfig[i].Short).
				Title(tabConfig[i].Title).
				OnClick(func(ctx app.Context, e app.Event) { h.showTab(i) }),
			),
		)
		contents = append(contents, getTabContent(h, i, cbs))
	}
	elems = append(elems,
		app.Table().Body(contents...),
	)

	return app.Div().Body(elems...)
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
