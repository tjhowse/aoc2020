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
	x, y   int
	wX, wY int
	h      heading
}

func rotate(deg int, v vec) vec {
	var result vec
	m := math.Sqrt(math.Pow(float64(v.x), 2) + math.Pow(float64(v.y), 2))
	angle := math.Atan2(float64(v.y), float64(v.x))
	result.x = int(math.Round(m * math.Cos(angle+float64(deg)*(math.Pi/180))))
	result.y = int(math.Round(m * math.Sin(angle+float64(deg)*(math.Pi/180))))
	return result
}

func (s *ship) doInstruction(i pa) {
	switch i.d {
	case byte('R'):
		// s.h.turn(i.d, i.n)
		var v vec
		v.x = s.wX
		v.y = s.wY
		n := rotate(-i.n, v)
		s.wX = n.x
		s.wY = n.y
	case byte('L'):
		// s.h.turn(i.d, i.n)
		var v vec
		v.x = s.wX
		v.y = s.wY
		n := rotate(i.n, v)
		s.wX = n.x
		s.wY = n.y
	case byte('N'):
		s.wY += i.n
	case byte('S'):
		s.wY -= i.n
	case byte('E'):
		s.wX += i.n
	case byte('W'):
		s.wX -= i.n
	case byte('F'):
		for j := 0; j < i.n; j++ {
			s.x += s.wX
			s.y += s.wY
		}
	default:
		log.Fatal("Bad command: ", i.d)
	}
	// log.Print("Pos: ", s.x, ", ", s.y)
	// log.Print(" wp: ", s.wX, ", ", s.wY)
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
	s.wX = 10
	s.wY = 1
	for _, c := range commands {
		s.doInstruction(c)
	}
	log.Print(s.x, ", ", s.y)
	log.Print(math.Abs(float64(s.x)) + math.Abs(float64(s.y)))
}
func main() {
	load()
	// var v vec
	// v.x = 10
	// v.y = 0
	// log.Print(v)
	// log.Print(rotate(90, v))
}
