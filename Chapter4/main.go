package main

import "fmt"

func main() {
	// fmt.Println("=== Validate User Input ===")
	// validateUserInput(3)

	// fmt.Println("\n=== Right Angle ===")
	// rightAngle(5)

	// fmt.Println("\n=== Largest and Second Largest ===")
	// largestAndSecondLargest(4, 7)

	// fmt.Println("\n=== Dangling Else Statement ===")
	// danglingElseStatement()

	// fmt.Println("\n=== Palindrome ===")
	// number := 1232
	// fmt.Println(palindrome(number))

	// fmt.Println("\n=== Binary to Decimal ===")
	// binary := 1101
	// fmt.Println(binaryToDecimal(binary))

	// fmt.Println("\n === Collect input till user input add up to target number ===")
	// var targetNumber int
	// fmt.Print("Enter a target number: ")
	// fmt.Scan(&targetNumber)

	// readingInputToSatisfaction(targetNumber)

	fmt.Println("\n === Comperator ===")
	var firNstNumber int
	var secondNumber int
	
	fmt.Print("Enter first number: ")
	fmt.Scan(&firNstNumber)
	fmt.Print("Enter second number: ")
	fmt.Scan(&secondNumber)
	
	fmt.Println(comparator(firNstNumber, secondNumber))
}