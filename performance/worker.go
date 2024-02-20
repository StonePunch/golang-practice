package main

import (
	"sync"
)

const workerNum = 10

// Job represents a single to job be done by a worker
type job struct {
	offsetNum  int
	workSubset int
	work       func(offsetNum, workSubset int) []int
}

func getSliceWorker(totalWork int) []int {
	out := make(chan []int, workerNum)

	// Start up workers
	var wgWorkers sync.WaitGroup
	for i := 0; i < workerNum; i++ {
		wgWorkers.Add(1)
		go func(offsetNum int) {
			defer wgWorkers.Done()

			j := job{
				offsetNum:  offsetNum,
				workSubset: totalWork / workerNum,
				work:       addToSlice,
			}

			worker(j, out)
		}(i)
	}

	// Close the out channel when all workers stop
	go func() {
		wgWorkers.Wait()
		close(out)
	}()

	// Saved 20ms by doing this instead of "data := []int{}", the overhead of
	// expanding the underlying array is considerable.
	data := make([]int, 0, totalWork)
	for v := range out {
		data = append(data, v...)
	}

	return data
}

func worker(j job, out chan<- []int) {
	result := j.work(j.offsetNum, j.workSubset)
	out <- result
}

func addToSlice(offsetNum, workSubset int) []int {
	offsetNum = offsetNum * workSubset
	data := make([]int, workSubset)

	for i := 0; i < workSubset; i++ {
		data[i] = offsetNum + i
	}

	return data
}
