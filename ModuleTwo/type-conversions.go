package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x*x + y*y))
	var z uint = uint(f)
	a := uint(f)
	fmt.Printf("%v %v %v %v %v\n", x, y, f, z, a)
}
