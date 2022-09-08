package crdts

type Set map[interface{}]struct{}

type GSet struct {
	set Set
}

func NewGSet() *GSet {
	return &GSet{
		set: make(Set),
	}
}

func (s *GSet) Elems() []interface{} {
	elems := make([]interface{}, 0, len(s.set))

	for elem := range s.set {
		elems = append(elems, elem)
	}

	return elems
}

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

func (s *GSet) Add(elem interface{}) {
	s.set[elem] = struct{}{}
}

func (g *GSet) Contains(elem interface{}) bool {
	_, ok := g.set[elem]
	return ok
}

func (g *GSet) Compare(h *GSet) bool {
	return g.Subset(h)
}

func (g *GSet) Merge(h *GSet) *GSet {
	return g.Union(h)
}
