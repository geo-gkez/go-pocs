package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "George")

	got := buffer.String()
	want := "Hello, George"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
