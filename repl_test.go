package main

import "testing"

func TestCleanInput(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HI                 TherE      ",
			expected: []string{"hi", "there"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, tc := range tests {
		actual := cleanInput(tc.input)

		if len(actual) != len(tc.expected) {
			t.Errorf("wanted slice = %#v, actual slice = %#v\n", tc.expected, actual)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := tc.expected[i]

			if word != expectedWord {
				t.Errorf("expected word = %s, actual word = %s\n", expectedWord, word)
			}
		}
	}
}
