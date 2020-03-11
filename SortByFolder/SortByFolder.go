package main

import (
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
	folders.InitFileRegexp()

	fc := folders.Init()
	series := fc.ChooseSeries()
	if series == nil {
		return
	}

	series.Rename()
}
