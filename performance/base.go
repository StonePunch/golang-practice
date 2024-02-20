package main

func getSliceBase(totalWork int) []int {
	data := []int{}
	for i := 0; i < totalWork; i++ {
		data = append(data, i)
	}

	return data
}
