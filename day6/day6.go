package main

import (
	"bufio"
	"log"
	"os"
)

func load() {
	var p []groupAnswers
	newP := groupAnswers{}
	newP.init()
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
			newP := groupAnswers{}
			newP.init()
			p = append(p, newP)
		} else {
			p[len(p)-1].parse(s)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	var total int
	for _, s := range p {
		// log.Print(s.count())
		total += s.count()
	}
	log.Print(total)
}

func main() {
	load()
	// s := "0123456"
	// log.Print(s[len(s)-2:])
}
