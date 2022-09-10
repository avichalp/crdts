package vectorclocks

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSet(t *testing.T) {
	vc := VectorClock{}
	vc.Set("a", 1)
	assert.Equal(t, uint64(1), vc.Get("a"))
}

func TestTick(t *testing.T) {
	vc := VectorClock{}
	var i uint64
	for i = 1; i <= 10; i++ {
		vc.Tick("a")
		assert.Equal(t, i, vc.Get("a"))
	}
}

func TestCopy(t *testing.T) {
	vc1 := VectorClock{"a": 1}
	vc2 := vc1.Copy()
	vc2.Tick("a")
	assert.Equal(t, uint64(1), vc1.Get("a"))
	assert.Equal(t, uint64(2), vc2.Get("a"))
}

func TestMerge(t *testing.T) {
	vc1 := VectorClock{"a": 1, "b": 1}
	vc2 := VectorClock{"b": 2, "c": 1}
	vc := vc1.Merge(vc2)
	expected := VectorClock{"a": 1, "b": 2, "c": 1}
	assert.Equal(t, expected, vc)
}

func TestDescendant(t *testing.T) {
	for _, tt := range []struct {
		vc         VectorClock
		descendant bool
	}{
		{VectorClock{"a": 1, "b": 2}, true},
		{VectorClock{"a": 1, "c": 1}, false},
		{VectorClock{"a": 1}, false},
		{VectorClock{"a": 1, "b": 1, "c": 1}, true},
	} {
		vc := VectorClock{"a": 1, "b": 1}
		t.Run(vc.String(), func(t *testing.T) {
			assert.Equal(t, tt.descendant, tt.vc.Descendant(vc))
		})
	}
}

func TestRelation(t *testing.T) {
	for _, tt := range []struct {
		vc       VectorClock
		relation Relation
	}{
		{VectorClock{"a": 1, "b": 1}, Equal},
		{VectorClock{"a": 1, "b": 1, "c": 1}, Ancestor},
		{VectorClock{"b": 1}, Descendant},
		{VectorClock{"a": 1, "c": 1}, Concurrent},
	} {
		vc := VectorClock{"a": 1, "b": 1}
		t.Run(tt.vc.String(), func(t *testing.T) {
			assert.Equal(t, tt.relation, vc.Relation(tt.vc))
		})
	}
}

func BenchmarkMerge(b *testing.B) {
	for _, bb := range []struct {
		vc1 VectorClock
		vc2 VectorClock
	}{
		{VectorClock{"a": 1, "b": 1}, VectorClock{"a": 1, "b": 1}},
		{VectorClock{"a": 1, "b": 1, "c": 1}, VectorClock{"a": 1, "b": 1}},
		{VectorClock{"b": 1}, VectorClock{"a": 1, "b": 1}},
		{VectorClock{"a": 1, "c": 1}, VectorClock{"a": 1, "b": 1}},
	} {
		b.Run(fmt.Sprintf("{%s}+{%s}", bb.vc1.String(), bb.vc2.String()), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				bb.vc1.Merge(bb.vc2)
			}
		})
	}
}

func BenchmarkEqual(b *testing.B) {
	for _, bb := range []struct {
		vc1 VectorClock
		vc2 VectorClock
	}{
		{VectorClock{"a": 1, "b": 1}, VectorClock{"a": 1, "b": 1}},
		{VectorClock{"a": 1, "b": 1, "c": 1}, VectorClock{"a": 1, "b": 1}},
		{VectorClock{"b": 1}, VectorClock{"a": 1, "b": 1}},
		{VectorClock{"a": 1, "c": 1}, VectorClock{"a": 1, "b": 1}},
	} {
		b.Run(fmt.Sprintf("{%s}+{%s}", bb.vc1.String(), bb.vc2.String()), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				bb.vc1.Equal(bb.vc2)
			}
		})
	}
}

func BenchmarkDecendant(b *testing.B) {
	for _, bb := range []struct {
		vc1 VectorClock
		vc2 VectorClock
	}{
		{VectorClock{"a": 1, "b": 1}, VectorClock{"a": 1, "b": 1}},
		{VectorClock{"a": 1, "b": 1, "c": 1}, VectorClock{"a": 1, "b": 1}},
		{VectorClock{"b": 1}, VectorClock{"a": 1, "b": 1}},
		{VectorClock{"a": 1, "c": 1}, VectorClock{"a": 1, "b": 1}},
	} {
		b.Run(fmt.Sprintf("{%s}+{%s}", bb.vc1.String(), bb.vc2.String()), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				bb.vc1.Descendant(bb.vc2)
			}
		})
	}

}
