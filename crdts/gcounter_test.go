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
