package crdts

import "math/rand"

// state-based grow only counter
type GCounter struct {
	id      int
	counter map[int]int
}

func NewGCounter() *GCounter {
	return &GCounter{
		id:      rand.Int(),
		counter: make(map[int]int),
	}
}

func (g *GCounter) Increment() {
	g.counter[g.id] += 1
}

func (g *GCounter) Value() (total int) {
	for _, val := range g.counter {
		total += val
	}
	return
}

// Merge combies the counter values across multiple replicas.
// The property of idempotency is preserved here across
// multiple merges as when no state is changed across any replicas,
// the result should be exactly the same everytime
func (g *GCounter) Merge(c *GCounter) {
	for id, val := range c.counter {
		if v, ok := g.counter[id]; !ok || v < val {
			g.counter[id] = val
		}
	}
}

func (g *GCounter) Compare(c *GCounter) bool {
	return g.counter[g.id] <= c.counter[c.id]
}
