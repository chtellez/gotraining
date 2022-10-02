// --Summary:
//
//	Create a program to calculate the area and perimeter
//	of a rectangle.
package main

import (
	"fmt"
)

// --Requirements:
// * Create a rectangle structure containing its coordinates
type rectangle struct {
	Height int
	Width  int
}

// * Using functions, calculate the area and perimeter of a rectangle,
//   - The functions must use the rectangle structure as the function parameter
func area(r rectangle) int {
	//* The area of a rectangle is length*width
	return r.Height * r.Width
}
func perimeter(r rectangle) int {
	//* The perimeter of a rectangle is the sum of the lengths of all sides
	return (r.Width * 2) + (r.Height * 2)
}

func main() {
	rec := rectangle{10, 20}
	//  - Print the results to the terminal
	fmt.Println("El área del rectángulo es:", area(rec))
	fmt.Println("El perímetro del rectángulo es:", perimeter(rec))
	//* After performing the above requirements, double the size
	//  of the existing rectangle and repeat the calculations
	rec.Height *= 2
	rec.Width *= 2
	//  - Print the new results to the terminal
	fmt.Println("El área del rectángulo es:", area(rec))
	fmt.Println("El perímetro del rectángulo es:", perimeter(rec))

}
