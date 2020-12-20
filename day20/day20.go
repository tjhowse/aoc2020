package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const imageSize = 3

// const imageSize = 12
const tileCount = imageSize * imageSize
const tileSize = 10
const (
	up    = iota // 0
	down  = iota // 1
	left  = iota // 2
	right = iota // 3
)

type dir int

type border [4 * tileSize]bool

func shift(a border, v int64, flip bool) (result border) {
	// This returns array a shifted right by v count, wrapped etc
	// if flip is true, flip the final result
	k := 0
	for i := v; ; i++ {
		i %= 4 * tileSize
		result[k] = a[i]
		k++
		if k >= 4*tileSize {
			break
		}
	}
	if !flip {
		return result
	}
	temp := border{}
	for i := 0; i < 4*tileSize; i++ {
		temp[i] = result[4*tileSize-1-i]
	}
	return temp
}

func checkTenOverlapBorder(a, b border) (result dir) {
	var count int
	for i := 0; i < 4*tileSize; i++ {
		if a[i] == b[i] {
			count++
			if count >= 10 && (i%10 == 0) {
				return dir((i / 10))
			}
		} else {
			count = 0
		}
	}
	return -1
}

type image struct {
	m [imageSize][imageSize]tile
}

type tile struct {
	m  [tileSize][tileSize]bool
	id int64
	b  border
}

func (t *tile) calcBorders() {
	// This returns an array of the boundary of this tile.

	var k int64
	// Top
	for i := 0; i < tileSize; i++ {
		t.b[k] = t.m[0][i]
		k++
	}
	// Right
	for i := 0; i < tileSize; i++ {
		t.b[k] = t.m[i][tileSize-1]
		k++
	}
	// Bottom
	for i := 0; i < tileSize; i++ {
		t.b[k] = t.m[tileSize-1][tileSize-i-1]
		k++
	}
	// Left
	for i := 0; i < tileSize; i++ {
		t.b[k] = t.m[tileSize-i-1][0]
		k++
	}
}
func (t *tile) printTile() {
	for j := 0; j < tileSize; j++ {
		for k := 0; k < tileSize; k++ {
			if t.m[j][k] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
func (t *tile) checkSharedEdge(t2 tile) int64 {
	// If the provided tile matches an edge on this tile, return the direction.
	// else return -1

	// Look for an overlapping 10 values in the border of each.
	for i := int64(0); i < tileSize*4; i += tileSize {
		temp := shift(t2.b, i, false)
		d := checkTenOverlapBorder(t.b, temp)
		if d >= 0 {
			log.Print(t2.id, " shares a border with ", t.id, " d: ", d)
		}
	}
	for i := int64(0); i < tileSize*4; i += tileSize {
		temp := shift(t2.b, i, true)
		d := checkTenOverlapBorder(t.b, temp)
		if d >= 0 {
			log.Print(t2.id, " flipped shares a border with ", t.id, " d: ", d)
		}
	}
	return -1
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
	var tiles []tile
	for i := 0; i < len(fileContents); i += 12 {
		line := fileContents[i]
		if line[:4] != "Tile" {
			log.Fatal("OHNO")
		}
		t := tile{}
		id, err := strconv.ParseInt(line[5:9], 10, 64)
		if err != nil {
			log.Fatal("Fuck! ", line[6:10])
		}
		t.id = id
		for j := 0; j < tileSize; j++ {
			for k := 0; k < tileSize; k++ {
				t.m[j][k] = fileContents[i+j+1][k] == '#'
			}
		}
		t.calcBorders()
		tiles = append(tiles, t)
	}
	for j := 0; j < tileCount; j++ {
		// log.Print(tiles[j].b)
		for k := 0; k < tileCount; k++ {
			if j == k {
				continue
			}
			tiles[k].checkSharedEdge(tiles[j])
		}
	}
	// log.Print(tiles[0].id)
	// tiles[0].printTile()

}

func main() {
	load()
}
