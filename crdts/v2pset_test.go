package crdts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVSetAddContains(t *testing.T) {
	set, testStr := NewVTwoPhaseSet(), "foo"
	contains, vc := set.Contains(testStr)
	assert.EqualValues(t, map[string]uint64{set.id: 1}, vc)
	assert.False(t, contains)

	set.Add(testStr)
	contains, vc = set.Contains(testStr)
	assert.True(t, contains)
	assert.EqualValues(t, map[string]uint64{set.id: 2}, vc)

}

func TestVSetAddRemoveContains(t *testing.T) {
	set, testStr := NewVTwoPhaseSet(), "foo"

	set.Add(testStr)
	set.Remove(testStr)

	contains, vc := set.Contains(testStr)

	// should not contain since remove follows add
	assert.False(t, contains)
	assert.EqualValues(t, map[string]uint64{set.id: 3}, vc)
}

func TestVRemoveBias(t *testing.T) {
	lww := NewVTwoPhaseSet()
	e1 := "foo"
	e2 := "bar"

	lww.Add(e1)
	lww.Remove(e1)
	contains, _ := lww.Contains(e1)
	// todo: test vc
	assert.False(t, contains)

	lww.Remove(e2)
	lww.Add(e2)
	contains, _ = lww.Contains(e2)
	assert.False(t, contains)
}

/* func TestVLWWSetMerge(t *testing.T) {
	lww1 := NewVTwoPhaseSet()
	lww2 := NewVTwoPhaseSet()

	lww1.Add("bar")
	lww1.Add("bar") // next time lww1 see bar it will increase the counter

	lww2.Remove("foo")
	lww2.Remove("foo")
	// lww2.rmMap["foo"].Tick(lww2.id)

	lww1.Add("foo")

	// fmt.Println(lww1, lww2)

	lww1.Merge(lww2)

	// fmt.Println(lww1, lww2)

	contains, vc := lww1.Contains("foo")
	fmt.Println(vc)
	assert.True(t, contains)
}
*/
