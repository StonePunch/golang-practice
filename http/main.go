package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

// Func that implements the io.Writer interface
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))

	return len(bs), nil
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(logWriter{}, resp.Body)
}
