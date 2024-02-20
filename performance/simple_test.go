package main

import (
	"testing"
)

var getSliceSimpleResult []int

func Benchmark_GetSliceSimple(b *testing.B) {
	var r []int

	for i := 0; i < b.N; i++ {
		r = getSliceSimple(totalWork)
	}

	getSliceSimpleResult = r
}
