package main

import (
	"bufio"
	"log"
	"os"
	"runtime/pprof"
	"sort"
	"strconv"
)

type seq struct {
	a []int
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

// func getCombo(array []int)

func getLinkableIndices(index int, array []int) []int {
	linkables := []int{}
	for i := index; i < len(array); i++ {
		jolt := array[i]
		if jolt <= array[index] {
			continue
		}
		if jolt-array[index] > 3 {
			break
		}
		linkables = append(linkables, i)
	}
	return linkables
}

func do(result map[[100]int]struct{}, input []int, part seq, index int) {
	part.a = append(part.a, input[index])
	next := getLinkableIndices(index, input)
	if len(next) == 0 {
		var k [100]int
		copy(k[:], part.a[:])
		result[k] = struct{}{}
		if len(result) > 1000000 {
			f, _ := os.Create("dump")
			defer f.Close()
			pprof.WriteHeapProfile(f)
			log.Fatal("too big, ", len(result))
		}

		return
	}
	for _, i := range next {
		// log.Print(part)
		split := part
		do(result, input, split, i)
	}
}

func load() {
	fileContents := fileToSlice("input")
	jolts := []int{0}
	for _, line := range fileContents {
		i, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatal("Couldn't parse jolts")
		}
		jolts = append(jolts, int(i))
	}
	dist := make(map[int]int)
	sort.Ints(jolts)
	for i := 0; i < len(jolts)-1; i++ {
		// log.Print(i)
		diff := jolts[i+1] - jolts[i]
		if diff > 3 {
			log.Fatal("Ohno")
		}
		dist[diff]++
	}
	// We start at zero
	dist[jolts[0]]++
	// We end at +3 above the last one
	dist[3]++
	log.Print(dist[1] * dist[3])

	m := make(map[[100]int]struct{})
	var s seq
	s.a = append(s.a, 0)
	do(m, jolts, s, 0)
	log.Print(len(m))
	// for _, i := range jolts {
	// 	log.Print(i)
	// }

}

func main() {
	// runtime.SetBlockProfileRate(1)
	load()
}
