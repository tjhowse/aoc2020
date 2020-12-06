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
	var totalPart1 int
	var totalPart2 int
	for _, s := range p {
		totalPart1 += s.countPart1()
		totalPart2 += s.countPart2()
	}
	log.Print(totalPart1)
	log.Print(totalPart2)
}

func main() {
	load()
}
