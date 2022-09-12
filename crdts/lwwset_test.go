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
