package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type config struct {
	timeout      time.Duration
	workerNumber int
}

func main() {
	ctx := context.Background()
	dartNumber := 5

	fmt.Println("Methods: ")
	fmt.Printf("Select option: ")
	var selection string
	fmt.Scanln(&selection)

	// Measure execution time
	start := time.Now()
	defer func() {
		fmt.Println("The execution time was:", time.Since(start))
	}()

	app := config{
		timeout:      time.Millisecond * 3000,
		workerNumber: 5,
	}

	var scores *[]int
	var err error

	switch selection {

	default:
		fmt.Println("Error: Unrecognized code")
		return
	}

	total := 0
	for _, v := range *scores {
		total += v
	}

	fmt.Println("\nThe individual scores are: ", *scores)
	fmt.Printf("The total score is: %d/%d\n", total, dartNumber*10)
}

func throwDart() int {
	time.Sleep(time.Millisecond * 500)

	rand.NewSource(time.Now().UnixNano())
	throwScore := rand.Intn(10)

	return throwScore
}
