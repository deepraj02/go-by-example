package main

import (
	"testing"
)

func TestLoopSingleCondition(t *testing.T) {
	repeated := LoopSingleCondition("A")
	expected := "AAAAA"
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}

}

func BenchmarkLoopSingleCondition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoopSingleCondition("a")
	}
}

func TestLoopInitialCondition(t *testing.T) {
	total := LoopInitialCondition(5)
	expected := 15
	if total != expected {
		t.Errorf("expected %d but got %d", expected, total)
	}

}
func BenchmarkLoopInitialCondition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoopInitialCondition(5)
	}
}

func TestLoopRange(t *testing.T) {
	nums := LoopRange(1, 2, 3, 4, 5)
	expected := []int{1, 2, 3, 4, 5}
	for i, num := range nums {
		if num != expected[i] {
			t.Errorf("expected %d but got %d", expected[i], num)
		}
	}
}

func BenchmarkLoopRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoopRange(1, 2, 3, 4, 5)
	}
}
