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
	go func() {
		for i := 0; i < len(lines); i++ {
			prep := make([]string, 0, 3)
			currLine := 0
			if i >= 3 {
				prep = append(prep, lines[i-3], lines[i-2], lines[i-1])
				currLine += 3
			}
			prep = append(prep, lines[i])

			go func() {
				result := FindXMAS(prep, currLine)
				signal <- result
			}()
		}
	}()

	for i := 0; i < len(lines); i++ {
		sum += <-signal
	}
	fmt.Println(sum)
}
