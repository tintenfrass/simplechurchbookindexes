package main

import (
	"fmt"
	"indexfuzzysearch/config"
	"indexfuzzysearch/search"
	"indexfuzzysearch/ui"
	"runtime"
)

func main() {
	fmt.Println(fmt.Sprintf("Index Fuzzy Search %s (%s)", ui.Version, runtime.GOARCH))

	config.Load()
	search.Update()
	search.ImportFromGit()
	ui.BuildAndRun()
}
