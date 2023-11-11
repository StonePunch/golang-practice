package main

import (
	"fmt"
	"io"
	"os"
)

/*

Create a program that reads the contents of a text file then prints its contents
to the terminal.
The file to open should be provided as an argument to the program when it is
executed at the terminal. For example, 'go run main.go myfile.txt" should open
up the myfile.txt file
To read in the arguments provided to a program, you can reference the
variable 'os.Args', which is a slice of type string
To open a file, check out the documentation for the 'Open' function in the 'os'
package - https://golang.org/pkg/os/#Open
What interfaces does the 'File' type implement?
If the 'File' type implements the 'Reader interface, you might be able to reuse
that io.Copy function!

*/

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println("File contents:", string(bs))

	return len(bs), nil
}

func readFile(file string) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(logWriter{}, f)
}

func main() {
	if args := os.Args; len(args) == 2 {
		readFile(args[1])
	} else {
		fmt.Println("No file provided")
	}
}
