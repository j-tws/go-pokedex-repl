package main

import "testing"

type context struct {
	input string
	expected []string
}

func TestCleanInputBasic(t *testing.T) {
	cases := []context{
		{
			input: "  hello world   ",
			expected: []string{"hello", "world"},
		},
		{
			input: "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected){
			t.Errorf("Got length of cleanInput(%s) = %d, expect: %d", c.input, len(actual), len(c.expected))
		}

		for i, actualWord := range actual {
			expectedWord := c.expected[i]

			if actualWord != expectedWord {
				t.Errorf("Expected word to be '%s', but got '%s' instead", expectedWord, actualWord)
			}
		}
	}
}
