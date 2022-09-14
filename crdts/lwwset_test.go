package crdts

import (
	"fmt"
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/stretchr/testify/assert"
)

func TestLWWSetAddContains(t *testing.T) {
	lww, err := NewLLWSet(BiasAdd)
	if err != nil {
		t.Fatalf("cannot create lww set")
	}

	testStr := "foo"
	assert.False(t, lww.Contains(testStr))

	lww.Add(testStr)
	assert.True(t, lww.Contains(testStr))

}

func TestLWWSetAddRemoveContains(t *testing.T) {
	lww, err := NewLLWSet(BiasAdd)
	if err != nil {
		t.Fatal("cannot create new lww set")
	}

	testStr := "foo"
	lww.Add(testStr)
	lww.Remove(testStr)
	assert.False(t, lww.Contains(testStr))
}

func TestInvalidBias(t *testing.T) {
	var InvalidBias BiasType = "invalid bias"
	_, err := NewLLWSet(InvalidBias)
	assert.EqualError(t, err, "given bias is not valid")

}

func TestRemoveBias(t *testing.T) {
	lww, _ := NewLLWSet(BiasRemove)

	mock := clock.NewMock()
	lww.clock = mock

	e := "foo"

	// Remove before Add. Since it is a set with
	// remove bias the any add after remove will not
	// 'add' element to the lww set
	lww.Add(e)
	// rollback the clock by 10 mins
	mock.Add(-10 * time.Minute)

	lww.Remove(e)

	fmt.Println(lww.addMap)
	fmt.Println(lww.rmMap)

	assert.True(t, lww.Contains(e))

}

func TestLWWSetAddRemoveConflict(t *testing.T) {
	for _, tt := range []struct {
		bias       BiasType
		testObject string
		elapsed    time.Duration
		contains   bool
	}{
		{
			BiasAdd,
			"foo",
			0,    // concurrent Add and remove
			true, // when add time is **not** before rm time
		},
		{
			BiasRemove,
			"bar",
			0,
			false,
		},
		{
			BiasAdd,
			"baz",
			1 * time.Minute,
			false,
		},
		{
			BiasAdd,
			"qux",
			-1 * time.Minute,
			true,
		},
		{
			BiasRemove,
			"foo",
			1 * time.Minute,
			false,
		},
		{
			BiasRemove,
			"bar",
			-1 * time.Minute,
			true,
		},
	} {
		t.Run(fmt.Sprintf("bias:%s,elapsed:%d", tt.bias, tt.elapsed), func(t *testing.T) {
			lww, _ := NewLLWSet(tt.bias)

			// replace clock with a mock clock
			mock := clock.NewMock()
			lww.clock = mock

			// Add object
			lww.Add(tt.testObject)
			// move time forward or backward
			mock.Add(tt.elapsed)
			// remove object
			lww.Remove(tt.testObject)
			assert.Equal(t, tt.contains, lww.Contains(tt.testObject))

		})

	}
}

func TestLWWSetMerge(t *testing.T) {
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
}
