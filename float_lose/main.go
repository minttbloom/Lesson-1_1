package main

import "fmt"

func main() {

	var f64 float64 = 1
	var i64min float64 = -9223372036854775808
	var i64max float64 = 9223372036854775807

	fmt.Println(f64+i64max/i64min == 0)
	fmt.Println(f64+i64min/i64max == 1)

	var f32 float32 = 1
	var i32min float32 = -2147483648
	var i32max float32 = 2147483647

	fmt.Println(f32+i32max/i32min == 0)
	fmt.Println(f32+i32min/i32max == 1)

}
