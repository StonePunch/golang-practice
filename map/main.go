package main

import "fmt"

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}

func main() {
	// Empty map creation
	colors := make(map[string]string)
	fmt.Println(colors)

	// Assignment value
	colors["red"] = "#ff0000"
	colors["green"] = "#4bf745"
	colors["white"] = "#ffffff"
	fmt.Println(colors)

	// Retrieve value
	a := colors["red"]
	fmt.Println("red:", a)

	// Retrieve value with non existent key
	b := colors["black"]
	fmt.Println("black:", b)

	// Check if the key exists
	_, c := colors["black"]
	fmt.Println("Key exists:", c)

	// Delete single from map
	delete(colors, "green")
	fmt.Println(colors)

	// Clear all entries from map
	clear(colors)
	fmt.Println(colors)

	// Another way to initialize a map
	colors = map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)
}
