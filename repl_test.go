package main

import (
	"testing"

	"github.com/boxy-pug/pokedexcli/utils"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello  world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "SMAll anD Big",
			expected: []string{"small", "and", "big"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}
	for _, c := range cases {
		actual := utils.CleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("For input %q, expected len %d got: %d", c.input, len(c.expected), len(actual))
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("For input %q, expected %q at index %d, got %q", c.input, expectedWord, i, word)
			}

		}
	}
}
