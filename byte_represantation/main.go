package main

import "fmt"

func main() {

	var a int = 180
	fmt.Printf("Print %b\n", a)

	// z := 1 << 1
	// az := a ^ z
	// fmt.Printf("Second byte %b\n", az)
	// fmt.Printf("Second byte check %b\n", az&z)

	// y := 1 << 3
	// ay := a ^ y
	// fmt.Printf("Fourth byte %b\n", ay)
	// fmt.Printf("Fourth byte check %b\n", ay&y)

	// x := 1 << 6
	// ax := a ^ x
	// fmt.Printf("Seventh byte %b\n", ax)
	// fmt.Printf("Seventh byte check %b\n", ax&x)

	z := 1 << 1
	y := 1 << 3
	x := 1 << 6
	a = a ^ z ^ x ^ y
	fmt.Printf("Byte representation %b", a)

}
