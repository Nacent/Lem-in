package main

import (
	"bufio"
	"fmt"
	"os"
)


func readFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . test.txt")
		os.Exit(1)
	}

	lines, err := readFile(os.Args[1])
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}