package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type bag struct {
	count  int64
	colour string
	sub    []bag
}

func (b *bag) containsGold(bags []bag) bool {
	for _, woo := range b.sub {
		if woo.colour == "shiny gold" {
			return true
		}
		for _, boo := range bags {
			if boo.colour == woo.colour {
				if boo.containsGold(bags) {
					return true
				}
			}
		}
	}
	return false
}

func (b *bag) countBags(bags []bag) int64 {
	var total int64
	for _, woo := range b.sub {
		total += woo.count
		for _, boo := range bags {
			if boo.colour == woo.colour {
				total += woo.count * boo.countBags(bags)
			}
		}
	}
	return total
}

func load() {
	// var p []groupAnswers
	// newP := groupAnswers{}
	// newP.init()
	// p = append(p, newP)
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var bags []bag
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		newBag := bag{}
		newBag.colour = s[0] + " " + s[1]
		if s[4] != "no" {
			for i := 4; i < len(s); i += 4 {
				newSubBag := bag{}
				newSubBag.colour = s[i+1] + " " + s[i+2]
				newSubBag.count, _ = strconv.ParseInt(s[i], 10, 64)
				newBag.sub = append(newBag.sub, newSubBag)
			}
		}
		// log.Print(newBag)
		bags = append(bags, newBag)
	}
	count := 0
	for _, b := range bags {
		if b.containsGold(bags) {
			count++
		}
	}
	log.Print(count)

	for _, b := range bags {
		if b.colour != "shiny gold" {
			continue
		}
		log.Print(b.countBags(bags))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	load()
}
