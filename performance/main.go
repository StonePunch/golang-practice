package main

import (
	"fmt"
	"time"
)

// Assumptions:
// - totalWork is bigger or equal to 20 and always a even number
// - the order of the final result does not matter
// Logic could be added to account for these situations, but it was outside
// the scope of what I'm trying to do.
const totalWork = 10000000

// Run benchmark with: go test -bench=. -benchtime=5s
func main() {
	var start time.Time
	var slice []int

	start = time.Now()
	slice = getSliceBase(totalWork)
	fmt.Printf("getSliceBase execution time: %s | Length of slice: %d\n", time.Since(start), len(slice))

	start = time.Now()
	slice = getSliceWorker(totalWork)
	fmt.Printf("getSliceWorker execution time: %s | Length of slice: %d\n", time.Since(start), len(slice))

	start = time.Now()
	slice = getSliceSimple(totalWork)
	fmt.Printf("getSliceSimple execution time: %s | Length of slice: %d\n", time.Since(start), len(slice))
}
