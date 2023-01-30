package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.OpenFile("./input.txt", os.O_RDONLY, 0o777)
	if err != nil {
		fmt.Println(err.Error())
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	max := make([]int, 3)
	var sum int
	for scanner.Scan() {
		i := scanner.Text()
		if len(i) != 0 {
			val, _ := strconv.Atoi(i)
			sum += val
		} else {
			if sum > max[0] {
				max[0], max[1], max[2] = sum, max[0], max[1]
			} else if sum > max[1] {
				max[1], max[2] = sum, max[1]
			} else if sum > max[2] {
				max[2] = sum
			}
			sum = 0
		}
	}
	fmt.Println(max[0] + max[1] + max[2])
}
