package main

import "testing"

func TestHello(t string) {
	want := "Hello, World."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want $q", got, want)
	}
}

func main() {
	TestHello("test")
}

//5/07/2020 Come back to this and fix it.
