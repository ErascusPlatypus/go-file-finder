package finder

import (
	"sync"
	"path/filepath"
	"os"
)

func (f *Finder) processDir(dir string, fileName string, dirs chan<- string, wg *sync.WaitGroup) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return
	}
	
	for _, e := range entries {
		path := filepath.Join(dir, e.Name())
		
		if e.IsDir() {
			wg.Add(1)
			go func (p string)  {
				dirs <- p 
			} (path)
		} else if e.Name() == fileName {
			f.mu.Lock()
			f.Res = append(f.Res, path)
			f.mu.Unlock()
		}
	}
}


func (f *Finder) JobFinder(root string, fileName string) {
	var dirWG sync.WaitGroup
	var workerWG sync.WaitGroup

	dirs := make(chan string, 100)

	workerWG.Add(20)
	for i := 0; i < 20; i++ {
		go func() {
			defer workerWG.Done()
			for dir := range dirs {
				f.processDir(dir, fileName, dirs, &dirWG)
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
