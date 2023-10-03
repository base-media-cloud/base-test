package yodal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/base-media-cloud/base-test/internal/adapters/delivery/yodal"
)

func TestYodal_Price(t *testing.T) {
	tests := map[string]struct {
		grams    float64
		ppGram   float64
		expected float64
	}{
		"calculates the price correctly": {
			grams:    5.00,
			ppGram:   5.00,
			expected: 25.00,
		},
		"calculates the price correctly with 0": {
			grams:    0.00,
			ppGram:   5.00,
			expected: 0.00,
		},
		"calculates the price correctly with minus numbers": {
			grams:    -5.00,
			ppGram:   5.00,
			expected: -25.00,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			delSvc := yodal.New(test.ppGram)

			result := delSvc.Price(test.grams)

			assert.Equal(t, test.expected, result)
		})
	}
}
