package main

import (
	"bufio"
	"fmt"
	"os"
)

func prepareInput(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	lines := prepareInput(file)
	sum := 0
	signal := make(chan int)
    defer close(signal)
	go solvepart2(lines, signal)

	for i := 0; i < len(lines)-2; i++ {
		sum += <-signal
	}
	fmt.Println(sum)
}
