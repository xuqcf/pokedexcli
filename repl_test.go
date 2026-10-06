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
			t.Errorf("the length of the words is not matched")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord { // Check each word in the slice
				t.Errorf("The word: %s to do not match the expected word: %s", word, expectedWord) // if they don't match, use t.Errorf to print an error message
				// and fail the test
			}
		}
	}
}
