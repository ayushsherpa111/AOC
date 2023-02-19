package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const STAGE_SIZE = 9

func parse_stage(line string) []string {
	line_parsed := make([]string, STAGE_SIZE)
	for i, j := 1, 0; i < len(line); i += 4 {
		line_parsed[j] = string(line[i])
		j++
	}
	return line_parsed
}

func parse_moves(line string) (int, int, int) {
	list := strings.Split(line, " ")
	quantity, _ := strconv.Atoi(list[1])
	source, _ := strconv.Atoi(list[3])
	target, _ := strconv.Atoi(list[5])
	return quantity, source - 1, target - 1
}

func main() {
	input_file, err := os.Open("./input.txt")
	var stage [][]string
	var stage_received = false
	var line string
    var temp []string

	stage = make([][]string, STAGE_SIZE)

	if err != nil {
		fmt.Println("Failed to open input file." + err.Error())
		return
	}
	scanner := bufio.NewScanner(input_file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line = scanner.Text()

		if len(line) == 0 {
			stage_received = true
			continue
		}

		if !stage_received {
			for i, column := range parse_stage(line) {
				if _, err := strconv.Atoi(column); err == nil || column == " " {
					continue
				}
				stage[i] = append(stage[i], column)
			}
		} else {
			quantity, source, target := parse_moves(line)
			temp, stage[source] = pop_stack(stage[source], quantity)
			stage[target] = append(temp, stage[target]...)
		}
	}
	for _, rows := range stage {
		fmt.Println(rows)
	}
}

func pop_stack(stack []string, quantity int) ([]string, []string) {
    temp := make([]string,0, quantity)
    for i := range stack[:quantity] {
        temp = append([]string{stack[i]}, temp...) 
    }
    return temp, stack[quantity:]
}
