package main

type rect struct {
	width, height int
}

// pointer based method
func (r *rect) area() int {
	return r.width * r.height
}

// value based method
func (r rect) perim() int {
	return 2*r.height + 2*r.width
}

func main() {
	r := rect{width: 10, height: 5}

	// calling the methods
	println("Area:", r.area())
	println("Perimeter:", r.perim())

	// Go automatically handles conversion between
	// value and pointer for methods
	rp := &r
	println("Area:", rp.area())
	println("Perimeter:", rp.perim())

}
