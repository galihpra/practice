package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMostappearItem(t *testing.T) {
	type TestCase struct {
		Parameter []string
		Expected  string
		Message   string
	}

	theTestCase := []TestCase{
		{[]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}, "golang -> 1\nruby -> 2\njs -> 4\n", "Output tidak sesuai"},
		{[]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}, "C -> 1\nD -> 2\nB -> 3\nA -> 4\n", "Output tidak sesuai"},
	}

	for i := 0; i < len(theTestCase); i++ {
		t.Run(fmt.Sprint("Case ", i+1), func(t *testing.T) {
			// theRune := theTestCase[i].Parameter
			assert.Equal(t, theTestCase[i].Expected, MostappearItem(theTestCase[i].Parameter),
				theTestCase[i].Message)
		})
	}

}
