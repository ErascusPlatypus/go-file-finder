package finder

import (
	"os"
	"path/filepath"
	"sync"
)

var sem = make(chan	struct {}, 20)

func (f *Finder) SemFinder(dir, fileName string, regexFlag bool, wg *sync.WaitGroup) {
	defer wg.Done()

	sem <- struct{}{}
	defer func() { <- sem }()
	entries, err := os.ReadDir(dir)

	// ch <- val    === value flows into ch 
	// x := <- ch   ===== value flows out of ch 

	if err != nil {
		return 
	} 
	
	for _, e := range entries {
		path := filepath.Join(dir, e.Name())
		normalSearch :=  (e.Name() == fileName)
		regexSearch := (f.re != nil && f.re.MatchString(e.Name()))

		if e.IsDir() && !f.isExcluded(e.Name()) {			
			wg.Add(1)
			go f.SemFinder(path, fileName, regexFlag, wg) 
		} else if regexFlag && regexSearch {
			f.appendVal(path)
		} else if !regexFlag && normalSearch {
			f.appendVal(path)
		}
	}
}
