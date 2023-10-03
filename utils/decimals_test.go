package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/base-media-cloud/base-test/utils"
)

func TestToDecimalString(t *testing.T) {
	tests := map[string]struct {
		input    float64
		expected string
	}{
		"works when you pass in 0": {
			input:    0,
			expected: "0.00",
		},
		"works when you have a decimal to 2 places": {
			input:    1.23,
			expected: "1.23",
		},
		"works with a negative number": {
			input:    -1.23,
			expected: "-1.23",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := utils.ToDecimalString(test.input)

			assert.Equal(t, test.expected, result)
		})
	}
}
