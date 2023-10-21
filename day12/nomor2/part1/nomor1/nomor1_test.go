package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrimeX(t *testing.T) {
	type TestCase struct {
		Parameter int
		Expected  int
		Message   string
	}

	theTestCase := []TestCase{
		{5, 11, "Output tidak sesuai"},
		{9, 23, "Output tidak sesuai"},
		{10, 29, "Output tidak sesuai"},
	}

	for i := 0; i < len(theTestCase); i++ {
		t.Run(fmt.Sprint("Case ", i+1), func(t *testing.T) {
			assert.Equal(t, theTestCase[i].Expected, PrimeX(theTestCase[i].Parameter),
				theTestCase[i].Message)
		})
	}

}
