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
		{"Even length string", "hello", "olleh"},
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
