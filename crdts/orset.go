package crdts

import (
	"github.com/google/uuid"
)

type ORSet struct {
	addMap map[interface{}]map[string]struct{}
	rmMap  map[interface{}]map[string]struct{}
}

func NewORSet() *ORSet {
	return &ORSet{
		addMap: make(map[interface{}]map[string]struct{}),
		rmMap:  make(map[interface{}]map[string]struct{}),
	}
}

func (o *ORSet) Add(value interface{}) {
	if m, ok := o.addMap[value]; ok {
		m[uuid.New().String()] = struct{}{}
		o.addMap[value] = m
		return
	}

	m := make(map[string]struct{})
	m[uuid.New().String()] = struct{}{}
	o.addMap[value] = m
}

func (o *ORSet) Remove(value interface{}) {
	r, ok := o.rmMap[value]
	if !ok {
		r = make(map[string]struct{})
	}

	if m, ok := o.addMap[value]; ok {
		for uid := range m {
			r[uid] = struct{}{}
		}
	}

	o.rmMap[value] = r

}

func (o *ORSet) Contains(value interface{}) bool {
	addMap, ok := o.addMap[value]
	if !ok {
		return false
	}

	rmMap, ok := o.rmMap[value]
	if !ok {
		return true
	}

	// for all occurance in add set
	// if anyone of them is not present
	// in the remove set
	for uid := range addMap {
		if _, ok := rmMap[uid]; !ok {
			return true
		}
	}

	return false
}

func (o *ORSet) Merge(r *ORSet) {
	for value, m := range r.addMap {
		addMap, ok := o.addMap[value]
		if ok {
			// set union of o and r's add set
			for uid := range m {
				addMap[uid] = struct{}{}
			}
			continue
		}
		o.addMap[value] = m
	}

	for value, m := range r.rmMap {
		rmMap, ok := o.rmMap[value]
		if ok {
			for uid := range m {
				rmMap[uid] = struct{}{}
			}
			continue
		}
		o.rmMap[value] = m
	}
}
