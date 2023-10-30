package calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyString(t *testing.T) {
	res := Add("")
	assert.Equal(t, res, int64(0))
}

func TestOneNumber(t *testing.T) {
	res := Add("1")
	assert.Equal(t, res, int64(1))
}

func TestTwoNumbers(t *testing.T) {
	res := Add("1,2")
	assert.Equal(t, res, int64(3))
	assert.Equal(t, int64(10), Add("1,2,3,4"))
	assert.Equal(t, int64(2), Add("1,2,3,-4"))
}

func TestManualDelimiter(t *testing.T) {
	res := Add("1:2")
	assert.Equal(t, res, int64(3))
	assert.Equal(t, int64(10), Add("1,2,3:4"))
	assert.Equal(t, int64(2), Add("1,2:3,-4"))
}
