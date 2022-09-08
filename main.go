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

	// fmt.Println("gset2 is subset of gset1", gset2.set.Subset(gset1.set))

	// gset3 := gset1.Union(gset2)
	// fmt.Println(gset3)

	fmt.Println("TWO PHASE SET")
	ppset := crdts.NewTwoPhaseSet()
	ppset.Add(obj1)

	qqset := crdts.NewTwoPhaseSet()
	qqset.Add(obj2)
	qqset.Remove(obj2)
	ppqqset := ppset.Union(qqset)

	fmt.Println("ppqqset contains obj1", ppqqset.Contains(obj1))
	fmt.Println("ppqqset contains obj2", ppqqset.Contains(obj2))

	AddLWWSet, _ := crdts.NewLLWSet(crdts.BiasAdd)
	AddLWWSet.Remove(obj1)
	AddLWWSet.Add(obj1)
	fmt.Println("Add LWW set contains obj1 after removal", AddLWWSet.Contains(obj1))

	RmLWWSet, _ := crdts.NewLLWSet(crdts.BiasRemove)
	RmLWWSet.Remove(obj1)
	RmLWWSet.Remove(obj1)
	fmt.Println("Remove LWW set contains obj1 after removal", RmLWWSet.Contains(obj1))

	AddLWWSet1, _ := crdts.NewLLWSet(crdts.BiasAdd)
	AddLWWSet2, _ := crdts.NewLLWSet(crdts.BiasAdd)
	AddLWWSet1.Add(obj1)
	AddLWWSet2.Add(obj2)
	AddLWWSet1.Merge(AddLWWSet2)
	fmt.Println("Merge obj2 from set 2 into 1?", AddLWWSet1.Contains(obj2))

	// ORSet
	orset := crdts.NewORSet()
	// inserting same object
	orset.Add(obj1)
	orset.Add(obj1)

	// removing one
	orset.Remove(obj1)

	fmt.Println("must not contain obj1", orset.Contains(obj1))

	orset2 := crdts.NewORSet()
	orset2.Add(obj1)
	orset2.Remove(obj1)
	orset2.Add(obj1)

	fmt.Println("must contain obj1", orset2.Contains(obj1))

}
