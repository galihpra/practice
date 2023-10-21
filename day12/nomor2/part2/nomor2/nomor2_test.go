package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumBuyProduct(t *testing.T) {
	type TestCase struct {
		Parameter1 int
		Parameter2 []int
		Expected   int
		Message    string
	}

	theTestCase := []TestCase{
		{30000, []int{15000, 10000, 12000, 5000, 3000}, 4, "Output tidak sesuai"},
		{10000, []int{2000, 3000, 1000, 2000, 10000}, 4, "Output tidak sesuai"},
		{0, []int{10000, 30000}, 0, "Output tidak sesuai"},
	}

	for i := 0; i < len(theTestCase); i++ {
		t.Run(fmt.Sprint("Case ", i+1), func(t *testing.T) {
			assert.Equal(t, theTestCase[i].Expected, MaximumBuyProduct(theTestCase[i].Parameter1, theTestCase[i].Parameter2),
				theTestCase[i].Message)
		})
	}

}
