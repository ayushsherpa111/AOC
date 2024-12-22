package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err.Error())
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	lines := make([]string, 0)
	sum := 0

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := 0; i < len(lines); i++ {
		prep := make([]string, 0, 3)
		currLine := 0
		if i >= 3 {
			prep = append(prep, lines[i-3], lines[i-2], lines[i-1])
			currLine += 3
		}
		prep = append(prep, lines[i])
		sum += findXMAS(prep, currLine)
	}
	fmt.Println(sum)
}

func checkXMAS(searchB []byte) int {
	if reXMAS.Match(searchB) || reSAMX.Match(searchB) {
		return 1
	}
	return 0
}

func findXMAS(inpXmas []string, currentLine int) int {
	xmasCount := 0
	xmasCount += len(reSAMX.FindAllIndex([]byte(inpXmas[currentLine]), -1))
	xmasCount += len(reXMAS.FindAllIndex([]byte(inpXmas[currentLine]), -1))

	if currentLine >= 3 {
		// look north
		for col := 0; col < len(inpXmas[0]); col++ {
			if inpXmas[currentLine][col] != byte('X') && inpXmas[currentLine][col] != byte('S') {
				continue
			}

			var isXMAS []byte
			isXMAS = []byte{
				inpXmas[currentLine][col],
				inpXmas[currentLine-1][col],
				inpXmas[currentLine-2][col],
				inpXmas[currentLine-3][col],
			}
			xmasCount += checkXMAS(isXMAS)

			if col < len(inpXmas[0])-3 {
				// look north east
				isXMAS = []byte{
					inpXmas[currentLine][col],
					inpXmas[currentLine-1][col+1],
					inpXmas[currentLine-2][col+2],
					inpXmas[currentLine-3][col+3],
				}
				xmasCount += checkXMAS(isXMAS)
			}
			if col >= 3 {
				// look north west
				isXMAS = []byte{
					inpXmas[currentLine][col],
					inpXmas[currentLine-1][col-1],
					inpXmas[currentLine-2][col-2],
					inpXmas[currentLine-3][col-3],
				}
				xmasCount += checkXMAS(isXMAS)
			}
		}
	}

	return xmasCount
}
