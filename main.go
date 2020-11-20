package main

import (
	"fmt"
	"os"
	"strconv"

	"go-pythagoras/pythagoras"
)

func main() {

	if len(os.Args) == 3 {
		a, _ := strconv.ParseFloat(os.Args[1], 0)
		b, _ := strconv.ParseFloat(os.Args[2], 0)
		hypotenuse := pythagoras.GetHypotenuse(a, b)
		area := pythagoras.GetArea(a, b)
		perimeter := pythagoras.GetPerimeter(a, b)
		fmt.Printf("Results for triangle [%f,%f]:\n", a, b)
		fmt.Printf("Hypotenuse: %f\n", hypotenuse)
		fmt.Printf("Area: %f\n", area)
		fmt.Printf("Perimeter: %f\n", perimeter)
	} else {
		fmt.Println("Must be called with 2 arguments.")
	}

}
