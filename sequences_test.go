package factorize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateFibonacci(t *testing.T) {
	expected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	actual := GenerateFibonacci(10)

	assert.Equal(t, expected, actual)
}

func TestGetNthFibonacci(t *testing.T) {
	expected := 34
	actual := GetNthFibonacci(10)

	assert.Equal(t, expected, actual)
}
