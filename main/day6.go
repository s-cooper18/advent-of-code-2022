package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func read_file(filename string) string {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Printf("Could not read this file due to this %s error\n", err)
	}

	fileContent := string(file)

	return fileContent
}

func is_unique(char_list []string) (int, bool) {
	index_seen := make(map[string]int)
	for index, val := range char_list {
		if index_seen[val] == 0 {
			index_seen[val] = (index + 1) // count from 1 to k index
		} else {
			return index_seen[val], false
		}
	}
	return 1, true
}

func check_kmers(all_chars []string, k int) (int, bool) {
	n_kmers := len(all_chars) - (k - 1)
	var err bool
	for i := 0; i < n_kmers; {
		end_index := i + k
		next_increment, unique := is_unique(all_chars[i:end_index])
		if unique {
			return end_index, err
		} else {
			i = i + next_increment
		}
	}
	err = true
	return -1, err

}

func main() {
	content := read_file("input.txt")
	char_list := strings.Split(content, "")
	end_index, _ := check_kmers(char_list, 14)
	fmt.Printf("Index: %d\n", end_index)
}
