package ui

import (
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func StartLocal() {
	app.Route("/", &searchComp{})
	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "FuzzySearch",
		Description: "FuzzySearch Online",
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

func StartGithub() {
	app.Route("/", &searchComp{})
	app.RunWhenOnBrowser()

	err := app.GenerateStaticWebsite(".", &app.Handler{
		Name:        "FuzzySearch",
		Description: "FuzzySearch Online",
		Resources:   app.GitHubPages("simplechurchbookindexes"),
	})

	if err != nil {
		log.Fatal(err)
	}
}
