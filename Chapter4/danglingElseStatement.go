package main

import "fmt"

func danglingElseStatement() {
	fmt.Println("Enter first number")
	var number int
	fmt.Scan(&number)

	fmt.Println("Enter the second number")
	var secondNumber int
	fmt.Scan(&secondNumber)

	if number > 5 {
		if secondNumber > 5 {
			fmt.Println("first number and second number is greater than 5")
		} else {
			fmt.Println("Only first number is greater than 5")
		}
	} else {
		fmt.Println("none is greater than 5")
	}
}
