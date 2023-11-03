package service

import (
	"convy/internal/errorext"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidationBuilder_Validate(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		input          string
		expectedResult bool
		expectedError  error
		flow           func(value string) (bool, error)
	}{
		{
			name:           "when password is invalid",
			expectedResult: false,
			input:          "pass",
			expectedError:  errorext.NewValidationError("password is invalid"),
			flow: func(value string) (bool, error) {
				return NewValidation().SetPassword(value).Validate()
			},
		},
		{
			name:           "when username is invalid",
			expectedResult: false,
			input:          "user",
			expectedError:  errorext.NewValidationError("username is invalid"),
			flow: func(value string) (bool, error) {
				return NewValidation().SetUsername(value).Validate()
			},
		},
		{
			name:           "when email is in valid",
			expectedResult: false,
			input:          "email",
			expectedError:  errorext.NewValidationError("email is invalid"),
			flow: func(value string) (bool, error) {
				return NewValidation().SetEmail(value).Validate()
			},
		},
	}

	for _, tt := range testCases {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, err := tt.flow(tt.input)

			assert.NotNil(t, err)
			assert.Equal(t, tt.expectedError.Error(), err.Error())
			assert.Equal(t, tt.expectedResult, result)
		})
	}
}
