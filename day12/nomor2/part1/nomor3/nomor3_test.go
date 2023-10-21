package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrimaSegiEmpat(t *testing.T) {
	type TestCase struct {
		Parameter []int
		Expected  string
		Message   string
	}

	theTestCase := []TestCase{
		{[]int{2, 3, 13}, "17 19 \n23 29 \n31 37 \n156", "Output tidak sesuai"},
		{[]int{5, 2, 1}, "2 3 5 7 11 \n13 17 19 23 29 \n129", "Output tidak sesuai"},
	}

	for i := 0; i < len(theTestCase); i++ {
		t.Run(fmt.Sprint("Case ", i+1), func(t *testing.T) {
			theRune := theTestCase[i].Parameter
			assert.Equal(t, theTestCase[i].Expected, PrimaSegiEmpat(theRune[0], theRune[1], theRune[2]),
				theTestCase[i].Message)
		})
	}

}
