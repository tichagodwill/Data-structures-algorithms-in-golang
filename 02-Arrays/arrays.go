package main

import "fmt"

func main() {
	//defining an array

	// method 1
	var numbers = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(numbers)

	//method 2
	numbersList := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(numbersList)

	//looping over an array in golang
	// method 1: using for loop

	for i := 0; i <= len(numbersList); i++ {
		fmt.Println(numbersList[i])
	}

	//Let's print the numbers using a loop
	for n := 0; n <= 10; n++ {
		fmt.Println(n)
	}

}
