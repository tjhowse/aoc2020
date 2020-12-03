package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type day3 struct {
	xmax, ymax int
	input      [31][323]bool
}

func (d *day3) tree(x int, y int) bool {
	return d.input[x%d.xmax][y]
}

func (d *day3) load() {
	d.xmax = 31
	d.ymax = 323
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var line int
	for scanner.Scan() {
		s := scanner.Text()
		for n, c := range s {
			d.input[n][line] = c == '#'
		}
		line++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (d *day3) ski(xStep, yStep int) int {
	var x int
	var trees int
	for y := 0; y < d.ymax; y += yStep {
		if d.tree(x, y) {
			trees++
		}
		x += xStep
	}
	return trees
}

func main() {
	var d day3
	d.load()
	fmt.Println(d.ski(3, 1))
	var m int
	m = d.ski(1, 1)
	m *= d.ski(3, 1)
	m *= d.ski(5, 1)
	m *= d.ski(7, 1)
	m *= d.ski(1, 2)
	fmt.Println(m)

}
