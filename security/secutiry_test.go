package security

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateToken(t *testing.T) {
	res, err := createToken("admin")
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestVerifyHappy(t *testing.T) {
	res, err := createToken("admin")
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	err = verifyToken(res)
}
