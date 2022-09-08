package crdts

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

