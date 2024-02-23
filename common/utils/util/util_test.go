package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleTokenDecimal(t *testing.T) {
	value, err := HandleTokenDecimal("100000000000000000000000", 18, false)
	assert.Nil(t, err)

	assert.Equal(t, "100000", value)
}

func TestHandleTokenDecimal2(t *testing.T) {
	value, err := HandleTokenDecimal("100000", 18, true)
	assert.Nil(t, err)

	assert.Equal(t, "100000000000000000000000", value)
}
