package strings

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	string := "hello world"
	reversedString := ReverseString(string)
	if reversedString != "dlrow olleh" {
		t.Fatalf("String is reversed incorrectly: %s", reversedString)
	}
}
