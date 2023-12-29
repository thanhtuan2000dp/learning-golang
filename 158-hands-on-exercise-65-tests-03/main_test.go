package main

import (
	"fmt"
	"testing"
)

func TestParadise(t *testing.T) {
	location := "New York"
	got := paradise(location)
	want := fmt.Sprint("My idea of paradise is ", location)
	if got != want {
		t.Errorf("Error - want:%s | got:%s.", want, got)
	}
}
