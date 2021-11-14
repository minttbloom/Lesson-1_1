package main

import "fmt"

func main() {

	var r rune
	var r1 rune
	var r2 rune
	var a byte
	var b []byte
	var c []byte

	r = '😂'
	r1 = '😂'
	r2 = '😁'
	a = '1'
	b = []byte("12d")
	c = []byte("3edkiw")

	fmt.Println(r)
	fmt.Println(r1, r2)
	fmt.Println(r, r1, r2)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
