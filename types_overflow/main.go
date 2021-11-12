package main

import "fmt"

func main() {

	var i8min int8 = -128
	var i8max int8 = 127

	fmt.Println(i8max + 100)
	fmt.Println(i8min-100 == i8max-99)
	fmt.Println(i8max-127 == i8min)
	fmt.Println(i8max-(i8max+i8min) == i8min)

	var a int16 = 10000
	fmt.Println(a * 4)

	var i32min int32 = -2147483648
	fmt.Println(i32min - 1)

	var i64min int64 = -9223372036854775808
	var i64max int64 = 9223372036854775807
	fmt.Println(i64min / i64max)
	fmt.Println(i64max + i64min*2)

	var r rune
	r = '😂'
	fmt.Println(r)

	var r1 rune
	var r2 rune
	r1 = '😂'
	r2 = '😁'

	fmt.Println(r1, r2)

}
