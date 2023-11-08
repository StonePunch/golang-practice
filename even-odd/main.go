package main

import "fmt"

func main() {
	input := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	for _, value := range input {
		if value%2 == 0 {
			fmt.Printf("%d is even \n", value)
		} else {
			fmt.Printf("%d is odd \n", value)
		}
	}
}
