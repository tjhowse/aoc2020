package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

const boop = 30
const maxX = boop
const maxY = boop
const maxZ = boop
const minX = -boop
const minY = -boop
const minZ = -boop

// const maxX = 20
// const maxY = 20
// const maxZ = 3
// const minX = -10
// const minY = -10
// const minZ = -3

type vec2 struct {
	x, y int64
}

type vec3 struct {
	x, y, z int64
}

type state struct {
	s map[vec3]bool
}

func (s *state) countAdjacent(c vec3) (result int64) {
	for k := c.z - 1; k <= c.z+1; k++ {
		for j := c.y - 1; j <= c.y+1; j++ {
			for i := c.x - 1; i <= c.x+1; i++ {
				if i == c.x && j == c.y && k == c.z {
					continue
				}
				if s.s[vec3{i, j, k}] {
					result++
				}
			}
		}
	}
	return result
}

func (s *state) step() state {
	result := state{}
	result.s = make(map[vec3]bool)

	for k := int64(minZ); k <= maxZ; k++ {
		for j := int64(minY); j <= maxY; j++ {
			for i := int64(minX); i <= maxX; i++ {
				vec := vec3{i, j, k}
				adj := s.countAdjacent(vec)
				if s.s[vec] {
					if !(adj == 2 || adj == 3) {
						result.s[vec] = false
					} else {
						result.s[vec] = true

					}
				} else {
					if adj == 3 {
						result.s[vec] = true
					} else {
						result.s[vec] = false

					}
				}
			}
		}
	}
	return result
}

func (s *state) countActive() (result int64) {
	for _, v := range s.s {
		if v {
			result++
		}
	}
	return result
}

func (s *state) print() {
	for k := int64(minZ); k <= maxZ; k++ {
		fmt.Println("\n-------------", k)
		for j := int64(minY); j <= maxY; j++ {
			for i := int64(minX); i <= maxX; i++ {
				vec := vec3{i, j, k}
				if s.s[vec] {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
	}
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
	s.s = make(map[vec3]bool)
	for y, line := range fileContents {
		// strconv.ParseInt(xxx, 10, 64)
		for x, cell := range line {
			if cell == '#' {
				vec := vec3{int64(x), int64(y), 0}
				s.s[vec] = true
			}
		}
	}
	// log.Print(s.s)
	// s.print()
	// log.Print(s.s[vec3{2, 2, 0}])
	// log.Print(s.countAdjacent(vec3{2, 2, -2}))
	s = s.step()
	s = s.step()
	s = s.step()
	s = s.step()
	s = s.step()
	s = s.step()
	log.Print(s.countActive())
	// s.print()
	// for k := int64(minZ); k <= maxZ; k++ {
	// 	fmt.Println("\n-------------", k)
	// 	for j := int64(minY); j <= maxY; j++ {
	// 		for i := int64(minX); i <= maxX; i++ {
	// 			vec := vec3{i, j, k}
	// 			if s.s[vec] {
	// 				fmt.Print("#")
	// 			} else {
	// 				fmt.Print(".")
	// 			}
	// 		}
	// 		fmt.Println()
	// 	}
	// }

	// var reFooBar = regexp.MustCompile(`[\w+:]{1,}(?P<fooNum>\d+),[\w+\s:]{1,}(?P<barNum>\d+)`)
	// log.Print(result)
}

func main() {
	load()
}
