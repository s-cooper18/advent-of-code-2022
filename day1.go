package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func day1() {
	filename := "input.txt"
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Printf("Could not read this file due to this %s error\n", err)
	}

	fileContent := string(file)

	result := strings.Split(fileContent, "\n\n")

	totals := make([]int, len(result))
	max_values := make([]int, 3)

	for index, substring := range result {
		nums := strings.Split(substring, "\n")

		for _, value := range nums {
			parsed_num, _ := strconv.Atoi(value)
			totals[index] = totals[index] + parsed_num
		}

		if totals[index] > max_values[1] {
			if totals[index] > max_values[0] {
				max_values[1] = max_values[0]
				max_values[0] = totals[index]
			} else {
				max_values[2] = max_values[1]
				max_values[1] = totals[index]
			}
		} else if totals[index] > max_values[2] {
			max_values[2] = totals[index]
		}

	}

	fmt.Println(max_values[0] + max_values[1] + max_values[2])
}
