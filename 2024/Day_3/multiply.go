package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("Failed to open file. %s", err.Error())
	}
	inputRaw, err := io.ReadAll(inputFile)
	if err != nil {
		log.Fatalf("Error reading file. %s", err.Error())
	}

	multiplyReg := regexp.MustCompile(`(mul\(([0-9]{1,3}),([0-9]{1,3})\))`)
	sum := 0
	for _, match := range multiplyReg.FindAllStringSubmatch(string(inputRaw), -1) {
		if len(match) == 4 {
			xStr, yStr := match[2], match[3]
			x, err := strconv.Atoi(xStr)
			if err != nil {
				panic(fmt.Sprintf("error parsing %d\n", x))
			}
			y, err := strconv.Atoi(yStr)
			if err != nil {
				panic(fmt.Sprintf("error parsing %d\n", y))
			}
			sum += x * y
		}
	}
	fmt.Println(sum)
}
