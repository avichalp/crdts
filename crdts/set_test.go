package crdts

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGSetAddContains(t *testing.T) {
	gset := NewGSet()

	elem := "foo"
	assert.False(t, gset.Contains(elem))

	gset.Add(elem)
	assert.True(t, gset.Contains(elem))

}

func TestGSetElems(t *testing.T) {
	for _, tt := range []struct {
		add []interface{}
	}{
		{[]interface{}{}},
		{[]interface{}{1}},
		{[]interface{}{1, -2, 3}},
	} {
		gset := NewGSet()
		expectedElems := map[interface{}]struct{}{}
		for _, i := range tt.add {
			expectedElems[i] = struct{}{}
			gset.Add(i)
		}

		actualElems := map[interface{}]struct{}{}
		for _, i := range gset.Elems() {
			actualElems[i] = struct{}{}
		}

		assert.EqualValues(t, expectedElems, actualElems)

	}

}

func TestGSetCompare(t *testing.T) {
	for _, tt := range []struct {
		addOne []interface{}
		addTwo []interface{}
	}{
		{[]interface{}{}, []interface{}{1}},
		{[]interface{}{"foo"}, []interface{}{"bar", "foo"}},
	} {
		gsetOne, gsetTwo := NewGSet(), NewGSet()

		expectedElemsOne := map[interface{}]struct{}{}
		for _, i := range tt.addOne {
			expectedElemsOne[i] = struct{}{}
			gsetOne.Add(i)
		}

		expectedElemsTwo := map[interface{}]struct{}{}
		for _, i := range tt.addTwo {
			expectedElemsTwo[i] = struct{}{}
			gsetTwo.Add(i)
		}

		assert.False(t, gsetTwo.Compare(gsetOne))
		assert.True(t, gsetOne.Compare(gsetTwo))
	}

}

// union
func TestGSetMerge(t *testing.T) {
	for _, tt := range []struct {
		addOne []interface{}
		addTwo []interface{}
		union  []interface{}
	}{
		{
			[]interface{}{},
			[]interface{}{1},
			[]interface{}{1},
		},
		{
			[]interface{}{"foo"},
			[]interface{}{"bar", "foo"},
			[]interface{}{"foo", "bar"},
		},
	} {
		gsetOne, gsetTwo := NewGSet(), NewGSet()

		expectedElemsOne := map[interface{}]struct{}{}
		for _, i := range tt.addOne {
			expectedElemsOne[i] = struct{}{}
			gsetOne.Add(i)
		}

		expectedElemsTwo := map[interface{}]struct{}{}
		for _, i := range tt.addTwo {
			expectedElemsTwo[i] = struct{}{}
			gsetTwo.Add(i)
		}

		gset := gsetOne.Merge(gsetTwo)

		fmt.Println("gsetOne", gsetOne.set)
		fmt.Println("gsetTwo", gsetTwo.set)
		fmt.Println("gset", gset.set)

		for e := range gsetOne.set {
			fmt.Println(e)
			assert.True(t, gset.Contains(e))
		}
		for e := range gsetTwo.set {
			fmt.Println(e)
			assert.True(t, gset.Contains(e))
		}

	}
}
