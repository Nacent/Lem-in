package main

import (
	"fmt"
	"os"
)

func readFile(filename string) {
	file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("test")
		os.Exit(1)
	}

	ant, err := readFile(os.Args[1])
	if err != nil {
		fmt.Println("error")
		return
	}
}