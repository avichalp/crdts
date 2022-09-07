package crdts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGCounter(t *testing.T) {
	for _, tt := range []struct {
		incsOne int
		incsTwo int
		result  int
	}{
		{5, 10, 15},
		{10, 5, 15},
		{100, 100, 200},
		{1, 2, 3},
	} {
		gOne, gTwo := NewGCounter(), NewGCounter()

		// increment gOne
		for i := 0; i < tt.incsOne; i++ {
			gOne.Increment()
		}

		// increment gTwo
		for i := 0; i < tt.incsTwo; i++ {
			gTwo.Increment()
		}

		gOne.Merge(gTwo)
		assert.Equal(t, tt.result, gOne.Value())

		gTwo.Merge(gOne)
		assert.Equal(t, tt.result, gTwo.Value())
	}

}

func TestPNCounter(t *testing.T) {
	for _, tt := range []struct {
		incOne int
		decOne int
		incTwo int
		decTwo int
		result int
	}{
		{5, 5, 6, 6, 0},
		{5, 6, 7, 8, -2},
		{8, 7, 6, 5, 2},
		{5, 0, 6, 0, 11},
		{0, 5, 0, 6, -11},
	} {
		pOne, pTwo := NewPNCounter(), NewPNCounter()

		for i := 0; i < tt.incOne; i++ {
			pOne.Increment()
		}
		
		for i := 0; i < tt.incTwo; i++ {
			pTwo.Increment()
		}

		for i := 0; i < tt.decOne; i++ {
			pOne.Decrement()
		}

		for i := 0; i < tt.decTwo; i++ {
			pTwo.Decrement()
		}

		pOne.Merge(pTwo)
		assert.Equal(t, tt.result, pOne.Value())

		pTwo.Merge(pOne)
		assert.Equal(t, tt.result, pTwo.Value())
	}
}
