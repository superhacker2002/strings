package strings

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{"Empty string", "", ""},
		{"Single character string", "a", "a"},
		{"Even length string", "hell", "lleh"},
		{"Odd length string", "world", "dlrow"},
		{"String with spaces", "hello world", "dlrow olleh"},
		{"Non ASCII string", "привет, мир", "рим ,тевирп"},
		{"String with escape sequence", "привет, мир\n", "\nрим ,тевирп"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := ReverseString(testCase.input)
			if actual != testCase.expectedOutput {
				t.Errorf("ReverseString(%q) = %q; expected %q", testCase.input, actual, testCase.expectedOutput)
			}
		})
	}
}

func TestCharactersCount(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		input          string
		expectedOutput int
	}{
		{"Empty string", "", 0},
		{"Single character string", "a", 1},
		{"Even length string", "hell", 4},
		{"Odd length string", "world", 5},
		{"String with spaces", "hello world", 11},
		{"Non ASCII string", "привет, мир", 11},
		{"String with escape sequence", "привет, мир\n", 12},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := CharactersCount(testCase.input)
			if actual != testCase.expectedOutput {
				t.Errorf("CharactersCount(%q) = %d; expected %d", testCase.input, actual, testCase.expectedOutput)
			}
		})
	}
}
