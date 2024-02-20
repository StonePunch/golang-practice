package main

import (
	"sync"
)

const (
	workerNum  = 10
	jobDivider = 20
)

// Job represents a single to job be done by a worker
type job struct {
	offsetNum  int
	workSubset int
	work       func(offsetNum, workSubset int) []int
}

func getSliceWorker(totalWork int) []int {
	out := make(chan []int, jobDivider)
	jobs := make(chan job, jobDivider)

	// Start up workers
	var wgWorkers sync.WaitGroup
	for i := 0; i < workerNum; i++ {
		wgWorkers.Add(1)
		go func() {
			defer wgWorkers.Done()
			worker(jobs, out)
		}()
	}

	// Add work to be done
	go func() {
		for i := 0; i < jobDivider; i++ {
			jobs <- job{
				offsetNum:  i,
				workSubset: totalWork / jobDivider,
				work:       addToSlice,
			}
		}

		close(jobs)
	}()

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

func worker(jobs <-chan job, out chan<- []int) {
	for job := range jobs {
		result := job.work(job.offsetNum, job.workSubset)
		out <- result
	}
}

func addToSlice(offsetNum, workSubset int) []int {
	offsetNum = offsetNum * workSubset
	data := make([]int, workSubset)

	for i := 0; i < workSubset; i++ {
		data[i] = offsetNum + i
	}

	return data
}
