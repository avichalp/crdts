package main

import (
	"fmt"
	"math/rand"
)

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

// state base CRDT interface:
// udapte, query, compare, merge

type PNCounter struct {
	pCounter *GCounter
	nCounter *GCounter
}

func NewPNCounter() *PNCounter {
	return &PNCounter{
		pCounter: NewGCounter(),
		nCounter: NewGCounter(),
	}
}

func (pn *PNCounter) Increment() {
	pn.pCounter.Increment()
}

func (pn *PNCounter) Decrement() {
	pn.nCounter.Increment()
}

func (pn *PNCounter) Value() int {
	return pn.pCounter.Value() - pn.nCounter.Value()
}

func (pn *PNCounter) Merge(pn2 *PNCounter) {
	pn.pCounter.Merge(pn2.pCounter)
	pn.nCounter.Merge(pn2.nCounter)
}

func main() {
	fmt.Println("hello world")
	gcounter1, gcounter2 := NewGCounter(), NewGCounter()

	gcounter1.Increment()
	gcounter2.Increment()

	gcounter1.Merge(gcounter2)

	fmt.Println(gcounter1.Value())

	gcounter2.Merge(gcounter1)
	fmt.Println(gcounter2.Value())

	fmt.Println("PN Counter examples")

	pncounter1, pncounter2 := NewPNCounter(), NewPNCounter()

	pncounter1.Increment()
	pncounter2.Increment()

	pncounter1.Merge(pncounter2)

	fmt.Println("PN Counter", pncounter1.Value())

}
