package main

import "fmt"

func main() {

	var s1 string = "👻"
	var s2 string = "🎃👻"
	var s3 string = "👽🎃👻"
	var st1 string = "a"

	fmt.Println([]byte(s1))
	fmt.Println(len(s1))

	fmt.Println([]byte(s1), []byte(st1))
	fmt.Println(len(s1) + len(st1))

	fmt.Println([]byte(s2))
	fmt.Println(len(s2))

	fmt.Println([]byte(s2), []byte(st1))
	fmt.Println(len(s2) + len(st1))

	fmt.Println([]byte(s3))
	fmt.Println(len(s3))

	fmt.Println([]byte(s3), []byte(st1))
	fmt.Println(len(s3) + len(st1))

	// var s1 string = "👻"
	// var s2 string = "🎃"
	// var s3 string = "👽"
	// var a1 string = "a"

	// fmt.Println([]byte(s1))
	// fmt.Println(len(s1))

	// fmt.Printf("%d, %d\n", []byte(s1), []byte(a1))
	// fmt.Printf("%d\n", len(s1)+len(a1))

	// fmt.Printf("%d, %d\n", []byte(s1), []byte(s2))
	// fmt.Printf("%d\n", len(s1)+len(s2))

	// fmt.Printf("%d, %d, %d\n", []byte(s1), []byte(s2), []byte(a1))
	// fmt.Printf("%d\n", len(s1)+len(s2)+len(a1))

	// fmt.Printf("%d, %d, %d\n", []byte(s1), []byte(s2), []byte(s3))
	// fmt.Printf("%d\n", len(s1)+len(s2)+len(s3))

	// fmt.Printf("%d, %d,%d, %d\n", []byte(s1), []byte(s2), []byte(s3), []byte(a1))
	// fmt.Printf("%d\n", len(s1)+len(s2)+len(s3)+len(a1))

}
