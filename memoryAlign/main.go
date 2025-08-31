package main

import (
    "fmt"
    "unsafe"
)

// Less efficient (24 bytes)
type BadLayout struct {
    a int8   // 1 byte + 7 padding
    b int64  // 8 bytes
    c int8   // 1 byte + 7 padding
}

// More efficient (16 bytes)
type GoodLayout struct {
    a int64  // 8 bytes
    b int8   // 1 byte
    c int8   // 1 byte + 6 padding
}

// Bad layout with string
type BadLayoutWithString struct {
	a int8   // 1 byte + 7 padding
	b string // 16 bytes (8-byte aligned)
	c int64  // 8 bytes
	d int8   // 1 byte + 7 padding
}

// Better layout with string
type GoodLayoutWithString struct {
	a string // 16 bytes (8-byte aligned)
	b int64  // 8 bytes
	c int8   // 1 byte
	d int8   // 1 byte + 6 padding
}

func main() {
    var (
		bad BadLayout
		good GoodLayout
		badS BadLayoutWithString
		goodS GoodLayoutWithString
	)
	
    fmt.Printf("Bad Struct - Size: %d bytes\n", unsafe.Sizeof(bad))
    fmt.Printf("Bad Struct - Alignment: %d bytes\n", unsafe.Alignof(bad))
    
    // Field offsets
    fmt.Printf("Bad Struct - Field a offset: %d\n", unsafe.Offsetof(bad.a))
    fmt.Printf("Bad Struct - Field b offset: %d\n", unsafe.Offsetof(bad.b))
    fmt.Printf("Bad Struct - Field c offset: %d\n", unsafe.Offsetof(bad.c))

	fmt.Println()

	fmt.Printf("Good Struct - Size: %d bytes\n", unsafe.Sizeof(good))
    fmt.Printf("Good Struct - Alignment: %d bytes\n", unsafe.Alignof(good))
    
    // Field offsets
    fmt.Printf("Good Struct - Field a offset: %d\n", unsafe.Offsetof(good.a))
    fmt.Printf("Good Struct - Field b offset: %d\n", unsafe.Offsetof(good.b))
    fmt.Printf("Good Struct - Field c offset: %d\n", unsafe.Offsetof(good.c))

	fmt.Println()
	fmt.Println()

	fmt.Printf("Bad String Struct - Size: %d bytes\n", unsafe.Sizeof(badS))
    fmt.Printf("Bad String Struct - Alignment: %d bytes\n", unsafe.Alignof(badS))

	// Field offsets
	fmt.Printf("Bad String Struct - Field a offset: %d\n", unsafe.Offsetof(badS.a))
	fmt.Printf("Bad String Struct - Field b offset: %d\n", unsafe.Offsetof(badS.b))
	fmt.Printf("Bad String Struct - Field c offset: %d\n", unsafe.Offsetof(badS.c))
	fmt.Printf("Bad String Struct - Field c offset: %d\n", unsafe.Offsetof(badS.d))

	fmt.Println()

	fmt.Printf("Good String Struct - Size: %d bytes\n", unsafe.Sizeof(goodS))
	fmt.Printf("Good String Struct - Alignment: %d bytes\n", unsafe.Alignof(goodS))

	// Field offsets
	fmt.Printf("Good Struct - Field a offset: %d\n", unsafe.Offsetof(goodS.a))
	fmt.Printf("Good Struct - Field b offset: %d\n", unsafe.Offsetof(goodS.b))
	fmt.Printf("Good Struct - Field c offset: %d\n", unsafe.Offsetof(goodS.c))
	fmt.Printf("Good Struct - Field c offset: %d\n", unsafe.Offsetof(goodS.d))
}
