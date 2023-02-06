package main

import (
	"fmt"
)

func main() {

	var number int

	fmt.Println("Input number: ")
	fmt.Scan(&number)

	sum, a := 0, number

	for a > 0 {

		remain := a % 10

		sum = sum*10 + remain

		a /= 10
	}

	if sum == number {
		fmt.Println("Palindrom number")
		return
	}

	fmt.Println("Not Palindrome")
}
