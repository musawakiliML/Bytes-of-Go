// Weight Loss Program

package main

import "fmt"

func main() {
	fmt.Print("My weight on the surface of Mars is ")
	fmt.Print(149.0 * 0.3783)
	fmt.Print(" lbs, and I would be ")
	fmt.Print(41 * 365 / 687)
	fmt.Print(" years old.")

	fmt.Printf("My weight on the surface of Mars is %v\n", 149.0*0.3783)
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
}