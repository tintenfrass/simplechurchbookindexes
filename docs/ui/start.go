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
		Name:        "Trauindex Sachsen",
		Description: "Onlinesuche zum Trauindex",
		Title:       "Trauindex Sachsen",
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

func StartGithub() {
	app.Route("/", &searchComp{})
	app.RunWhenOnBrowser()

	err := app.GenerateStaticWebsite(".", &app.Handler{
		Name:        "Trauindex Sachsen",
		Description: "Onlinesuche zum Trauindex",
		Title:       "Trauindex Sachsen",
		Resources:   app.GitHubPages("simplechurchbookindexes"),
		RawHeaders: []string{
			`<!-- Google tag (gtag.js) -->
			<script async src="https://www.googletagmanager.com/gtag/js?id=G-S2D2V6SHXE"></script>
			<script>
			window.dataLayer = window.dataLayer || [];
			function gtag(){dataLayer.push(arguments);}
			gtag('js', new Date());

			gtag('config', 'G-S2D2V6SHXE');
			</script>`,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
