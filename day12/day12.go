package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

const (
	east  = iota // 0
	south = iota // 1
	west  = iota // 2
	north = iota // 3
)

type heading struct {
	h int
}

func (h *heading) turn(dir byte, amount int) {
	if dir == byte('R') {
		// Turning right
		h.h += amount / 90
		// log.Print("Right")
	} else if dir == byte('L') {
		// log.Print("Left")
		h.h -= amount / 90
	}
	// h.h = int(math.Abs(float64(h.h % 4)))
	if h.h < 0 {
		h.h += 4
	} else if h.h >= 4 {
		h.h -= 4
	}
	// log.Print(h.h)
}

type vec struct {
	x, y int
}

type pa struct {
	d byte
	n int
}

type ship struct {
	x, y int
	h    heading
}

func (s *ship) doInstruction(i pa) {
	switch i.d {
	case byte('R'):
		s.h.turn(i.d, i.n)
	case byte('L'):
		s.h.turn(i.d, i.n)
	case byte('N'):
		s.y += i.n
	case byte('S'):
		s.y -= i.n
	case byte('E'):
		s.x += i.n
	case byte('W'):
		s.x -= i.n
	case byte('F'):
		switch s.h.h {
		case east:
			s.x += i.n
		case south:
			s.y -= i.n
		case west:
			s.x -= i.n
		case north:
			s.y += i.n
		default:
			log.Fatal("Bad direction: ", s.h.h)
		}
	default:
		log.Fatal("Bad command: ", i.d)
	}
	// log.Print(s.x, ", ", s.y, ", ", s.h.h)
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

	var commands []pa

	for _, line := range fileContents {
		newPA := pa{}
		newPA.d = line[0]
		a, err := strconv.ParseInt(line[1:], 10, 64)
		if err != nil {
			log.Fatal("Ohno")
		}
		newPA.n = int(a)
		commands = append(commands, newPA)
	}
	// log.Print(commands)
	var s ship
	for _, c := range commands {
		s.doInstruction(c)
	}
	log.Print(s.x, ", ", s.y)
	log.Print(math.Abs(float64(s.x)) + math.Abs(float64(s.y)))
}
func main() {
	load()
}
