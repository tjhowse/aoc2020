package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func convertStringtoNumber(s string) int64 {
	var b string
	for _, l := range s {
		switch l {
		case 'F':
			b = b + "0"
		case 'B':
			b = b + "1"
		case 'L':
			b = b + "0"
		case 'R':
			b = b + "1"
		default:
			log.Fatal("Encountered dud letter")
		}
	}
	var a int64
	a, err := strconv.ParseInt(b, 2, 64)
	if err != nil {
		log.Fatal("Couldn't parse binary number")
	}
	return a
}

func seatID(s string) int64 {
	row := convertStringtoNumber(s[:7])
	column := convertStringtoNumber(s[7:])
	return row*8 + column
}

func load() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var max int64
	var all []int
	for scanner.Scan() {
		s := scanner.Text()
		id := seatID(s)
		all = append(all, int(id))
		if id > max {
			max = id
		}
	}
	sort.Ints(all)
	log.Print(max)
	for i := len(all) - 1; i > 0; i-- {
		if all[i] != all[i-1]+1 {
			log.Print(all[i] - 1)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	load()
}
