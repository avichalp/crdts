package main

import (
	"encoding/json"
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

type mainSet map[interface{}]struct{}

type GSet struct {
	mainSet mainSet
}

func NewGSet() *GSet {
	return &GSet{
		mainSet: mainSet{},
	}
}

func (g *GSet) Add(elem interface{}) {
	g.mainSet[elem] = struct{}{}
}

func (g *GSet) Contains(elem interface{}) bool {
	_, ok := g.mainSet[elem]
	return ok
}

func (g *GSet) Len() int {
	return len(g.mainSet)
}

func (g *GSet) Elems() []interface{} {
	elems := make([]interface{}, 0, len(g.mainSet))

	for elem := range g.mainSet {
		elems = append(elems, elem)
	}

	return elems
}

// Merge method for Gsets
func (g *GSet) Union(h *GSet) *GSet {
	s := NewGSet()
	for _, e := range g.Elems() {
		s.Add(e)
	}
	for _, e := range h.Elems() {
		if !s.Contains(e) {
			s.Add(e)
		}
	}

	return s
}

// Compare method for GSets
func (g *GSet) Subset(h *GSet) bool {
	for _, e := range g.Elems() {
		if !h.Contains(e) {
			return false
		}
	}
	return true
}

type gsetJSON struct {
	T string        `json:"type"`
	E []interface{} `json:"e"`
}

func (g *GSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(&gsetJSON{
		T: "g-set",
		E: g.Elems(),
	})
}

type TwoPhaseSet struct {
	addSet *GSet
	rmSet  *GSet
}

func NewTwoPhaseSet() *TwoPhaseSet {
	return &TwoPhaseSet{
		addSet: NewGSet(),
		rmSet:  NewGSet(),
	}
}

func (t *TwoPhaseSet) Add(elem interface{}) {
	t.addSet.Add(elem)
}

func (t *TwoPhaseSet) Remove(elem interface{}) {
	t.rmSet.Add(elem)
}

func (t *TwoPhaseSet) Contains(elem interface{}) bool {
	return t.addSet.Contains(elem) && !t.rmSet.Contains(elem)
}

// Compare method for TPSet
func (t *TwoPhaseSet) Subset(u *TwoPhaseSet) bool {
	return t.addSet.Subset(u.addSet) && t.rmSet.Subset(u.rmSet)
}

// Merge method for TPSet
func (t *TwoPhaseSet) Union(u *TwoPhaseSet) *TwoPhaseSet {
	s := NewTwoPhaseSet()
	s.addSet = t.addSet.Union(u.addSet)
	s.rmSet = t.rmSet.Union(u.rmSet)
	return s
}

type tpsetJSON struct {
	T string        `json:"type"`
	A []interface{} `json:"a"`
	R []interface{} `json:"r"`
}

func (t *TwoPhaseSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(&tpsetJSON{
		T: "2p-set",
		A: t.addSet.Elems(),
		R: t.rmSet.Elems(),
	})
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

	fmt.Println("Grow Only Set")

	obj1 := "dummy-object1"
	obj2 := "dummy-object2"
	gset1 := NewGSet()
	gset2 := NewGSet()

	gset1.Add(obj1)
	fmt.Println(gset1)

	gset2.Add(obj1)
	gset2.Add(obj2)
	fmt.Println(gset2)

	fmt.Println("gset2 is subset of gset1", gset2.Subset(gset1))

	gset3 := gset1.Union(gset2)
	fmt.Println(gset3)

	fmt.Println("TWO PHASE SET")
	ppset := NewTwoPhaseSet()
	ppset.Add(obj1)

	qqset := NewTwoPhaseSet()
	qqset.Add(obj2)
	qqset.Remove(obj2)
	ppqqset := ppset.Union(qqset)

	fmt.Println("ppqqset contains obj1", ppqqset.Contains(obj1))
	fmt.Println("ppqqset contains obj2", ppqqset.Contains(obj2))

}
