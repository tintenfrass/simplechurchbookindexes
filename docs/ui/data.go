package ui

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type searchComp struct {
	app.Compo
	searchValue   string
	slideValueMin int
	slideValueMax int
	results       []app.UI
	debug         string
	checked       map[string]bool
	algo          int
}
