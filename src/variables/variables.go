// Go compiler will throw error if variable is declared and not used.
package main

import "fmt"

// declare a global variable. This is lower case so only accessible in this package
var j int

// var block to define multiple variables at package level
var (
	w int     = 12
	y float32 = 12.0
	x bool    = true
)

func main() {
	fmt.Println(w, x, y)

	j = 12
	fmt.Println(j)

	// declare variable. Variables are initialized to zero in go
	var i int
	i = 12
	fmt.Println(i)

	// declare and initialize variable. Auto-detect type
	var a = "initial"
	fmt.Println(a)

	// declare and define multiple variables
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	// Short notation for declare and define(type auto-detect)
	f := "apple"
	// print value(%v) and type(%T) of variable
	// notice that this needs printf
	fmt.Printf("%v, %T\n", f, f)

	// Converting types
	var g float32 = 12.4
	var k int
	k = int(g)
	fmt.Println(k)
}
