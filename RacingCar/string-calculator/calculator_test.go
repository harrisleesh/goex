package string_calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLibrary(t *testing.T) {
	assert.Equal(t, "a", "a")
}

func TestCalculatorBinary(t *testing.T) {
	assert.Equal(t, Calculate("1 + 2"), int64(3))
	assert.Equal(t, Calculate("1 - 2"), int64(-1))
	assert.Equal(t, Calculate("2 * 2"), int64(4))
	assert.Equal(t, Calculate("10 / 2"), int64(5))
}

func TestCalculatorMultiple(t *testing.T) {
	assert.Equal(t, Calculate("1 + 2 - 1"), int64(2))
}
