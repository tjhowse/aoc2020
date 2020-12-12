package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

const yMax = 92
const xMax = 91

// const yMax = 10
// const xMax = 10

type state struct {
	s [yMax][xMax]byte
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func (s *state) adjCount(x, y int, debug bool) int {
	var count int
	for j := max(y-1, 0); j <= min(y+1, xMax); j++ {
		for i := max(x-1, 0); i <= min(x+1, 90); i++ {
			if i == x && j == y {
				continue
			}
			if debug {
				log.Print("Checking ", i, j, " : ", s.s[j][i])
			}
			if s.s[j][i] == 2 {
				count++
			}
		}
	}
	return count
}

func (s *state) visDirection(x, y, xStep, yStep int, debug bool) byte {
	j := y + yStep
	i := x + xStep
	for j >= 0 && j < yMax && i >= 0 && i < xMax {
		if debug {
			log.Print("Looking at ", i, " ", j, ": ", s.s[j][i])
		}
		switch s.s[j][i] {
		case 1:
			return 0
		case 2:
			return 1
		}
		i += xStep
		j += yStep
	}
	return 0
}

func (s *state) visCount(x, y int, debug bool) int {
	var total int
	total += int(s.visDirection(x, y, 0, -1, debug))  // N
	total += int(s.visDirection(x, y, 1, -1, debug))  // NE
	total += int(s.visDirection(x, y, 1, 0, debug))   // E
	total += int(s.visDirection(x, y, 1, 1, debug))   // SE
	total += int(s.visDirection(x, y, 0, 1, debug))   // S
	total += int(s.visDirection(x, y, -1, 1, debug))  // SW
	total += int(s.visDirection(x, y, -1, 0, debug))  // W
	total += int(s.visDirection(x, y, -1, -1, debug)) // NW
	return total
}

func (s *state) step() state {
	next := state{}
	for y, row := range s.s {
		for x, seat := range row {
			switch seat {
			case 0:
				// Floor
				next.s[y][x] = 0
			case 1:
				// Empty seat
				if s.adjCount(x, y, false) == 0 {
					next.s[y][x] = 2
				} else {
					next.s[y][x] = 1
				}
			case 2:
				// Filled seat
				if s.adjCount(x, y, false) >= 4 {
					next.s[y][x] = 1
				} else {
					next.s[y][x] = 2
				}
			}
		}
	}
	return next
}

func (s *state) stepVis() state {
	next := state{}
	for y, row := range s.s {
		for x, seat := range row {
			switch seat {
			case 0:
				// Floor
				next.s[y][x] = 0
			case 1:
				// Empty seat
				if s.visCount(x, y, false) == 0 {
					next.s[y][x] = 2
				} else {
					next.s[y][x] = 1
				}
			case 2:
				// Filled seat
				if s.visCount(x, y, false) >= 5 {
					next.s[y][x] = 1
				} else {
					next.s[y][x] = 2
				}
			}
		}
	}
	return next
}

func (s *state) print() {
	for _, row := range s.s[0:10] {
		for _, seat := range row[0:10] {
			fmt.Printf("%d", seat)
		}
		// log.Print(row[0:10])
		fmt.Println()
	}
	log.Print("----------------")
}

func (s *state) printLetters() {
	for _, row := range s.s[0:10] {
		for _, seat := range row[0:10] {
			var op string
			switch seat {
			case 0:
				op = "."
			case 1:
				op = "L"
			case 2:
				op = "#"
			}
			fmt.Printf("%s", op)
		}
		// log.Print(row[0:10])
		fmt.Println()
	}
	fmt.Println("----------")
}

func (s *state) count() int {
	var count int
	for _, row := range s.s {
		for _, seat := range row {
			if seat == 2 {
				count++
			}
		}
	}
	return count
}

func (s *state) adjVis() state {
	result := state{}
	for y, row := range s.s {
		for x := range row {
			result.s[y][x] = byte(s.adjCount(x, y, false))
		}
	}
	return result
}
func (s *state) visVis() state {
	result := state{}
	for y, row := range s.s {
		for x := range row {
			result.s[y][x] = byte(s.visCount(x, y, false))
		}
	}
	return result
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

func load() {
	fileContents := fileToSlice("input")
	initial := state{}
	for y, line := range fileContents {
		for x, s := range line {
			switch s {
			case '.':
				initial.s[y][x] = 0
			case 'L':
				initial.s[y][x] = 1
			case '#':
				initial.s[y][x] = 2
			}
		}
	}
	// initial.print()
	next := initial.step()
	// next.print()
	// next = next.step()
	// next.print()
	// log.Print(next.adjCount(1, 0, true))
	// vis := next.adjVis()
	// vis.print()
	// next = next.step()
	// next.print()
	// next = next.step()
	// next = next.step()

	var prevState state
	prevState = next
	for {
		next = next.step()
		if prevState == next {
			break
		}
		prevState = next
	}
	log.Print(next.count())

}
func loadVis() {
	fileContents := fileToSlice("input")
	initial := state{}
	for y, line := range fileContents {
		for x, s := range line {
			switch s {
			case '.':
				initial.s[y][x] = 0
			case 'L':
				initial.s[y][x] = 1
			case '#':
				initial.s[y][x] = 2
			}
		}
	}
	// initial.print()
	next := initial.stepVis()
	// next.printLetters()
	// log.Print("Debug:")
	// debug := next.visVis()
	// debug.print()
	// next = next.stepVis()
	// next.printLetters()
	// log.Print("Debug:")
	// debug = next.visVis()
	// debug.print()
	// next = next.stepVis()
	// next.printLetters()
	// log.Print("Debug:")
	// debug = next.visVis()
	// debug.print()
	// next = next.stepVis()
	// next.print()
	// next = next.stepVis()
	// next.print()
	// next.visCount(0, 1, true)
	// next = next.stepVis()
	// next.print()
	// next = next.stepVis()
	// next.print()
	var prevState state
	prevState = next
	for {
		next = next.stepVis()
		if prevState == next {
			break
		}
		prevState = next
	}
	// next.printLetters()
	log.Print(next.count())

}

func main() {
	// runtime.SetBlockProfileRate(1)
	load()
	loadVis()
}
