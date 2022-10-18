package main

import "testing"

func sumMaxInt(m int) int {
	sum := 0
	for i := 0; i < m; i++ {
		sum += i
	}
	return sum
}
func sumWithDefer() {
	defer func() {
		sumMaxInt(10)
	}()
}

func sumWithoutDefer() {
	sumMaxInt(10)
}
func BenchmarkWithSumWithDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumWithDefer()
	}
}
func BenchmarkWithoutSumWithDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumWithoutDefer()
	}
}
