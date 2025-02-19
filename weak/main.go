package main

import (
	"fmt"
	"runtime"
	"weak"
)

type Data struct {
	value string
}

func main() {
	wp := weak.Make(&Data{value: "Temporary Data"})
	p := &Data{value: "Temporary Data"}
	fmt.Println("Weak Pointer Value:", wp.Value())
	fmt.Println("Pointer Value:", p)
	runtime.GC()                                             // Trigger garbage collection
	fmt.Println("After GC, Weak Pointer Value:", wp.Value()) // Likely nil if GC has run
	fmt.Println("After GC, Pointer Value:", p)
}

func mainWeak() {
	_ = weakPointer()
}

func mainPointer() {
	_ = pointer()
}

//go:noinline
func weakPointer() weak.Pointer[Data] {
	return weak.Make(&Data{value: "Temporary Data"})
}

//go:noinline
func pointer() *Data {
	return &Data{
		value: "12113",
	}
}
