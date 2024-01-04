package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	var buffer bytes.Buffer

	Greet(&buffer, "User")

	got := buffer.String()
	want := "Hello, User"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
