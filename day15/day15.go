package main

import (
	"bufio"
	"log"
	"os"
)

type a struct {
	a int
}

const fieldLen = 36

func fileToSlice(filename string) []string {
	var contents []string
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return contents
}

func do(limit int64, numbers []int64) int64 {
	spoken := make(map[int64]int64)
	var said int64 = -1
	for i := int64(1); i <= limit; i++ {
		if i < int64(len(numbers)+1) {
			if said >= 0 {
				spoken[said] = i
			}
			said = numbers[i-1]
		} else {
			if spoken[said] == 0 {
				// We haven't seen the previously said number before
				spoken[said] = i
				said = 0
			} else {
				t := spoken[said]
				spoken[said] = i
				said = i - t
			}
		}
	}
	return said
}

func load() {
	// said := do(2020, []int64{6, 3, 15, 13, 1, 0})
	// log.Print(" Said: ", said)
	// said := do(30000000, []int64{1, 3, 2})
	// log.Print(" Said: ", said)
	said := do(30000000, []int64{6, 3, 15, 13, 1, 0})
	log.Print(" Said: ", said)
}

func main() {
	load()
}
