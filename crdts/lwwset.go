package crdts

import (
	"errors"
	"time"

	"github.com/benbjohnson/clock"
)

type BiasType string

type LWWSet struct {
	addMap map[interface{}]time.Time
	rmMap  map[interface{}]time.Time

	bias  BiasType
	clock clock.Clock
}

const (
	BiasAdd    BiasType = "a"
	BiasRemove BiasType = "r"
)

var (
	ErrInvalidBias = errors.New("given bias is not valid")
)

func NewLLWSet(bias BiasType) (*LWWSet, error) {
	if bias != BiasAdd && bias != BiasRemove {
		return nil, ErrInvalidBias
	}

	return &LWWSet{
		addMap: make(map[interface{}]time.Time),
		rmMap:  make(map[interface{}]time.Time),
		bias:   bias,
		clock:  clock.New(),
	}, nil
}

func (s *LWWSet) Add(value interface{}) {
	s.addMap[value] = s.clock.Now()
}

func (s *LWWSet) Remove(value interface{}) {
	s.rmMap[value] = s.clock.Now()
}

func (s *LWWSet) Contains(value interface{}) bool {
	addTime, addOk := s.addMap[value]
	if !addOk {
		return false
	}

	rmTime, rmOk := s.rmMap[value]
	if !rmOk {
		return true
	}

	switch s.bias {
	case BiasAdd:
		return !addTime.Before(rmTime)
	case BiasRemove:
		return rmTime.Before(addTime)
	}

	return false
}

func (s *LWWSet) Merge(r *LWWSet) {
	for value, ts := range r.addMap {
		if t, ok := s.addMap[value]; ok && t.Before(ts) {
			s.addMap[value] = ts
		} else {
			if t.Before(ts) {
				s.addMap[value] = ts
			} else {
				s.addMap[value] = t
			}
		}
	}

	for value, ts := range r.rmMap {
		if t, ok := s.rmMap[value]; ok && t.Before(ts) {
			s.rmMap[value] = ts
		} else {
			if t.Before(ts) {
				s.rmMap[value] = ts
			} else {
				s.rmMap[value] = t
			}
		}
	}
}
