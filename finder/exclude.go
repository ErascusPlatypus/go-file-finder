package finder

func (f *Finder) ToMap(excludeDir []string) {
	f.Once.Do(func() {
		mp := make(map [string]struct {}, len(excludeDir))
		for _, v := range excludeDir {
			mp[v] = struct{}{}
		}
	})
}

func (f *Finder) isExcluded(dir string) bool {
	_, ok := f.excludeMap[dir]
	return ok
}