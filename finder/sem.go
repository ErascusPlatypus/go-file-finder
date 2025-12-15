package finder

import (
	"os"
	"path/filepath"
	"sync"
)

var sem = make(chan	struct {}, 20)

func (f *Finder) SemFinder(dir string, fileName string, wg *sync.WaitGroup) {
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

		if e.IsDir() {
			wg.Add(1)
			go f.SemFinder(path, fileName, wg) 
		} else if e.Name() == fileName {
			f.mu.Lock()
			f.Res = append(f.Res, path)
			f.mu.Unlock()
		}
	}
}
