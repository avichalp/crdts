package crdts

import (
	// "fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestORSetAddContains(t *testing.T) {
	orSet, testValue := NewORSet(), "foo"

	assert.False(t, orSet.Contains(testValue))
	orSet.Add(testValue)
	assert.True(t, orSet.Contains(testValue))

}

func TestORSetAddRemoveContains(t *testing.T) {
	orSet, testValue := NewORSet(), "foo"

	orSet.Add(testValue)
	orSet.Remove(testValue)

	assert.False(t, orSet.Contains(testValue))
}

func TestORSetAddRemoveAddContains(t *testing.T) {
	orSet, testValue := NewORSet(), "foo"

	orSet.Add(testValue)
	orSet.Remove(testValue)
	orSet.Add(testValue)

	assert.True(t, orSet.Contains(testValue))
}

func TestORSetAddAddRemoveContains(t *testing.T) {
	orSet, testValue := NewORSet(), "foo"

	orSet.Add(testValue)
	orSet.Add(testValue)
	orSet.Remove(testValue)

	assert.False(t, orSet.Contains(testValue))
}

func TestORSetMerge(t *testing.T) {
	type addRm struct {
		addSet []string
		rmSet  []string
	}

	for _, tt := range []struct {
		set1        addRm
		set2        addRm
		contains    map[string]struct{}
		notContains map[string]struct{}
	}{
		{
			addRm{[]string{"foo"}, []string{}},
			addRm{[]string{}, []string{"foo"}},
			map[string]struct{}{
				"foo": {},
			},
			map[string]struct{}{},
		},
		{
			// rm in 1st set; add in 2nd set
			addRm{[]string{}, []string{"foo"}},
			addRm{[]string{"foo"}, []string{}},
			map[string]struct{}{
				"foo": {},
			},
			map[string]struct{}{},
		},
		{
			addRm{[]string{"foo"}, []string{"foo"}},
			addRm{[]string{}, []string{}},
			map[string]struct{}{},
			map[string]struct{}{
				"foo": {},
			},
		},
		{
			addRm{[]string{}, []string{}},
			addRm{[]string{"bar"}, []string{"bar"}},
			map[string]struct{}{},
			map[string]struct{}{
				"bar": {},
			},
		},
		{
			// adds foo in 1st set and rm it in the 2nd set
			// rm bar from 1st set and add it in 2nd set
			addRm{[]string{"foo"}, []string{"bar"}},
			addRm{[]string{"bar"}, []string{"foo"}},
			map[string]struct{}{
				"bar": {},
				"foo": {},
			},
			map[string]struct{}{},
		},
		{
			// contains returns 'true' if any one 'addMap' entry
			// is not present in the 'rmSet'
			addRm{[]string{"foo", "bar"}, []string{"bar"}},
			addRm{[]string{"bar", "foo"}, []string{"foo"}},
			map[string]struct{}{
				"bar": {},
				"foo": {},
			},
			map[string]struct{}{},
		},
		{
			addRm{[]string{"foo", "bar"}, []string{"bar", "foo"}},
			addRm{[]string{"bar", "foo"}, []string{"foo", "bar"}},
			map[string]struct{}{},
			map[string]struct{}{
				"bar": {},
				"foo": {},
			},
		},
	} {
		orset1, orset2 := NewORSet(), NewORSet()

		for _, add := range tt.set1.addSet {
			orset1.Add(add)
		}

		for _, rm := range tt.set1.rmSet {
			orset1.Remove(rm)
		}

		for _, add := range tt.set2.addSet {
			orset2.Add(add)
		}

		for _, rm := range tt.set2.rmSet {
			orset2.Remove(rm)
		}

		orset1.Merge(orset2)

		for obj := range tt.contains {
			assert.True(t, orset1.Contains(obj))
		}

		for obj := range tt.notContains {
			assert.False(t, orset1.Contains(obj))
		}
	}

}
