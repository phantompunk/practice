package kata

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected [][]string
	}{
		{"Example 1", []string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}},
		{"Empty Input", []string{}, [][]string{}},
		{"Single Element", []string{"a"}, [][]string{{"a"}}},
		{"No Anagrams", []string{"abc", "def", "ghi"}, [][]string{{"abc"}, {"def"}, {"ghi"}}},
		{"All Anagrams", []string{"abc", "bca", "cab"}, [][]string{{"abc", "bca", "cab"}}},
		{"Mixed Case", []string{"Eat", "Tea", "ate"}, [][]string{{"Eat", "Tea", "ate"}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := groupAnagrams(tc.input)
      if !reflect.DeepEqual(result, tc.expected) {
        t.Errorf("got %v, want %v", result, tc.expected)
      }
		})
	}
}
