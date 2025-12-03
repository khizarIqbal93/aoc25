package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidV2(t *testing.T) {
	cases := []struct {
		input    int
		expected bool
	}{
		{
			input:    11,
			expected: false,
		},
		{
			input:    222222,
			expected: false,
		},
		{
			input:    1188511885,
			expected: false,
		},
		{
			input:    824824824,
			expected: false,
		},
		{
			input:    2121212121,
			expected: false,
		},
	}

	for _, tc := range cases {
		actual := IsValidV2(tc.input)
		assert.Equal(t, tc.expected, actual, fmt.Sprintf("failed for %v", tc.input))
	}
}
