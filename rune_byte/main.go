package main

import "fmt"

func main() {

	// var s4 string = "👻"
	// var s2 string = "🎃👻"
	// var s3 string = "👽🎃👻"
	// var st1 string = "a"

	// fmt.Println([]byte(s4))
	// fmt.Println(len(s4))

	// fmt.Println([]byte(s4), []byte(st1))
	// fmt.Println(len(s4) + len(st1))

	// fmt.Println([]byte(s2))
	// fmt.Println(len(s2))

	// fmt.Println([]byte(s2), []byte(st1))
	// fmt.Println(len(s2) + len(st1))

	// fmt.Println([]byte(s3))
	// fmt.Println(len(s3))

	// fmt.Println([]byte(s3), []byte(st1))
	// fmt.Println(len(s3) + len(st1))

	var s1 string = "a"  //1
	var s2 string = "я"  //2
	var s3 string = "字"  //3
	var s4 string = "👻"  //4
	var s6 string = "❄️" //6
	var s7 string = "🏳️" //7

	fmt.Println([]byte(s4))
	fmt.Println(len(s4)) //4

	fmt.Printf("%d, %d\n", []byte(s2), []byte(s2))
	fmt.Printf("%d\n", len(s2)+len(s2)) //4

	fmt.Printf("%d, %d\n", []byte(s4), []byte(s1))
	fmt.Printf("%d\n", len(s4)+len(s1)) //5

	fmt.Printf("%d, %d\n", []byte(s4), []byte(s2))
	fmt.Printf("%d\n", len(s4)+len(s2)) //6

	fmt.Printf("%d, %d\n", []byte(s6), []byte(s1))
	fmt.Printf("%d\n", len(s6)+len(s1)) //7

	fmt.Printf("%d, %d\n", []byte(s6), []byte(s2))
	fmt.Printf("%d\n", len(s6)+len(s2)) //8

	fmt.Printf("%d, %d\n", []byte(s6), []byte(s3))
	fmt.Printf("%d\n", len(s6)+len(s3)) //9

	fmt.Printf("%d, %d\n", []byte(s6), []byte(s4))
	fmt.Printf("%d\n", len(s6)+len(s4)) //10

	fmt.Printf("%d, %d\n", []byte(s7), []byte(s4))
	fmt.Printf("%d\n", len(s7)+len(s4)) //11

	fmt.Printf("%d, %d, %d\n", []byte(s7), []byte(s3), []byte(s2))
	fmt.Printf("%d\n", len(s7)+len(s3)+len(s2)) //12

}
