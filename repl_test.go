package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  bulbasaur     pikachu  ",
			expected: []string{"bulbasaur", "pikachu"},
		},
		{
			input:    "  mewto      charizard     squirtle  ",
			expected: []string{"mewto", "charizard", "squirtle"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf(
				"Length of actual, %d is not what is expected, %d.",
				len(actual),
				len(c.expected),
			)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Actual word does not match expected word.")
			}
		}

	}
}
