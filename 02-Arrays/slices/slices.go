package main

import "fmt"

func main() {
	myArray := [5]int{1, 2, 3, 4, 5}
	mySlice := myArray[1:3]
	fmt.Println(mySlice) // prints [2 3]

	/// Appending method to a slice

	mySlice1 := []int{1, 2, 3, 4, 5}
	myNewSlice := append(mySlice1[:2], 6, 7)
	fmt.Println(myNewSlice) // prints [1 2 6 7]

	// creating a slice from literals
	sliceOfIntegers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// accessing slice values with their indexes
	firstInteger := sliceOfIntegers[0]  // 1
	secondInteger := sliceOfIntegers[1] // 2
	lastInteger := sliceOfIntegers[9]   // 10
	fmt.Println(firstInteger)
	fmt.Println(secondInteger)
	fmt.Println(lastInteger)

	// using a for loop to access a slice
	for i := 0; i < 10; i++ {
		fmt.Println(sliceOfIntegers[i])
	}

	// using range to access a slice
	for index, value := range sliceOfIntegers {
		fmt.Println(index, value)
	}

}
