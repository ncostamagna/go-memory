package main

import "testing"

func BenchmarkPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mainPointer()
	}
}

// uses heap, less performance 
func BenchmarkWeakPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mainWeak()
	}
}
