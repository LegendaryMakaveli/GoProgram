package main

import "fmt"

func rightAngle(number int) {
	fmt.Println("Enter a number")
	fmt.Scan(&number)

	if number < 1 || number > 10 {
		fmt.Println("Number must between 1 to 10")
		return
	}

	for count := 1; count <= number; count++ {
		for innerCount := 1; innerCount <= count; innerCount++ {
			fmt.Print("*")
		}
		fmt.Println(" ")
	}
}
