package finder

import (
	"sync"
)

type Finder struct {
	mu  sync.Mutex
	Res []string

	excludeMap map [string] struct {}
	Once sync.Once
}
