package crdts


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
