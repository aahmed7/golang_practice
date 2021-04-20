package main

import "fmt"

func main() {
	// Unlike arrays, slices are typed only by the elements they
	// contain (not the number of elements). To create an empty slice
	// with non-zero length, use the builtin make. Here we make a slice
	// of strings of length 3 (initially zero-valued).
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	// add one element
	s = append(s, "d")

	// add multiple elements
	s = append(s, "e", "f")

	fmt.Println("apd:", s)

	// slices can also be copied. create a new slice
	// and copy the older one into it
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// slice operation: slice[low:high]
	// this will copy elements between low and high
	// including low
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	// delcare and initialize is supported as well
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// multidimensional slice
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
