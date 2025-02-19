package main

import "testing"


func BenchmarkStackIt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main1()
	}
}

func BenchmarkStackIt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main2()
	}
}


func BenchmarkVector(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = mainVector()
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = mainSlices()
	}
}