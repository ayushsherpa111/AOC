package main

import (
	"regexp"
)

/*
- For line n read n-3, n, n+3 lines.
- If n-3 < 0 skip reading lines above
- If n+3 < numOfLines skip reading lines below.

- Store read lines -> [n-3,n-2,n-1,n,n+1,n+2,n+3] and find the words XMAS vertically, horizontally and diagnally.
*/

var (
	reXMAS = regexp.MustCompile("(?i)(xmas)")
	reSAMX = regexp.MustCompile("(?i)(samx)")
)

func checkXMAS(searchB []byte) int {
	if reXMAS.Match(searchB) || reSAMX.Match(searchB) {
		return 1
	}
	return 0
}

func solvePart1(lines []string, signal chan int) {
	for i := 0; i < len(lines); i++ {
		prep := make([]string, 0, 3)
		currLine := 0
		if i >= 3 {
			prep = append(prep, lines[i-3], lines[i-2], lines[i-1])
			currLine += 3
		}
		prep = append(prep, lines[i])

		go func() {
			result := findXMAS(prep, currLine)
			signal <- result
		}()
	}
}

func findXMAS(inpXmas []string, currentLine int) int {
	mainChan := make(chan int)
	defer close(mainChan)

	sumval := 0

	go func(mChan chan int) {
		count := 0
		count += len(reSAMX.FindAllIndex([]byte(inpXmas[currentLine]), -1))
		count += len(reXMAS.FindAllIndex([]byte(inpXmas[currentLine]), -1))
		mChan <- count
	}(mainChan)

	// look north
	go func(ch chan int, input []string, currLine int) {
		northCount := 0
		if currLine == 3 {
			// look north
			var isXMAS []byte
			for col := 0; col < len(input[0]); col++ {
				if input[currentLine][col] != byte('X') &&
					input[currentLine][col] != byte('S') {
					continue
				}
				isXMAS = []byte{
					input[currentLine][col],
					input[currentLine-1][col],
					input[currentLine-2][col],
					input[currentLine-3][col],
				}
				northCount += checkXMAS(isXMAS)
			}
		}
		ch <- northCount
		// close(ch)
	}(mainChan, inpXmas, currentLine)

	go func(ch chan int, input []string, currLine int) {
		northEastCount := 0
		if currLine == 3 {
			// look north
			var isXMAS []byte
			for col := 0; col < len(input[0]); col++ {
				if input[currLine][col] != byte('X') &&
					input[currLine][col] != byte('S') {
					continue
				}
				if col < len(input[0])-3 {
					// look north east
					isXMAS = []byte{
						input[currLine][col],
						input[currLine-1][col+1],
						input[currLine-2][col+2],
						input[currLine-3][col+3],
					}
					northEastCount += checkXMAS(isXMAS)
				}
			}
		}
		ch <- northEastCount
		// close(ch)
	}(mainChan, inpXmas, currentLine)

	go func(ch chan int, input []string, crrLine int) {
		northWestCount := 0
		// look north
		if crrLine == 3 {
			var isXMAS []byte
			for col := 0; col < len(input[0]); col++ {
				if input[crrLine][col] != byte('X') &&
					input[crrLine][col] != byte('S') {
					continue
				}
				if col >= 3 {
					// look north west
					isXMAS = []byte{
						input[crrLine][col],
						input[crrLine-1][col-1],
						input[crrLine-2][col-2],
						input[crrLine-3][col-3],
					}
					northWestCount += checkXMAS(isXMAS)
				}
			}
		}
		ch <- northWestCount
	}(mainChan, inpXmas, currentLine)

	for i := 0; i < 4; i++ {
		sumval += <-mainChan
	}

	return sumval
}
