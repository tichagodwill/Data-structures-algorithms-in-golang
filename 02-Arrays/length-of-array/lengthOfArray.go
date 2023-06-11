package main

import "fmt"

func main() {
	// lenght of an array
	myArray := [7]string{"John", "peter", "Ali"}
	fmt.Println(len(myArray)) //output 7

	// To get the number of elements in the array, we iterate over it.. see code below

	arrayElements := [10]string{"paul", "peter", "trump", "Mohammed"}

	count := 0

	for i := 0; i < len(arrayElements); i++ {
		if arrayElements[i] != "" {
			count++
		}
	}
	fmt.Println("Number of elements in the array:", count)
	//Number of elements in the array: 4

}
