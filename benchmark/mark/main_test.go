package benchmark

import (
	"fmt"
	"testing"
)

func TestMark(t *testing.T) {
	teststring := Mark("Demo")
	if teststring != "Target Demo has been marked" {
		t.Error("Got", teststring, "expected", "Target Demo has been marked\"")
	}
}

func ExampleMark() {
	fmt.Println(Mark("Demo"))
	//out put: Target Demo has been marked
}

func BenchmarkMark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mark("Demo")
	}
}
