package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	// INVALID_X =
	FOREST_SIZE = 99
)

func main() {
	input, err := os.Open("./input.txt")
	forest := make([][]int8, FOREST_SIZE)
	idx := 0
	edges := 99 + 99 + 99 + 99

	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		build_forest(scanner.Text(), forest, idx)
		idx++
	}
	edges += visible_trees(forest)
}

func build_forest(trees string, forest [][]int8, index int) {
	row := make([]int8, FOREST_SIZE)
	for idx, tree := range trees {
		row[idx] = int8(tree - '0')
	}
	forest[index] = row
}

func vis_row(tree_index int, trees []int8) bool {
	left_vis := true
	right_vis := true

	for left, right := tree_index-1, tree_index+1; right < len(trees) || left > 0; {
		if trees[left] > trees[tree_index] {
			left_vis = false
		}
		if trees[right] > trees[tree_index] {
			right_vis = false
		}

		if left_vis {
			left--
		}
		if right_vis {
			right++
		}
		if left_vis || right_vis {
			break
		}
	}
	// for _, tree := range trees[tree_index+1:] {
	//     if trees[tree_index] < tree {
	//         is_visible = false
	//         break
	//     }
	// }
	return left_vis || right_vis
}

func vis_col(tree_index, col_index int, tree [][]int8) bool {
	return true
}

func visible_trees(forest [][]int8) int {
	count := 0
	for row, trees := range forest[1:FOREST_SIZE] {
		for col, tree := range trees[1:FOREST_SIZE] {
			fmt.Println()
		}
	}
	return count
}
