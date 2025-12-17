package finder

import (
	"sync"
	"path/filepath"
	"os"
)

func (f *Finder) processDir(dir, fileName string, dirs chan<- string, regexFlag bool, wg *sync.WaitGroup) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return
	}
	
	for _, e := range entries {
		path := filepath.Join(dir, e.Name())

		normalSearch := (e.Name() == fileName) 
		regexSearch := (f.re != nil && f.re.MatchString(e.Name()))
		
		if e.IsDir() && !f.isExcluded(e.Name()) {
			wg.Add(1)

			go func (p string)  {
				dirs <- p 
			} (path)

		} else if regexFlag && regexSearch {
			f.appendVal(path)
		} else if !regexFlag && normalSearch {
			f.appendVal(path)
		}
	}
}


func (f *Finder) JobFinder(root, fileName string, regexFlag bool) {
	var dirWG sync.WaitGroup
	var workerWG sync.WaitGroup

	dirs := make(chan string, 100)

	workerWG.Add(20)
	for i := 0; i < 20; i++ {
		go func() {
			defer workerWG.Done()
			for dir := range dirs {
				f.processDir(dir, fileName, dirs, regexFlag, &dirWG)
				dirWG.Done()
			}
		}()
	}

	dirWG.Add(1)
	dirs <- root

	go func() {
		dirWG.Wait()
		close(dirs)
	}()

	workerWG.Wait()
}
