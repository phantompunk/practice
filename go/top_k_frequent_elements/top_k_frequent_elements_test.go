package kata

import (
        "testing"

        "github.com/stretchr/testify/assert"
)

func TestTopKFrequent(t *testing.T) {
        testCases := []struct {
                name     string
                input    string
                expected string
        }{
                // Add your test cases here
        }

        for _, tc := range testCases {
                t.Run(tc.name, func(t *testing.T) {
                        result := topKFrequent(tc.input)
                        assert.Equal(t, tc.expected, result)
                })
        }
}
