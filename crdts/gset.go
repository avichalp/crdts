package crdts

import "encoding/json"

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

