package main

import (
	"fmt"
	"runtime"
)

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MB\n", bToMb(m.Alloc)) // current allocated memory
	fmt.Printf("TotalAlloc = %v MB\n", bToMb(m.TotalAlloc)) // cumulative total allocated memory (only increases, never decreases)
	fmt.Printf("Sys = %v MB\n", bToMb(m.Sys)) // total memory allocated obtained from the OS (includes heap, stack, etc.)
	fmt.Printf("NumGC = %v\n", m.NumGC) // number of garbage collections
}

func bToMb(b uint64) uint64 {
	return b / 1000 / 1000
}

func main() {
	fmt.Println("Mem stats before:")
	printMemStats()

	s := make([]int, 10_000_000)
	for i := 0; i < len(s); i++ {
		s[i] = i
	}

	fmt.Println()
	fmt.Println("Mem stats after:")
	printMemStats()

	runtime.GC() // never use it in production

	fmt.Println()
	fmt.Println("Mem stats after GC, we'll use 's' after GC:")
	printMemStats()

	// if I use 's' here, the manual GC will not remove it from memory
	// because it's still in the stack
	fmt.Println(s[1])

	runtime.GC() // never use it in production

	fmt.Println()
	fmt.Println("Mem stats after GC:")
	printMemStats()
}