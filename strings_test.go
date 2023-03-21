package strings

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{input: "hello world", expected: "dlrow olleh"},
	}
	for _, testCase := range testCases {
		t.Run(testCase.input, testReverseStringFunc(testCase.input, testCase.expected))
	}
}

func testReverseStringFunc(input string, expected string) func(*testing.T) {
	return func(t *testing.T) {
		actual := ReverseString(input)
		if actual != expected {
			t.Errorf(fmt.Sprintf("Expected reversed string to be %s but instead got %s!", expected, actual))
		}
	}
}
