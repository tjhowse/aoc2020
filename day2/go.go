package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var wrongPasswords int64
	var wrongPasswordsPart2 int64
	var lines int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines++
		// fmt.Println(scanner.Text())
		// 3-8 h: kzkhgrffz
		s := strings.Split(scanner.Text(), " ")
		minmax := strings.Split(s[0], "-")
		min, _ := strconv.ParseInt(minmax[0], 10, 64)
		max, _ := strconv.ParseInt(minmax[1], 10, 64)
		letter := rune(s[1][0])
		password := s[2]
		var count int64
		for _, l := range password {
			// fmt.Println(l)
			if l == letter {
				count++
			}
		}
		if !((min <= count) && (count <= max)) {
			wrongPasswords++
		}
		first := rune(password[min-1]) == letter
		second := rune(password[max-1]) == letter
		if !(first != second) {
			wrongPasswordsPart2++

		}
	}
	fmt.Println(lines - wrongPasswords)
	fmt.Println(lines - wrongPasswordsPart2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
