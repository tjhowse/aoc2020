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

// func getCombo(array []int)

func getLinkableIndices(index uint8, array []uint8) []uint8 {
	linkables := []uint8{}
	for i := index; i < uint8(len(array)); i++ {
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

func do(result map[[100]uint8]struct{}, input []uint8, part seq, index uint8) int {
	var total int
	// m := make(map[[100]uint8]struct{})
	part.a = append(part.a, input[index])
	next := getLinkableIndices(index, input)
	if len(next) == 0 {
		var k [100]uint8
		copy(k[:], part.a[:])
		result[k] = struct{}{}
		// m[k] = struct{}{}

		if len(result) > 10000000 {
			f, _ := os.Create("dump")
			defer f.Close()
			pprof.WriteHeapProfile(f)
			log.Fatal("too big, ", len(result))
		}

		return len(result)
	}
	for _, i := range next {
		// log.Print(part)
		split := part
		total += do(result, input, split, i)

	}
	return total
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
	// sort.Ints(jolts)
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

	m := make(map[[100]uint8]struct{})
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
