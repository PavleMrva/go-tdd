package di

import (
	"testing"
	"bytes"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Pavle")

	got := buffer.String()
	want := "Hello, Pavle"

	if got != want {
		t.Errorf("got %q want %q", got, want)	
	}
}