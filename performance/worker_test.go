package main

import (
	"testing"
)

var getSliceWorkerResult []int

func Benchmark_GetSliceWorker(b *testing.B) {
	var r []int

	for i := 0; i < b.N; i++ {
		// always record the result of getSliceWorker to prevent the compiler
		// eliminating the function call.
		r, _ = getSliceWorker(totalWork)
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	getSliceWorkerResult = r
}
