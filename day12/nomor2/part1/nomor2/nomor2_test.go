package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	type TestCase struct {
		Parameter int
		Expected  int
		Message   string
	}

	theTestCase := []TestCase{
		{0, 0, "Output tidak sesuai"},
		{4, 3, "Output tidak sesuai"},
		{12, 144, "Output tidak sesuai"},
	}

	for i := 0; i < len(theTestCase); i++ {
		t.Run(fmt.Sprint("Case ", i+1), func(t *testing.T) {
			assert.Equal(t, theTestCase[i].Expected, Fibonacci(theTestCase[i].Parameter),
				theTestCase[i].Message)
		})
	}

}
