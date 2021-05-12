// Created by: Infinity de Guzman
// Created on: May 2021
//
// This program calculates the area of a trapezoid

package main

import "fmt"

func main() {
	// This function prints the address
	var aBase int
	var bBase int
  var height int

	// input
	fmt.Println("This program calculates the area of a trapezoid.")
	fmt.Println()
	fmt.Print("Enter the a base: ")
	fmt.Scanln(&aBase)
	fmt.Print("Enter the b base: ")
	fmt.Scanln(&bBase)
	fmt.Print("Enter the height: ")
	fmt.Scanln(&height)

	// output
	fmt.Println("The area is: ", ((aBase + bBase) / 2) * height, "cm²")
	fmt.Println("\nDone.")
}
