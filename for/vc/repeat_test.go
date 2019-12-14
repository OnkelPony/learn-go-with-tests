package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 9)
	want := "aaaaaaaaa"

	if got != want {
		t.Errorf("expected %q but got %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 9)
	}
}

func ExampleRepeat() {
	result := Repeat("6", 3)
	fmt.Println(result)
	// Output: 666
}
