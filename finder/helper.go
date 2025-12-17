package finder

import "regexp"

func (f *Finder) ToMap(excludeDir []string) {
	f.Once.Do(func() {
		f.excludeMap = make(map [string]struct {}, len(excludeDir))
		for _, v := range excludeDir {
			f.excludeMap[v] = struct{}{}
		}
	})
}

func (f *Finder) isExcluded(dir string) bool {
	_, ok := f.excludeMap[dir]
	return ok
}

func (f *Finder) SetRegex(pattern string) error {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return err 
	}

	f.re = re 

	return nil 
}

func (f *Finder) appendVal(path string) {
	f.mu.Lock()
	f.Res = append(f.Res, path)
	f.mu.Unlock()
}
