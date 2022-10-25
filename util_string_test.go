package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var sliceStr = []string{"a", "ab", "abc"}

func TestSliceContainsString(t *testing.T) {
	assert.Equal(t, true, SliceContainsString(sliceStr, "a"))
	assert.Equal(t, true, SliceContainsString(sliceStr, "ab"))
	assert.Equal(t, false, SliceContainsString(sliceStr, "abv"))
}
