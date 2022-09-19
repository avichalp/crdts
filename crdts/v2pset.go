package crdts

import (
	"fmt"
	"math/rand"

	vectorclocks "github.com/avichalp/crdts/vector_clocks"
)

// TODO: LogicalLWW: last writer wins with vector clocks as timestamp
type VTwoPhaseSet struct {
	id     string
	addMap map[interface{}]struct{}
	rmMap  map[interface{}]struct{}
	vc     vectorclocks.VectorClock
}

func NewVTwoPhaseSet() *VTwoPhaseSet {
	id := fmt.Sprintf("node_id:%d", rand.Int())
	return &VTwoPhaseSet{
		// set vector clocks process id here using v.Set()
		id:     id,
		addMap: make(map[interface{}]struct{}),
		rmMap:  make(map[interface{}]struct{}),
		vc:     vectorclocks.VectorClock{id: 1},
	}
}

func (s *VTwoPhaseSet) Add(value interface{}) {
	s.addMap[value] = struct{}{}
	s.vc.Tick(s.id)
}

func (s *VTwoPhaseSet) Remove(value interface{}) {
	s.rmMap[value] = struct{}{}
	s.vc.Tick(s.id)
}

func (s *VTwoPhaseSet) Contains(value interface{}) (bool, vectorclocks.VectorClock) {
	if _, addOk := s.addMap[value]; !addOk {
		return false, s.vc
	}

	if _, rmOk := s.rmMap[value]; !rmOk {
		return true, s.vc
	}

	return false, s.vc
}

func (s *VTwoPhaseSet) Merge(r *VTwoPhaseSet) {
	s.vc = r.vc.Merge(s.vc)

	for value := range r.addMap {
		s.addMap[value] = struct{}{}
	}

	for value := range r.rmMap {
		s.rmMap[value] = struct{}{}
	}
}
