package main

import (
	"fmt"
	"time"
)

func getSliceBase(totalWork int) ([]int, string) {
	start := time.Now()

	data := []int{}
	for i := 0; i < totalWork; i++ {
		data = append(data, i)
	}

	return data, fmt.Sprint("The getSliceBase execution time was: ", time.Since(start))
}
