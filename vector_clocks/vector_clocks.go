package vectorclocks

import "fmt"

type VectorClock map[string]uint64

type Relation uint8

const (
	Equal = iota
	Ancestor
	Descendant
	Concurrent
)

func (v VectorClock) Tick(pid string) {
	v[pid] += 1
}

func (v VectorClock) Get(pid string) uint64 {
	return v[pid]
}

func (v VectorClock) Set(pid string, tick uint64) {
	if tick == 0 {
		panic(fmt.Errorf("tick cannot be 0"))
	}
	v[pid] = tick
}

func (v VectorClock) Copy() (w VectorClock) {
	w = make(VectorClock, len(v))
	for pid := range v {
		w[pid] = v[pid]
	}
	return
}

func (v VectorClock) Merge(w VectorClock) (x VectorClock) {
	x = v.Copy()
	for pid := range w {
		if x[pid] < w[pid] {
			x[pid] = w[pid]
		}
	}
	return
}

func (v VectorClock) Equal(w VectorClock) bool {
	if len(w) != len(v) {
		return false
	}

	for pid := range w {
		if v[pid] != w[pid] {
			return false
		}
	}
	return true
}

// v is descendant of w iff:
// 1. all elements in w is less than or equal than v and
// 2. v != w
func (v VectorClock) Descendant(w VectorClock) bool {
	isEqual := len(w) == len(v)
	for pid := range w {
		if w[pid] > v[pid] {
			return false
		} else if isEqual && w[pid] < v[pid] {
			isEqual = false
		}
	}
	return !isEqual
}

func (v VectorClock) Relation(w VectorClock) Relation {
	if v.Equal(w) {
		return Equal
	} else if w.Descendant(v) {
		return Ancestor
	} else if v.Descendant(w) {
		return Descendant
	} else {
		return Concurrent
	}
}
