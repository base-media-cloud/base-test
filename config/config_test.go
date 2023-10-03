package config_test

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/base-media-cloud/base-test/config"
)

var allEnvVariables = []string{
	"HOST",
	"DELIVERY_SERVICE",
	"PRICE",
	"READ_HEADER_TIMEOUT",
}

func resetEnvVariables() {
	for _, v := range allEnvVariables {
		err := os.Unsetenv(v)
		if err != nil {
			panic(err)
		}
	}
}

func TestNewConfig(t *testing.T) {
	tests := map[string]struct {
		envVars     map[string]string
		expected    *config.Config
		expectedErr error
	}{
		"default values": {
			envVars: map[string]string{
				"HOST":             "mock-host",
				"DELIVERY_SERVICE": "mock-service",
				"PRICE":            "0.01",
			},
			expected: &config.Config{
				Host:              "mock-host",
				DeliverySvc:       "mock-service",
				Price:             0.01,
				ReadHeaderTimeout: 15 * time.Second,
			},
		},
		"all values": {
			envVars: map[string]string{
				"HOST":                "mock-host",
				"DELIVERY_SERVICE":    "mock-service",
				"PRICE":               "0.01",
				"READ_HEADER_TIMEOUT": "1s",
			},
			expected: &config.Config{
				Host:              "mock-host",
				DeliverySvc:       "mock-service",
				Price:             0.01,
				ReadHeaderTimeout: 1 * time.Second,
			},
		},
		"missing required values": {
			envVars:     map[string]string{"READ_HEADER_TIMEOUT": "1s"},
			expectedErr: errors.New("Host: missing required value: HOST"),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			resetEnvVariables()

			for k, v := range test.envVars {
				err := os.Setenv(k, v)
				if err != nil {
					t.Fatal(err)
				}
			}

			result, err := config.New()

			assert.Equal(t, test.expected, result)
			if err != nil {
				assert.Equal(t, test.expectedErr.Error(), err.Error())
			}
		})
	}
}
