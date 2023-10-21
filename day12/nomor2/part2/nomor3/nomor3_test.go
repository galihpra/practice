package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayingDomino(t *testing.T) {
	type TestCase struct {
		Parameter1 [][]int
		Parameter2 []int
		Expected   []int
		Message    string
	}

	theTestCase := []TestCase{
		{[][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3}, []int{3, 4}, "Output tidak sesuai"},
		{[][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6}, []int{6, 5}, "Output tidak sesuai"},
		{[][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}, []int(nil), "Output tidak sesuai"},
	}

	for i := 0; i < len(theTestCase); i++ {
		t.Run(fmt.Sprint("Case ", i+1), func(t *testing.T) {
			assert.Equal(t, theTestCase[i].Expected, PlayingDomino(theTestCase[i].Parameter1, theTestCase[i].Parameter2),
				theTestCase[i].Message)
		})
	}

}
