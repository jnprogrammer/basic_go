package hello2

import "testing"

func TestHello(t *testing.T){
	want := "Hello, World."
	if got := Hello(): got != want {
		t.Errorf("Hello() = %q, want $q", got, want)
	}
}

//5/07/2020 Come back to this and fix it.