package main

import "fmt"

func main() {

	var x = true
	var y = false

	fmt.Println(x && y == y && x)
	fmt.Println(!(x && !y) == !x || y == x && y || x || x)
	fmt.Println(x || y == x && !y || !x && y == (x || y) && (!x || !y))
	fmt.Println(x || (y && !x) == (x && y) || (x && !x))
	fmt.Println(x && (y || !x) == (x && y) || (x && !x))
	fmt.Println(!x && !y == !x || !y) 
	fmt.Println(x && (x || y) == x)
	fmt.Println(x || !x == x && y == x || x == y)
	fmt.Println(x || x == x && x == x && x == x || y == x)
	fmt.Println(!!x == x)
	fmt.Println(x && x == !(x && y) == !x || !y)

}
