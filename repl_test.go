package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HeLLo, World !",
			expected: []string{"hello,", "world", "!"},
		},
		{
			input:    "   ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)

		if len(c.expected) != len(actual) {
			t.Errorf(`---------------------------------
TEST FAILED
Different length
Expecting:  Error: %v
Actual:     Error: %v
Fail
`, len(c.expected), len(actual))

		} else {
			fmt.Printf(`---------------------------------
TEST PASSED
Matched length
Expecting:  Error: %v
Actual:     Error: %v
Pass
`, len(c.expected), len(actual))

		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if strings.Compare(word, expectedWord) != 0 {
				t.Errorf(`---------------------------------
TEST FAILED
Unmatched words
Expecting:  Error: %v
Actual:     Error: %v
Fail
				`, c.expected, actual)
			} else {
				fmt.Printf(`---------------------------------
TEST Passed
Expecting:  Error: %v
Actual:     Error: %v
Pass
`, c.expected, actual)
			}
		}
	}

	// ...
}
