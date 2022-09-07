package crdts

import (
	"math/rand"
)

// state-based grow only counter
type GCounter struct {
	id      int
	counter map[int]int
}

// NewGCounter returns a new instance of a GCounter
// The new replica can be uniquely identified using
// unique id.
func NewGCounter() *GCounter {
	return &GCounter{
		id:      rand.Int(),
		counter: make(map[int]int),
	}
}

// Increment adds 1 to the current counter value
func (g *GCounter) Increment() {
	g.counter[g.id] += 1
}

// Value returns the counter value. It is the sum
// of all counter values of every replica as recorded
// in this replica
func (g *GCounter) Value() (total int) {
	for _, v := range g.counter {
		total += v
	}
	return
}

// Merge combies the counter values across multiple replicas.
// The property of idempotency is preserved here across
// multiple merges as when no state is changed across any replicas,
// the result should be exactly the same everytime
func (g *GCounter) Merge(c *GCounter) {
	for id, u := range c.counter {
		v, ok := g.counter[id]
		if !ok || v < u {
			g.counter[id] = u
		}
	}
}
