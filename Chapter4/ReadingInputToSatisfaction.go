package main

import "fmt"


func readingInputToSatisfaction(target int) {
	var sum int = 0
	var input int

	for sum < target {
		fmt.Print("Enter an integer: ")
		fmt.Scan(&input)

		sum += input
		fmt.Println("Current sum:", sum)
	}

	fmt.Println("Final sum:", sum)
	fmt.Println("Target reached or exceeded!")
}
