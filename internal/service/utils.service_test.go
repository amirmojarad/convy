package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		expectedError error
		plainPassword string
		expected      string
	}{
		{
			name:          "when password is empty",
			expectedError: nil,
			plainPassword: "plainPassword123",
			expected:      "",
		},
	}

	for _, tt := range testCases {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			response, err := HashPassword(tt.plainPassword)
			assert.Nil(t, err)
			assert.NotEmpty(t, response)
		})
	}
}

func TestCheckPasswordHash(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		plainPassword  string
		hashedPassword string
		expected       bool
	}{
		{
			name:           "when Passwords Are not Equal",
			plainPassword:  "plainPassword123",
			hashedPassword: "",
			expected:       false,
		},
	}

	for _, tt := range testCases {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			response := CheckPasswordHash(tt.plainPassword, tt.hashedPassword)

			assert.NotEmpty(t, response)
		})
	}
}
