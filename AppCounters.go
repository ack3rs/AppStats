package AppStats

import (
	"sync"
)

type Counter struct {
	Stats map[string]int
	Mux   sync.Mutex
}

var Stats = Counter{Stats: make(map[string]int)}

func (c *Counter) IncStat(StatsticName string) {
	c.Mux.Lock()
	c.Stats[StatsticName] = c.Stats[StatsticName] + 1
	c.Mux.Unlock()
}

func (c *Counter) SetStat(StatsticName string, StatisticValue int) {
	c.Mux.Lock()
	c.Stats[StatsticName] = StatisticValue
	c.Mux.Unlock()
}

func (c *Counter) GetStat(StatsticName string) int {
	c.Mux.Lock()
	defer c.Mux.Unlock()
	return c.Stats[StatsticName]
}

func (c *Counter) GetStatKeys() []string {
	c.Mux.Lock()
	defer c.Mux.Unlock()
	out := []string{}
	for k, _ := range c.Stats {
		out = append(out, k)
	}
	return out
}
