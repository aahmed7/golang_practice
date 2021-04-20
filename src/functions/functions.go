package main

import "fmt"

// this func takes two int and returns an int
// go requires explicit return definition
func plus(a int, b int) int {
	return a + b
}

// type can be omitted for like-typed parameters
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2=", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3=", res)

}
