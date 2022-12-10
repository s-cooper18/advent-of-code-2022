package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Instruction struct {
	quantity  int
	old_index int
	new_index int
}

func read_file(filename string) []string {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Printf("Could not read this file due to this %s error\n", err)
	}

	fileContent := string(file)

	split_content := strings.Split(fileContent, "\n\n")

	return split_content
}

func compute_col_indices(last_line string) []int {
	num_cols := len(last_line) - strings.Count(last_line, " ")

	col_indices := make([]int, num_cols)

	for i := 0; i < num_cols; i++ {
		col_indices[i] = strings.Index(last_line, fmt.Sprint(i+1))
	}

	return col_indices
}

func add_row_to_piles(piles [][]string, this_row string, col_indices []int) {
	for i := 0; i < len(col_indices); i++ {
		index := col_indices[i]
		elem := string(this_row[index])

		if elem != " " {
			piles[i] = append(piles[i], elem)
		}
	}
}

func parse_piles(this_string string) [][]string {
	lines := strings.Split(this_string, "\n")
	col_indices := compute_col_indices(lines[len(lines)-1])

	piles := make([][]string, len(col_indices))

	n_lines := len(lines)

	for i := 1; i < n_lines; i++ {
		row_index := n_lines - i - 1
		add_row_to_piles(piles, lines[row_index], col_indices)
	}

	return piles
}

func map_line_to_instruction(line string) Instruction {
	segments := strings.Split(line, " ")
	quantity, _ := strconv.Atoi(segments[1])
	old_location, _ := strconv.Atoi(segments[3])
	new_location, _ := strconv.Atoi(segments[5])
	return Instruction{quantity: quantity, old_index: old_location - 1, new_index: new_location - 1}
}

func parse_instrcutions(instruction_str string) []Instruction {
	lines := strings.Split(instruction_str, "\n")
	n_lines := len(lines)
	instructions := make([]Instruction, n_lines)
	for i := 0; i < n_lines; i++ {
		instructions[i] = map_line_to_instruction(lines[i])
	}
	return instructions
}

func apply_instructions(piles [][]string, instructions []Instruction) {
	for _, instruction := range instructions {
		old_pile_height := len(piles[instruction.old_index])

		min_index := 0
		if old_pile_height > instruction.quantity {
			min_index = old_pile_height - instruction.quantity
		}

		elems_to_move := piles[instruction.old_index][min_index:old_pile_height]
		fmt.Printf("%s\n", elems_to_move)

		for _, elem := range elems_to_move {
			piles[instruction.new_index] = append(piles[instruction.new_index], elem)
		}

		piles[instruction.old_index] = piles[instruction.old_index][0:min_index]
	}
}

func get_top_of_each_stack(piles [][]string) []string {
	tops := make([]string, len(piles))
	for i, pile := range piles {
		if len(pile) > 0 {
			tops[i] = pile[len(pile)-1]
		}

	}
	return tops
}

func main() {
	split_content := read_file("input.txt")
	piles := parse_piles(split_content[0])
	instructions := parse_instrcutions(split_content[1])

	apply_instructions(piles, instructions)
	tops := get_top_of_each_stack(piles)

	joined := strings.Join(tops, "")

	fmt.Printf("%s\n", joined)

}
