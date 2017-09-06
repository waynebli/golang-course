package main

import (
	"fmt"
	"io"
	"os"
)

type fileWriter struct{}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please supply a filename. Usage: ./main 'filename'")
		os.Exit(1)
	}
	filename := args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fw := fileWriter{}

	io.Copy(fw, file)

}

func (fileWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
