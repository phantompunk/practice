package kata

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
  testCases := []struct {
    name     string
    input    string
    expected bool
  }{
    // Add your test cases here
    {"empty string", "", true},
    {"single character", "a", true},
    {"two same characters", "aa", true},
    {"two different characters", "ab", false},
    {"palindrome even length", "abba", true},
    {"palindrome odd length", "racecar", true},
    {"not a palindrome", "hello", false},
    {"palindrome with spaces", "A man a plan a canal Panama", true},
    {"palindrome with punctuation", "Madam, I'm Adam.", true},
    {"not a palindrome with punctuation", "Hello, World!", false},
  }

  for _, tc := range testCases {
    t.Run(tc.name, func(t *testing.T) {
      result := isPalindrome(tc.input)
      assert.Equal(t, tc.expected, result)
    })
  }
}
