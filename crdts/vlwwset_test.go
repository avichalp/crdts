package crdts

import (
	"fmt"
	"testing"

	vectorclocks "github.com/avichalp/crdts/vector_clocks"
	"github.com/stretchr/testify/assert"
)

func TestVLWWSetAddContains(t *testing.T) {
	lww, _ := NewVLLWSet(BiasAdd)

	testStr := "foo"
	assert.False(t, lww.Contains(testStr))

	lww.Add(testStr)
	assert.True(t, lww.Contains(testStr))

}

func TestVLWWSetAddRemoveContains(t *testing.T) {
	lww, _ := NewVLLWSet(BiasAdd)

	testStr := "foo"
	lww.Add(testStr)
	lww.Remove(testStr)

	// should not contain
	assert.False(t, lww.Contains(testStr))
}

func TestVLWWInvalidBias(t *testing.T) {
	var InvalidBias BiasType = "invalid bias"
	_, err := NewVLLWSet(InvalidBias)
	assert.EqualError(t, err, "given bias is not valid")

}

func TestVRemoveBias(t *testing.T) {
	lww, _ := NewVLLWSet(BiasRemove)
	e1 := "foo"
	e2 := "bar"

	lww.Add(e1)
	lww.Remove(e1)
	assert.False(t, lww.Contains(e1))

	lww.Remove(e2)
	lww.Add(e2)
	assert.False(t, lww.Contains(e2))
}

func TestVLWWSetAddRemoveConflict(t *testing.T) {
	for _, tt := range []struct {
		bias       BiasType
		testObject string
		relation   vectorclocks.Relation
		contains   bool
	}{
		{
			BiasAdd,
			"foo",
			vectorclocks.Concurrent, // concurrent Add and remove
			true,                    // when add time is **not** before rm time
		},
		{
			BiasRemove,
			"bar",
			vectorclocks.Concurrent,
			false,
		},
		{
			BiasAdd,
			"baz",
			vectorclocks.Ancestor, // add before remove
			false,
		},
		{
			BiasAdd,
			"qux",
			vectorclocks.Descendant, // remove before add
			true,
		},
		{
			BiasRemove,
			"foo",
			vectorclocks.Ancestor, // add before remove
			false,
		},
		{
			BiasRemove,
			"bar",
			vectorclocks.Descendant, // remove before add
			true,
		},
	} {
		t.Run(fmt.Sprintf("bias:%s,relation:%d", tt.bias, tt.relation), func(t *testing.T) {
			lww, _ := NewVLLWSet(tt.bias)

			switch tt.relation {
			case vectorclocks.Concurrent:
				// Add object
				lww.Add(tt.testObject)
				// remove object
				lww.Remove(tt.testObject)
				// patch the vc
				if _, ok := lww.rmMap[tt.testObject]; ok {
					addVC := lww.addMap[tt.testObject]
					lww.rmMap[tt.testObject] = addVC
				}
			case vectorclocks.Descendant:
				// remove object
				lww.Remove(tt.testObject)
				// Add object
				lww.Add(tt.testObject)
				// patch the add vc to make remove before add
				if _, ok := lww.addMap[tt.testObject]; ok {
					rmVC := lww.rmMap[tt.testObject]
					addVC := rmVC.Copy()
					addVC.Tick(lww.id)
					lww.addMap[tt.testObject] = addVC
				}
			case vectorclocks.Ancestor:
				// Add object
				lww.Add(tt.testObject)
				// remove object
				lww.Remove(tt.testObject)
			}

			assert.Equal(t, tt.contains, lww.Contains(tt.testObject))

		})
	}
}

/*
func TestVLWWSetMerge(t *testing.T) {
	type addRm struct {
		op string
		d  time.Duration
	}

	var addOp, rmOp string = "add", "remove"

	for _, tt := range []struct {
		mapOne      map[string]addRm
		mapTwo      map[string]addRm
		contains    map[string]struct{}
		notContains map[string]struct{}
	}{
		{
			map[string]addRm{
				"object1": {addOp, 1 * time.Minute},
				"object2": {addOp, 2 * time.Minute},
			},
			map[string]addRm{
				"object1": {rmOp, 2 * time.Minute},
				"object2": {rmOp, 2 * time.Minute},
			},
			map[string]struct{}{
				"object2": {},
			},
			map[string]struct{}{
				"object1": {},
			},
		},
	} {
		mock1, mock2 := clock.NewMock(), clock.NewMock()
		lww1, _ := NewLLWSet(BiasAdd)
		lww1.clock = mock1

		lww2, _ := NewLLWSet(BiasAdd)
		lww2.clock = mock2

		var totalDuration time.Duration

		for obj, addrm := range tt.mapOne {
			curTime := addrm.d - totalDuration

			totalDuration += curTime
			mock1.Add(curTime)

			switch addrm.op {
			case addOp:
				lww1.Add(obj)
			case rmOp:
				lww1.Remove(obj)
			}
		}

		totalDuration = 0 * time.Second

		for obj, addrm := range tt.mapTwo {
			curTime := addrm.d - totalDuration

			totalDuration += curTime
			mock2.Add(curTime)

			switch addrm.op {
			case addOp:
				lww2.Add(obj)
			case rmOp:
				lww2.Remove(obj)
			}
		}

		lww1.Merge(lww2)

		for obj := range tt.contains {
			assert.True(t, lww1.Contains(obj))
		}

		for obj := range tt.notContains {
			assert.False(t, lww1.Contains(obj))
		}

	}
} */
