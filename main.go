package main

import (
	"math/rand"
	"strconv"
	"sync"
	"time"

	cmap "github.com/orcaman/concurrent-map"
)

const count = 2_000_000
const num = 200

var sm = syncMap{}
var rwm = rwmMap{}
var mm = mMap{}
var cm = cMap{}

func init() {
	rand.Seed(rand.Int63())

	sm.Init()
	go sm.SetRandom()

	rwm.Init()
	go rwm.SetRandom()

	mm.Init()
	go mm.SetRandom()

	cm.Init()
	go cm.SetRandom()

}

type syncMap struct {
	m sync.Map
}

func (sm *syncMap) Init() {
	for i := 0; i < count; i++ {
		sm.m.Store(i, i)
	}
}

func (sm *syncMap) GetRandom() {
	c := rand.Intn(count)
	v, _ := sm.m.Load(c)
	i := v.(int)
	_ = i
}

func (sm *syncMap) SetRandom() {
	for i := 0; i < num; i++ {
		go func() {
			for {
				c := rand.Intn(count)
				sm.m.Store(c, c)
				time.Sleep(time.Millisecond * time.Duration(c))
			}
		}()
	}
}

type rwmMap struct {
	rwm sync.RWMutex
	m   map[int]int
}

func (rwm *rwmMap) Init() {
	rwm.m = make(map[int]int, count)
	for i := 0; i < count; i++ {
		rwm.m[i] = i
	}
}

func (rwm *rwmMap) GetRandom() {
	c := rand.Intn(count)
	rwm.rwm.RLock()
	defer rwm.rwm.RUnlock()
	_, _ = rwm.m[c]

}

func (rwm *rwmMap) SetRandom() {
	for i := 0; i < num; i++ {
		go func() {
			for {
				c := rand.Intn(count)
				rwm.rwm.Lock()
				rwm.m[c] = c
				rwm.rwm.Unlock()
				time.Sleep(time.Nanosecond * time.Duration(c))
			}
		}()
	}
}

type mMap struct {
	mu sync.Mutex
	m  map[int]int
}

func (mm *mMap) Init() {
	mm.m = make(map[int]int, count)
	for i := 0; i < count; i++ {
		mm.m[i] = i
	}
}

func (mm *mMap) GetRandom() {
	c := rand.Intn(count)
	mm.mu.Lock()
	defer mm.mu.Unlock()
	_, _ = mm.m[c]

}

func (mm *mMap) SetRandom() {
	for i := 0; i < num; i++ {
		go func() {
			for {
				c := rand.Intn(count)
				mm.mu.Lock()
				mm.m[c] = c
				mm.mu.Unlock()
				time.Sleep(time.Nanosecond * time.Duration(c))
			}
		}()
	}

}

type cMap struct {
	m cmap.ConcurrentMap
}

func (cm *cMap) Init() {
	cm.m = cmap.New()
	for i := 0; i < count; i++ {
		cm.m.Set(strconv.Itoa(i), i)
	}
}

func (cm *cMap) GetRandom() {
	c := rand.Intn(count)
	v, _ := cm.m.Get(strconv.Itoa(c))
	i := v.(int)
	_ = i
}

func (cm *cMap) SetRandom() {
	for i := 0; i < num; i++ {
		go func() {
			for {
				c := rand.Intn(count)
				cm.m.Set(strconv.Itoa(c), c)
				time.Sleep(time.Millisecond * time.Duration(c))
			}
		}()
	}
}

func main() {
}
