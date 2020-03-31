package validation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateEmail(t *testing.T) {
	assert.True(t, ValidateEmail("google@gmail.com"))
	assert.False(t, ValidateEmail("google@gmail@.com"))
	assert.False(t, ValidateEmail("google@gmail_com"))
	assert.False(t, ValidateEmail("ç$?§/az@gmail.comg"))
}
