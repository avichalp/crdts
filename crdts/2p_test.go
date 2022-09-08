package crdts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoPhaseSetAdd(t *testing.T) {
	tpset := NewTwoPhaseSet()
	elem := "foo"

	assert.False(t, tpset.Contains(elem))

	tpset.Add(elem)

	assert.True(t, tpset.Contains(elem))
}

func TestTwoPhaseSetAddRemove(t *testing.T) {
	elemOne := "foo"
	tpsetOne := NewTwoPhaseSet()
	tpsetOne.Add(elemOne)
	assert.True(t, tpsetOne.Contains(elemOne))
	tpsetOne.Remove(elemOne)
	assert.False(t, tpsetOne.Contains(elemOne))

	// first remove, then add
	elemTwo := "bar"
	tpsetTwo := NewTwoPhaseSet()
	tpsetTwo.Remove(elemTwo)
	tpsetTwo.Add(elemTwo)
	assert.False(t, tpsetTwo.Contains(elemTwo))
}

func TestTwoPhaseSetCompare(t *testing.T) {
	for _, tt := range []struct {
		addE1    interface{}
		addE2    interface{}
		removeE1 interface{}
		removeE2 interface{}
		compare  bool
	}{
		{ // both add and rm sets of 2 are empty
			"foo",
			nil,
			"bar",
			nil,
			true,
		},
		{ // idendical add set and rm sets
			"foo",
			"foo",
			nil,
			nil,
			true,
		},
		{ // idendical and and rm sets
			nil,
			nil,
			"bar",
			"bar",
			true,
		},
		{ // disjont add sets, overlapping rm set
			"foo",
			"bar",
			"baz",
			nil,
			false,
		},
		{ // disjoint rm sets, overlapping add set
			"baz",
			nil,
			"foo",
			"bar",
			false,
		},
		{ // both add and rm sets are disjoint
			"foo",
			"baz",
			"bar",
			"qux",
			false,
		},
	} {
		tpsetOne, tpsetTwo := NewTwoPhaseSet(), NewTwoPhaseSet()
		if tt.addE1 != nil {
			tpsetOne.Add(tt.addE1)
		}
		if tt.addE2 != nil {
			tpsetTwo.Add(tt.addE2)
		}
		if tt.removeE1 != nil {
			tpsetOne.Remove(tt.removeE1)
		}
		if tt.removeE2 != nil {
			tpsetTwo.Remove(tt.removeE2)
		}
		assert.Equal(t, tt.compare, tpsetTwo.Compare(tpsetOne))
	}
}

func TestTwoPhaseSetMerge(t *testing.T) {
	for _, tt := range []struct {
		addE1    interface{}
		addE2    interface{}
		removeE1 interface{}
		removeE2 interface{}
	}{
		{
			"foo",
			"baz",
			"bar",
			"qux",
		},
		{
			1,
			2,
			3,
			4,
		},
		{
			"hello",
			"hola",
			"bonjour",
			"guten tag",
		},
	} {
		tpsetOne, tpsetTwo := NewTwoPhaseSet(), NewTwoPhaseSet()
		tpsetOne.Add(tt.addE1)
		tpsetOne.Remove(tt.removeE1)
		tpsetTwo.Add(tt.addE2)
		tpsetTwo.Remove(tt.removeE2)

		// merged set
		tpset := tpsetOne.Merge(tpsetTwo)

		assert.True(t, tpset.Contains(tt.addE1))
		assert.False(t, tpset.Contains(tt.removeE1))
		assert.True(t, tpset.Contains(tt.addE1))
		assert.False(t, tpset.Contains(tt.removeE2))

	}

}
