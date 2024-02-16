package main

import (
	"fmt"
	"sync"
	"time"
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

func getSliceWorker(totalWork int) ([]int, string) {
	start := time.Now()

	out := make(chan []int, jobDivider)
	jobs := make(chan job, jobDivider)

	// Start up workers
	var wgWorkers sync.WaitGroup
	for i := 1; i <= workerNum; i++ {
		wgWorkers.Add(1)
		go func(id int) {
			defer wgWorkers.Done()
			worker(id, jobs, out)
		}(i)
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

	return data, fmt.Sprint("The getSliceWorker execution time was: ", time.Since(start))
}

func worker(id int, jobs <-chan job, out chan<- []int) {
	fmt.Println("worker", id, "started")

	for job := range jobs {
		fmt.Println("worker", id, "started job")

		result := job.work(job.offsetNum, job.workSubset)
		out <- result

		fmt.Printf("worker %d ended job, values between %d and %d\n",
			id,
			result[0],
			result[len(result)-1],
		)
	}

	fmt.Println("worker", id, "stopped")
}

func addToSlice(offsetNum, workSubset int) []int {
	offsetNum = offsetNum * workSubset
	data := make([]int, workSubset)

	for i := 0; i < workSubset; i++ {
		data[i] = offsetNum + i
	}

	return data
}
