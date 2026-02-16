package main

import "fmt"

func main() {
	var number int
	var largest int
	var secondLargest int

	fmt.Println("Enter a 10 numbers")

	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	largest = number
	secondLargest = number

	for count := 1; count < 10; count++ {
		fmt.Print("Enter a number: ")
		fmt.Scan(&number)

		if number > largest {
			secondLargest = largest
			largest = number
		} else if number > secondLargest {
			secondLargest = number
		}
	}

	fmt.Println("The largest number is: ", largest)
	fmt.Println("The second largest number is: ", secondLargest)


	
}