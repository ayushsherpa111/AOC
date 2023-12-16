package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	line_scanner := bufio.NewScanner(file)
	var sum int

	for line_scanner.Scan() {
		inp := line_scanner.Text()
		_, game_set, err := parse_input(inp)
		if err != nil {
			fmt.Printf("Found error: %s while parsing input: %s\n", err.Error(), inp)
			break
		}
        sum += minimum_set(game_set)
	}
	fmt.Println(sum)
}

type game struct {
	color string
	count int
}

var rule map[string]int = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

// Game 1: 9 red, 5 blue, 6 green; 6 red, 13 blue; 2 blue, 7 green, 5 red
func parse_input(game_input string) (id int, result []game, err error) {
	inp_pattern, err := regexp.Compile(`Game (\d+): (.*)`)
	if err != nil {
		return
	}

	results := inp_pattern.FindSubmatch([]byte(game_input))
	id, err = strconv.Atoi(string(results[1]))
	if err != nil {
		return
	}

	game_pattern, err := regexp.Compile(`(\d+)\s([a-z]+)`)
	game_status := string(results[2])

	for _, sets := range strings.Split(game_status, ";") {
		matches := game_pattern.FindAllSubmatch([]byte(sets), -1)
		for _, set := range matches {
			count, _ := strconv.Atoi(string(set[1]))
			game_output := game{color: string(set[2]), count: count}
			// result[i] = append(result[i], game{})
			result = append(result, game_output)
		}
	}
	return
}

func is_game_valid(game_input []game) bool {
	for _, set := range game_input {
		if set.count > rule[set.color] {
			return false
		}
	}
	return true
}

func minimum_set(game_input []game) int {
	min_val := map[string]int{
		"red":   0,
		"blue":  0,
		"green": 0,
	}
	for _, set := range game_input {
		if set.count > min_val[set.color] {
			min_val[set.color] = set.count
		}
	}
	return min_val["red"] * min_val["blue"] * min_val["green"]
}
