package finder

import (
	"sync"
	"regexp"
)

type Finder struct {
	mu  sync.Mutex
	Res []string

	excludeMap map [string] struct {}
	Once sync.Once

	re *regexp.Regexp


}
