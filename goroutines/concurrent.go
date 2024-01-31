package main

import (
	"context"
	"fmt"
)

func (c config) generateScoreConcurrent(ctx context.Context, dartNumber int) (*[]int, error) {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	out := make(chan int)
	for i := 0; i < dartNumber; i++ {
		go func() {
			out <- throwDart()
		}()
	}

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
