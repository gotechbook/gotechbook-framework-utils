package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var sliceStr = []string{"a", "ab", "abc"}

func TestSliceContainsString(t *testing.T) {
	assert.Equal(t, true, sliceContainsString(sliceStr, "a"))
	assert.Equal(t, true, sliceContainsString(sliceStr, "ab"))
	assert.Equal(t, false, sliceContainsString(sliceStr, "abv"))
}
