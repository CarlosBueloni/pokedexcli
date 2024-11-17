package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "Hello wOrld",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "  hEllo  ",
			expected: []string{
				"hello",
			},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("The lengths are not equal: '%v' vs '%v'",
				len(actual),
				len(cs.expected),
			)
			continue
		}

		for i := range actual {
			actual_word := actual[i]
			expected_word := cs.expected[i]
			if actual_word != expected_word {
				t.Errorf("The words are not equal: '%v' vs '%v'",
					actual_word,
					expected_word,
				)
			}
		}

	}
}
