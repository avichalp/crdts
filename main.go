package main

import (
	"fmt"

	"github.com/avichalp/crdts/crdts"
)

// state base CRDT interface:
// udapte, query, compare, merge

func main() {
	fmt.Println("hello world")
	gcounter1, gcounter2 := crdts.NewGCounter(), crdts.NewGCounter()

	gcounter1.Increment()
	gcounter2.Increment()

	gcounter1.Merge(gcounter2)

	fmt.Println(gcounter1.Value())

	gcounter2.Merge(gcounter1)
	fmt.Println(gcounter2.Value())

	fmt.Println("PN Counter examples")

	pncounter1, pncounter2 := crdts.NewPNCounter(), crdts.NewPNCounter()

	pncounter1.Increment()
	pncounter2.Increment()

	pncounter1.Merge(pncounter2)

	fmt.Println("PN Counter", pncounter1.Value())

	fmt.Println("Grow Only Set")

	obj1 := "dummy-object1"
	obj2 := "dummy-object2"
	gset1 := crdts.NewGSet()
	gset2 := crdts.NewGSet()

	gset1.Add(obj1)
	fmt.Println(gset1)

	gset2.Add(obj1)
	gset2.Add(obj2)
	fmt.Println(gset2)

	fmt.Println("gset2 is subset of gset1", gset2.Subset(gset1))

	gset3 := gset1.Union(gset2)
	fmt.Println(gset3)

	fmt.Println("TWO PHASE SET")
	ppset := crdts.NewTwoPhaseSet()
	ppset.Add(obj1)

	qqset := crdts.NewTwoPhaseSet()
	qqset.Add(obj2)
	qqset.Remove(obj2)
	ppqqset := ppset.Union(qqset)

	fmt.Println("ppqqset contains obj1", ppqqset.Contains(obj1))
	fmt.Println("ppqqset contains obj2", ppqqset.Contains(obj2))

}
