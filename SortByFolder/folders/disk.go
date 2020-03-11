package folders

import (
	"io/ioutil"
	"log"
	"strconv"
)

var fr *FileRegexp

const frPattern string = `.*-(\d+)[.](mp4|m4v)`

// Disk folder data
type Disk struct {
	name   string
	count  int
	videos []Video
}

// InitFileRegexp sets up the fr used in scan
func InitFileRegexp() {
	fr = NewFileRegexp(frPattern)
}

func (d *Disk) scan(parent string, season string) {
	files, err := ioutil.ReadDir(parent + d.name)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			if fr.Match(file.Name()) {
				path := parent + d.name + "/" + file.Name()
				video := new(Video)
				video.disk = d.name
				video.season = season
				video.name = path
				match := fr.Extract(file.Name())
				video.track, _ = strconv.Atoi(match[0])
				video.ext = match[1]
				d.videos = append(d.videos, *video)
				d.count++
			}
		}
	}
}
