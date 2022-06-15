package main

import (
	"testing"
)

func TestExistingDirectoryShouldSuccess(t *testing.T) {
	result := isExists(`C:/Users/ikiru.yoshizaki/go/src`)
	if !result {
		t.Fatal("failed test")
	}
}

func TestNotExistingDirectoryShouldFailed(t *testing.T) {
	result := isExists(`C:/Users/ikiru.yoshizaki/go/src/hoge`)
	if result {
		t.Fatal("failed test")
	}
}
