package main

import "fmt"

func validateUserInput(number int) {

	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	for number != 1 && number != 2 {
		fmt.Println("Invalid number")
		fmt.Print("Enter a number: ")
		fmt.Scan(&number)
	}

	fmt.Println("Valid number")
}
