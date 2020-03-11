package folders

import (
	"io/ioutil"
	"log"
)

// Season folder data
type Season struct {
	name  string
	count int
	disks []Disk
}

func (s *Season) scan(parent string) {
	files, err := ioutil.ReadDir(parent + s.name)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			disk := new(Disk)
			disk.name = file.Name()
			disk.scan(parent+s.name+"/", s.name)
			if disk.count != 0 {
				s.disks = append(s.disks, *disk)
				s.count += disk.count
			}
		}
	}

}
