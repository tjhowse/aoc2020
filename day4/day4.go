package main

import (
	"bufio"
	"log"
	"os"
)

func load() {
	var p []passportStrict
	newP := passportStrict{}
	p = append(p, newP)
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			newP := passportStrict{}
			p = append(p, newP)
		} else {
			p[len(p)-1].parse(s)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	var valid int
	for _, s := range p {
		if s.validate() {
			valid++
		}
	}
	log.Print(valid)
}

func main() {
	load()
	// s := "0123456"
	// log.Print(s[len(s)-2:])
}
