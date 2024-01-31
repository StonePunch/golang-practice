package main

import (
	"context"
	"fmt"
)

func (c config) generateScoreSynchronous(ctx context.Context, dartNumber int) (*[]int, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	out := make(chan int)
	go func() {
		for i := 0; i < dartNumber; i++ {
			out <- throwDart()
		}
		close(out)
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
		}
	}
}
