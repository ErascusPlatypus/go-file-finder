package finder

import (
	"os"
	"path/filepath"

)

func (f *Finder) BasicFinder(dir string, fileName string) {
	entries, err := os.ReadDir(dir)

	if err != nil {
		return 
	}

	for _, e := range entries {
		path := filepath.Join(dir, e.Name())

		if e.IsDir() {
			f.BasicFinder(path, fileName) 
		} else if e.Name() == fileName {
			f.Res = append(f.Res, path)
		}
	}
}
