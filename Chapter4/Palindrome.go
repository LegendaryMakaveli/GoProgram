package main


func palindrome(number int) bool {

	reversedNumber := 0
	originalNumber := number

	for number != 0 {
		remainder := number % 10
		reversedNumber = reversedNumber*10 + remainder
		number /= 10
	}

	return originalNumber == reversedNumber
}

