package main

import (
	"fmt"
)

func main() {
	// Assumptions:
	// - totalWork is bigger or equal to 20 and always a even number
	// - the order of the final result does not matter
	// Logic could be added to account for these situations, but it was outside
	// the scope of what I'm trying to do.
	totalWork := 10000000

	// execution time using base code: 63ms
	slice, executionTime := getSliceBase(totalWork)
	fmt.Printf("%s | Length of slice: %d\n", executionTime, len(slice))

	// execution time using workers code: 25ms
	slice, executionTime = getSliceWorker(totalWork)
	fmt.Printf("%s | Length of slice: %d\n", executionTime, len(slice))
}
