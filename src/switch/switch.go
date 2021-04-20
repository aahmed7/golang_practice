package main

import (
	"fmt"
	"time"
)

func main() {
	// Simple switch statement
	i := 2
	fmt.Println("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// Use commas to separate multiple expressions
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its a weekend")
	default:
		fmt.Println("Tis a weekday")
	}

	// switch without expression is an alternate way to use if/else
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// Type switch compares types instead of values
	WhatamI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	WhatamI(true)
	WhatamI(1)
	WhatamI("hey")
}
