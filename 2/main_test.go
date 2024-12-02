package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSafeLine(t *testing.T) {
	assert.Equal(t, true, safeLine([]int{1, 3, 5, 6}))
	assert.Equal(t, true, safeLine([]int{-1, 1, 3, 5}))
	assert.Equal(t, true, safeLine([]int{-1, -2, -3, -5}))
	assert.Equal(t, false, safeLine([]int{1, 0, 1, 2}))
	assert.Equal(t, false, safeLine([]int{0, -1, 1, 2}))
	assert.Equal(t, false, safeLine([]int{0, 1, -1, -2}))
	assert.Equal(t, false, safeLine([]int{1, -1, 5, 6}))
	assert.Equal(t, false, safeLine([]int{1, 1, 2, 3}))
}
