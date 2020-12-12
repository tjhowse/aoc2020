package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

type seq struct {
	a []uint8
}

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

func contains(array []uint8, check uint8) bool {
	for _, i := range array {
		if i == check {
			return true
		}
	}
	return false
}

func load() {
	fileContents := fileToSlice("input")
	jolts := []uint8{0}
	for _, line := range fileContents {
		i, err := strconv.ParseUint(line, 10, 8)
		if err != nil {
			log.Fatal("Couldn't parse jolts", err)
		}
		jolts = append(jolts, uint8(i))
	}
	dist := make(map[uint8]int)
	sort.Slice(jolts, func(i, j int) bool { return jolts[i] < jolts[j] })
	jolts = append(jolts, jolts[len(jolts)-1]+3)
	// sort.Ints(jolts)
	// log.Print(jolts)
	for i := 0; i < len(jolts)-1; i++ {
		// log.Print(i)
		diff := jolts[i+1] - jolts[i]
		// log.Print(diff)
		if diff > 3 {
			log.Fatal("Ohno")
		}
		dist[diff]++
	}
	// We start at zero // I don't understand this.
	dist[jolts[0]]++
	log.Print(dist[1] * dist[3])

	// https://old.reddit.com/r/adventofcode/comments/ka8z8x/2020_day_10_solutions/gfcxuxf
	// I got stumped on this one, with many dead ends and much frustration. I looked up someone
	// else's solution on reddit. I re-wrote it and go and understood how it worked, I think.

	// This stores the number of ways of reaching this point in the sequence.
	paths := make(map[uint8]int)
	// There's only one way to reach the start.
	paths[0] = 1
	// For all joltages, starting at zero
	for _, j := range jolts {
		// Look at 1, 2 and 3 up from the current joltage
		for i := 1; i < 4; i++ {
			// If there's a joltage in our list that matches current + 1,2 or 3
			if contains(jolts, j+uint8(i)) {
				// Increase the count of ways of reaching that joltage by the number of
				// ways we could reach the current joltage. Multiple paths to this joltage
				// will increase its path count by their path counts.
				paths[j+uint8(i)] += paths[j]
			}
		}
	}
	log.Print(paths[jolts[len(jolts)-1]])
}

func main() {
	load()
}
