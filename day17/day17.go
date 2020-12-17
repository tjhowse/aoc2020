package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"regexp"
)

const maxX = 8
const maxY = 8
const maxZ = 8

type vec2 struct {
	x, y int64
}

type vec3 struct {
	x, y, z int64
}

type state struct {
	s map[vec3]bool
}

func (s *state) getAdjacent(c vec2) []vec2 {
	result := []vec2{}

	return result
}
func (s *state) countAdjacent(c vec2) (result int64) {
	// var result int64
	for j := c.y - 1; j <= c.y+1; j++ {
		for i := c.x - 1; i <= c.x+1; i++ {
			if i == c.x && j == c.y {
				continue
			}
			if s.getValue(vec2{i, j}) {

				result++
			}
		}
	}
	return result
}

func (s *state) getValue(c vec2) bool {
	if c.x < 0 {
		// Forgive me
		c.x += 100 * maxX
	}
	if c.y < 0 {
		c.y += 100 * maxY
	}
	c.x = int64(math.Abs(float64(c.x % maxX)))
	c.y = int64(math.Abs(float64(c.y % maxY)))
	return s.s[c.y][c.x]
}

func reSubMatchMap(r *regexp.Regexp, str string) map[string]string {
	match := r.FindStringSubmatch(str)
	subMatchMap := make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i != 0 {
			subMatchMap[name] = match[i]
		}
	}

	return subMatchMap
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
	// log.Print(-21 % 5)
	fileContents := fileToSlice("input")
	s := state{}
	for y, line := range fileContents {
		// strconv.ParseInt(xxx, 10, 64)
		for x, cell := range line {
			if cell == '#' {
				s.s[y][x] = true
			}
		}
	}
	log.Print(s.countAdjacent(vec2{0, 0}))

	// var reFooBar = regexp.MustCompile(`[\w+:]{1,}(?P<fooNum>\d+),[\w+\s:]{1,}(?P<barNum>\d+)`)
	// log.Print(result)
}

func main() {
	load()
}
