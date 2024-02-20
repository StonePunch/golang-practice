package main

func getSliceSimple(totalWork int) []int {
	data := make([]int, totalWork)

	for i := 0; i < totalWork; i++ {
		data[i] = i
	}

	return data
}
