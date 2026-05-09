package main

import (
	"slices"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello    world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "egg head123, beach",
			expected: []string{"egg", "head123,", "beach"},
		},
		{
			input:    " THIS IS FULL CAPS ",
			expected: []string{"this", "is", "full", "caps"},
		},
		{
			input:    "         ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if slices.Compare(actual, c.expected) != 0 {
			t.Errorf("cleanInput(%q) == %v, expected %v", c.input, actual, c.expected)
			return
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("The give word %q does not match the expected word %q", word, expectedWord)
				return
			}
		}
	}
}
