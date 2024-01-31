package main

import (
	"context"
	"fmt"
)

func worker(id int, jobs <-chan func() int, out chan<- int) {
	for job := range jobs {
		fmt.Println("worker", id, "started job")
		out <- job()
		fmt.Println("worker", id, "ended job")
	}
	fmt.Println("worker", id, "closed")
}

func (c config) generateScoreWorkerPool(ctx context.Context, dartNumber int) (*[]int, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	out := make(chan int, dartNumber)
	jobs := make(chan func() int, dartNumber)

	// Only spawn a limited number of goroutines to do the work
	// In the context of a docker container, it keeps the memory utilization under control
	for i := 1; i <= c.workerNumber; i++ {
		go worker(i, jobs, out)
	}

	// Add work to be done
	go func() {
		for i := 0; i < dartNumber; i++ {
			jobs <- throwDart
		}
		close(jobs)
	}()

	scores := []int{}

	for {
		select {
		case <-ctx.Done():
			return &scores, fmt.Errorf("Context canceled")

		case score, open := <-out:
			if !open {
				return &scores, nil
			}

			scores = append(scores, score)

			if len(scores) == dartNumber {
				close(out)
			}
		}
	}
}
