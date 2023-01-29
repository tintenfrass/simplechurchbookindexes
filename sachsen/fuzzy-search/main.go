package main

import (
	"fmt"
	"indexfuzzysearch/config"
	"indexfuzzysearch/search"
	"indexfuzzysearch/ui"
	"runtime"
)

func main() {
	fmt.Println("Index Fuzzy Search v1.0", runtime.GOARCH)

	config.Load()
	search.ImportFromGit()
	ui.BuildAndRun()
}
