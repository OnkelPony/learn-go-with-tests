package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Collection of variable size", func(t *testing.T) {

		numbers := []int{1, 2, 3, 4, 5, 6}
		got := Sum(numbers)
		want := 21

		if got != want {
			t.Errorf("Got %d want %d given %v", got, want, numbers)
		}
	})

	t.Run("Slice of slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9, 10})
		want := []int{3, 19}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v want %v", got, want)
		}
	})

}

func TestSumAllTails(t *testing.T) {
	checkSums := func(got []int, want []int, t *testing.T) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v want %v", got, want)
		}
	}
	t.Run("make the sum of some slices", func(t *testing.T) {

		got := SumAllTails([]int{1, 2}, []int{0, 9, 10})
		want := []int{2, 19}

		checkSums(got, want, t)
	})

	t.Run("safe sum of empty slices", func(t *testing.T) {

		got := SumAllTails([]int{}, []int{6, 6, 6})
		want := []int{0, 12}

		checkSums(got, want, t)
	})
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2}, []int{0, 9, 10})
	}
}

func BenchmarkSumAllTails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2}, []int{0, 9, 10})
	}
}
