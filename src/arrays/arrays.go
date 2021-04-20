package main

import "fmt"

func main() {
	// Declare an array. All elements initialize to zero
	var a [5]int
	fmt.Println("emp:", a)

	// Set a single element
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// get length of array
	fmt.Println("len:", len(a))

	// Declare and initialize an array
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	// Multidimensional array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
