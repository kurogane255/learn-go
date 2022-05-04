package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	s := "hello world"
	result := Hello(s)
	if result != s {
		t.Fatalf("feaild test: %#v", result)
	}
}
