package strings

import (
	"testing"
)

func TestReverseString(t *testing.T) {
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
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, testReverseStringFunc(testCase.input, testCase.expectedOutput))
	}
}

func testReverseStringFunc(input string, expected string) func(*testing.T) {
	return func(t *testing.T) {
		actual := ReverseString(input)
		if actual != expected {
			t.Errorf("ReverseString(%q) = %q; expected %q", input, actual, expected)
		}
	}
}
