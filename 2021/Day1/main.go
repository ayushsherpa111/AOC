package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, e := os.Open("./input.txt")
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	input_scanner := bufio.NewScanner(file)
	input_scanner.Split(bufio.ScanLines)
	depth := make([]int, 0)
	var increment int
	for input_scanner.Scan() {
		dpt, _ := strconv.Atoi(input_scanner.Text())
		depth = append(depth, dpt)
	}
	var curr_dept int
	var next_dept int
	for i := 0; i < len(depth)-3; i++ {
		comm_depth := depth[i+1] + depth[i+2]
		curr_dept = depth[i] + comm_depth
		next_dept = comm_depth + depth[i+3]
		if next_dept > curr_dept {
			increment += 1
		}
	}
	fmt.Println(increment)
}
