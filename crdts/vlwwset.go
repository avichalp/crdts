package crdts

import (
	"fmt"
	"math/rand"

	vectorclocks "github.com/avichalp/crdts/vector_clocks"
)

// TODO: LogicalLWW: last writer wins with vector clocks as timestamp
type VLWWSet struct {
	id     string
	addMap map[interface{}]vectorclocks.VectorClock
	rmMap  map[interface{}]vectorclocks.VectorClock

	bias BiasType
}

func NewVLLWSet(bias BiasType) (*VLWWSet, error) {
	if bias != BiasAdd && bias != BiasRemove {
		return nil, ErrInvalidBias
	}

	return &VLWWSet{
		// set vector clocks process id here using v.Set()
		id:     fmt.Sprintf("%d", rand.Int()),
		addMap: make(map[interface{}]vectorclocks.VectorClock),
		rmMap:  make(map[interface{}]vectorclocks.VectorClock),
		bias:   bias,
	}, nil
}

func (s *VLWWSet) Add(value interface{}) {
	// if addMap is empty create a new vector clock
	// else do a Tick on the previous vector clock
	// store all of them?  -- yes bc of LWW semantics

	// we overwrite the value since it is a LWW set
	_, ok := s.addMap[value]
	if ok {
		// tick
		s.addMap[value].Tick(s.id)
	} else {
		// set
		vc := vectorclocks.VectorClock{}
		vc.Set(s.id, 1)
		s.addMap[value] = vc
	}

}

func (s *VLWWSet) Remove(value interface{}) {

	// if value is in add set, add it in the rm set
	// with incremented vector clock
	addVC, addOK := s.addMap[value]
	if addOK {
		rmVC := addVC.Copy()
		rmVC.Tick(s.id)
		s.rmMap[value] = rmVC
		return
	}

	// if value is in rm set already, re add in the
	// increment its vector clock
	_, rmOK := s.rmMap[value]
	if rmOK {
		// tick
		s.rmMap[value].Tick(s.id)
	} else {
		// if value is not in add or remove sets
		//  add it to remove set with a new vector clock
		// set
		vc := vectorclocks.VectorClock{}
		vc.Set(s.id, 1)
		s.rmMap[value] = vc
	}

}

func (s *VLWWSet) Contains(value interface{}) bool {
	addVC, addOk := s.addMap[value]
	if !addOk {
		return false
	}

	rmVC, rmOk := s.rmMap[value]
	if !rmOk {
		return true
	}

	switch s.bias {
	case BiasAdd:
		// value is memeber of the set
		// if rmVC is not decendant of the addVC
		// if it is equal or concurrent or comes before (ancestor)
		// then the given value is member of the set
		//
		// rmVC  -> addVC: true
		// addVC -> rmVC : false
		// addVC == rmVC: true
		// addVC || rmVC: true (cannot compare, add bias)
		return !rmVC.Descendant(addVC)
	case BiasRemove:
		// in LWW Remove Bias set an element cannot be added
		// if it is removed at an earlier point in time

		// rmVC  -> addVC: true
		// addVC -> rmVC : false
		// addVC == rmVC: false
		// addVC || rmVC: false (cannot compare) remove wins
		return addVC.Descendant(rmVC)
	}

	return false
}

func (s *VLWWSet) Merge(r *VLWWSet) {
	for value, rvc := range r.addMap {
		if svc, ok := s.addMap[value]; ok && svc.Relation(rvc) == vectorclocks.Ancestor {
			s.addMap[value] = rvc
		} else {
			if svc.Relation(rvc) == vectorclocks.Ancestor {
				s.addMap[value] = rvc
			} else {
				s.addMap[value] = svc
			}
		}
	}

	for value, rvc := range r.rmMap {
		if svc, ok := s.rmMap[value]; ok && svc.Relation(rvc) == vectorclocks.Ancestor {
			s.rmMap[value] = rvc
		} else {
			if svc.Relation(rvc) == vectorclocks.Ancestor {
				s.rmMap[value] = rvc
			} else {
				s.rmMap[value] = svc
			}
		}
	}
}

// todo: lwwset tests should pass of vlww-set
