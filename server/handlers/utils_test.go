package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatePassword(t *testing.T) {
	testCases := []struct {
		name                string
		password            string
		expectErrorContains string
	}{
		{
			name:                "invalid length",
			password:            "this",
			expectErrorContains: "expected password to be a minimum of 10 characters",
		},
		{
			name:                "no uppercase",
			password:            "this is a test",
			expectErrorContains: "at least one uppercase letter is required",
		},
		{
			name:                "no lowercase",
			password:            "THISISATEST",
			expectErrorContains: "at least one lowercase letter is required",
		},
		{
			name:                "no digits",
			password:            "thisIsATest",
			expectErrorContains: "at least one number is required",
		},
		{
			name:     "happy",
			password: "Thisisatest123",
		},
	}
	for _, tt := range testCases {
		t.Run(
			tt.name,
			func(t *testing.T) {
				err := validatePassword(tt.password)
				if len(tt.expectErrorContains) > 0 {
					assert.ErrorContains(t, err, tt.expectErrorContains)
					return
				}
				assert.Nil(t, err)
			},
		)
	}
}
