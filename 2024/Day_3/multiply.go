package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

func ParseFromList(rePattern *regexp.Regexp, mulInst string) int {
	sum := 0
	for _, match := range rePattern.FindAllStringSubmatch(string(mulInst), -1) {
		if len(match) < 4 {
			continue
		}
		xStr, yStr := match[2], match[3]
		x, _ := strconv.Atoi(xStr)
		y, _ := strconv.Atoi(yStr)
		sum += x * y
	}
	return sum
}

func main() {
	isDo := true
	sum := 0
	doInst := []byte("do()")

	inputFile, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("Failed to open file. %s", err.Error())
	}
    defer inputFile.Close()

	inputRaw, err := io.ReadAll(inputFile)
	if err != nil {
		log.Fatalf("Error reading file. %s", err.Error())
	}
    defer inputFile.Close()


	multiplyReg := regexp.MustCompile(`(mul\(([0-9]{1,3}),([0-9]{1,3})\))`)
	doDont := regexp.MustCompile(`(do\(\))|(don't\(\))`)

	startIdx := 0
	searchSpace := doDont.FindAllStringIndex(string(inputRaw), -1)
	for i, match := range searchSpace {
		isDo = bytes.Equal(inputRaw[match[0]:match[1]], doInst)
		if !isDo {
			if startIdx != -1 {
				sum += ParseFromList(multiplyReg, string(inputRaw[startIdx:match[0]]))
			}
			startIdx = -1
		}

		if isDo && startIdx == -1 {
			startIdx = match[1]
		}

		if i == len(searchSpace)-1 {
			sum += ParseFromList(multiplyReg, string(inputRaw[startIdx:]))
		}
	}

	fmt.Println(sum)
}
