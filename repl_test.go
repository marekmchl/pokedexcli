package main

import "testing"

// TestCleanInput repeatedly calls cleanInput with a test cases checking for an error.
func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    " hello ",
			expected: []string{"hello"},
		},
		{
			input:    " hello there    world  ",
			expected: []string{"hello", "there", "world"},
		},
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    " hello ",
			expected: []string{"hello"},
		},
		{
			input: `   hello
			there
			  world  `,
			expected: []string{"hello", "there", "world"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Failed testing cleanInput\nExpected words: %v\nGot words: %v\n", c.expected, actual)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Failed in cleanInput test case %v\nExpected word: %v\nGot word: %v\n", c.input, expectedWord, word)
			}
		}
	}
}
