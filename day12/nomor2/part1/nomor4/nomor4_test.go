package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSequence(t *testing.T) {
	type TestCase struct {
		Parameter []int
		Expected  int
		Message   string
	}

	theTestCase := []TestCase{
		{[]int{-2, 1 - 3, 4, -1, 2, 1, -5, 4}, 6, "Output tidak sesuai"},
		{[]int{-2, -3, 4, -1, -2, 1, 5, -3}, 7, "Output tidak sesuai"},
		{[]int{-2, -5, 6, 2, -3, 1, 6, -6}, 12, "Output tidak sesuai"},
	}

	for i := 0; i < len(theTestCase); i++ {
		t.Run(fmt.Sprint("Case ", i+1), func(t *testing.T) {
			assert.Equal(t, MaxSequence(theTestCase[i].Parameter), theTestCase[i].Expected,
				theTestCase[i].Message)
		})
	}

}
