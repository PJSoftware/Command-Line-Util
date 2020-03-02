# Ripping Utilities

For convenience, I recently started ripping my DVDs to digital videos, using Handbrake to perform the rip, and MetaX to add appropriate metadata so they can be added to iTunes and not look out of place.

For movies this is fairly straightforward. However, for TV shows, juggling all the output files and simplifying running them through MetaX took a bit of thought. Ultimately I came up with a couple of utilities to help reduce the pain of ripping a full season (or more) of a show.

These were originally written in Perl, but it seemed worthwhile rewriting them in Go and making them available for anyone else who might want them.

## CreateFolders

The first step is to create folders into which each disk of your series can be ripped.

**CreateFolders.exe** prompts for three pieces of information: series name, number of seasons, and number of disks per season. It will then create a folder tree structure containing one folder per disk per season, under the parent series folder. Each disk in the series should be ripped into the appropriate folder.

## SortByFolder

Once all disks of a season have been ripped into the folder structure described above, **SortByFolder.exe** will rename them all to the format "S#E##". This makes labelling them via MetaX much easier, since it recognises this naming convention.
