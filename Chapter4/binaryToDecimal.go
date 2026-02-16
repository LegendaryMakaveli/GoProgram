package main

func binaryToDecimal(binary int) int {

	decimal := 0
	base := 1 

	for binary != 0 {

		remainder := binary % 10  
		decimal += remainder * base
		base *= 2   
		binary /= 10  
	}

	return decimal
}

