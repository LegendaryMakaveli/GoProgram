package main


func comparator(number int, secondNumber int) int {
	if number == secondNumber {
		return 0
	} else if number > secondNumber {
		return 1
	}
	return -1
}
