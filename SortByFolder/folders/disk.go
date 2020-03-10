package folders

import (
	"io/ioutil"
	"log"
	"regexp"
)

// Disk folder data
type Disk struct {
	name   string
	count  int
	videos []string
}

func (d *Disk) scan(parent string) {
	files, err := ioutil.ReadDir(parent + d.name)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			if ok, _ := regexp.MatchString(".*[.]mp4", file.Name()); ok {
				path := parent + d.name + "/" + file.Name()
				d.videos = append(d.videos, path)
				d.count++
			}
		}
	}

}
