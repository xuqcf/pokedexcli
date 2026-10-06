package main

import (
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
			input:    "  my name is Abdullah  ",
			expected: []string{"my", "name", "is", "abdullah"},
		},
		{
			input:    "  weather is sunny  ",
			expected: []string{"weather", "is", "sunny"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice
		// if they don't match, use t.Errorf and continue to the next case
		if len(actual) != len(c.expected) {
			t.Errorf("length mismatch: got %d words, expected %d", len(actual), len(c.expected))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord { // Check each word in the slice
				t.Errorf("input: %q: got %q, expected %q", c.input, word, c.expected) // if they don't match, use t.Errorf to print an error message
				// and fail the test
			}
		}
	}
}
