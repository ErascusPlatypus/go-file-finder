package finder

import (
	"sync"
)

type Finder struct {
	mu  sync.Mutex
	Res []string
}
