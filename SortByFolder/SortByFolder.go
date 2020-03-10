package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"./folders"
)

// SortByFolder is intended to be run from the "series" parent folder (the same
// folder CreateFolders will be run from.)
//
// Therefore, there will be three levels of folders to parse:
//   [series_name]/[s#]/[d#]/[mp4_files]
//
// We want some level of interactivity. It will be necessary to:
//  1. Determine which series folders have sortable content
//  2. Prompt for series to sort
//  3. Do the sorting
func main() {
	fc := folders.Init()
	series := fc.ChooseSeries()
	fmt.Println(series.Name())
}

func getString(prompt string, def string) string {
	reader := bufio.NewReader(os.Stdin)
	if def != "" {
		prompt = prompt + " [" + def + "]"
	}
	fmt.Print(prompt + ": ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r", "", -1)
	text = strings.Replace(text, "\n", "", -1)
	if text == "" {
		text = def
	}
	return text
}
