package kata

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected []string
	}{
		// Add your test cases here
		{"1 pair", 1, []string{"()"}},
		{"3 pairs", 3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := generateParenthesis(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("got %v want %v", result, tc.expected)
			}
		})
	}
}
