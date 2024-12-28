package main

import (
	"bytes"
	"regexp"
)

func solvepart2(inputLines []string, signalChan chan int) {
	for i, v := range inputLines {
		prep := make([]string, 0, 3)
		if i == 0 || i == len(inputLines)-1 {
			continue
		}
		// take line above, current and below
		prep = append(prep, inputLines[i-1], v, inputLines[i+1])
		go func() {
			result := findMAS(prep)
			signalChan <- result
		}()
	}
}

var (
	reMAS = regexp.MustCompile("(?i)(mas)")
	reSAM = regexp.MustCompile("(?i)(sam)")
)

func checkMAS(searchB []byte) bool {
	return reMAS.Match(searchB) || reSAM.Match(searchB)
}

func findMAS(input []string) (sum int) {
	for i := 1; i < len(input[1])-1; i++ {
		if bytes.Equal([]byte{input[1][i]}, []byte("A")) {
			left := []byte{input[0][i-1], input[1][i], input[2][i+1]}
			right := []byte{input[0][i+1], input[1][i], input[2][i-1]}
			if checkMAS(left) && checkMAS(right) {
				sum += 1
			}
		}
	}
	return
}
