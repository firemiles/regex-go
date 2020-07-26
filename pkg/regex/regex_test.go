package regex

import (
	"testing"
)

func TestSuffixExpression(t *testing.T) {
	var testCase []string = []string{
		"a|b",
	}

	for _, tc := range testCase {
		se, _ := SuffixExpression(tc)
		t.Errorf(string(se))
	}
}
