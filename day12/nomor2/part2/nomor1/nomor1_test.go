package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMinAndMax(t *testing.T) {
	type TestCase struct {
		Parameter []int
		Expected  string
		Message   string
	}

	theTestCase := []TestCase{
		{[]int{5, 7, 4, -2, -1, 8}, "min: -2, index: 3,max: 8,index: 5", "Output tidak sesuai"},
		{[]int{4, 3, 9, 4, -21, 7}, "min: -21, index: 4,max: 9,index: 2", "Output tidak sesuai"},
		{[]int{-1, 5, 6, 4, 2, 18}, "min: -1, index: 0,max: 18,index: 5", "Output tidak sesuai"},
	}

	for i := 0; i < len(theTestCase); i++ {
		t.Run(fmt.Sprint("Case ", i+1), func(t *testing.T) {
			assert.Equal(t, theTestCase[i].Expected, FindMinAndMax(theTestCase[i].Parameter),
				theTestCase[i].Message)
		})
	}

}
