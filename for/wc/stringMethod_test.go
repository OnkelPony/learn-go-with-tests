package wc

import (
	"fmt"
	"testing"
)

const text = "Let all beings be happy!"

func TestReplace(t *testing.T) {
	got := ReplaceAndUpper(text, "e", "I", 2)
	want := "LIT ALL BIINGS BE HAPPY!"

	if got != want {
		t.Errorf("Expected %q but got %q", want, got)
	}

}

func ExampleReplaceAndUpper() {
	result := ReplaceAndUpper("Neumíme, jdi pryč!", "m", "gh", -1)
	fmt.Println(result)
	//Output: NEUGHÍGHE, JDI PRYČ!
}

func BenchmarkReplaceAndUpper(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReplaceAndUpper("A poslední čtyři se jmenovali Root", "i", "y", 2)
	}
}
