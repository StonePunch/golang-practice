package main

import (
	"testing"
)

var getSliceBaseResult []int

func Benchmark_GetSliceBase(b *testing.B) {
	var r []int

	for i := 0; i < b.N; i++ {
		r, _ = getSliceBase(totalWork)
	}

	getSliceBaseResult = r
}
