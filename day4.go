package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func read_file(filename string) []string {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Printf("Could not read this file due to this %s error\n", err)
	}

	fileContent := string(file)

	split_content := strings.Split(fileContent, "\n")

	return split_content
}

type NumberSet struct {
	lower int
	upper int
}

func parse_line(line string) []NumberSet {
	number_set_strs := strings.Split(line, ",")
	num_sets := make([]NumberSet, 2)
	for i, num_set_str := range number_set_strs {
		split_nums := strings.Split(num_set_str, "-")
		lower, err := strconv.Atoi(split_nums[0])
		upper, err := strconv.Atoi(split_nums[1])

		if err != nil {
			fmt.Printf("Could not read this file due to this %s error\n", err)
		}
		num_sets[i] = NumberSet{lower: lower, upper: upper}
	}

	return num_sets
}

func compare_ranges(num1 NumberSet, num2 NumberSet) bool {

	if (num1.lower <= num2.lower) && ((num1.upper >= num2.upper) || (num1.upper >= num2.lower)) {
		return true
	} else if (num2.lower <= num1.lower) && ((num2.upper >= num1.upper) || (num2.upper >= num1.lower)) {
		return true
	}

	return false

}

func main() {
	result := read_file("input.txt")

	total := 0
	for _, line := range result {
		result := parse_line(line)
		fmt.Printf("%d\n", result)
		if compare_ranges(result[0], result[1]) {
			total++
		}
	}

	fmt.Printf("%d\n", total)
}
