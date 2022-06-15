package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsExists(t *testing.T) {
	assert.True(t, isExists(`C:/Users/ikiru.yoshizaki/go/src`), "Directory exists")
	assert.False(t, isExists(`C:/Users/ikiru.yoshizaki/go/src/hoge`), "Directory not exists")
}
