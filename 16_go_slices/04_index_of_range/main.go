package main

import "fmt"

func main() {

	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "buenos dias!"

	// greeting[3] = "suprabadham" //Gives index Out of range.
	// panic: runtime error: index out of range
	// Length is limited to 3. You've to use append
	greeting = append(greeting, "suprabadham")

	fmt.Println(greeting[2])
}
