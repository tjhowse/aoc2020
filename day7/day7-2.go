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

func (b *bag) containsGold(bags map[string]bag) bool {
	for _, woo := range b.sub {
		if woo.colour == "shiny gold" {
			return true
		}
		t := bags[woo.colour]
		if t.containsGold(bags) {
			return true
		}
	}
	return false
}

func (b *bag) countBags(bags map[string]bag) int64 {
	var total int64
	for _, woo := range b.sub {
		total += woo.count
		t := bags[woo.colour]
		total += woo.count * t.countBags(bags)
	}
	return total
}

func load() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	bags := make(map[string]bag)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		colour := s[0] + " " + s[1]
		b := bag{}
		b.colour = colour
		if s[4] != "no" {
			for i := 4; i < len(s); i += 4 {
				newSubBag := bag{}
				newSubBag.colour = s[i+1] + " " + s[i+2]
				newSubBag.count, _ = strconv.ParseInt(s[i], 10, 64)
				b.sub = append(b.sub, newSubBag)
			}
		}
		bags[b.colour] = b
	}
	count := 0
	for _, v := range bags {
		if v.containsGold(bags) {
			count++
		}
	}
	log.Print(count)
	b := bags["shiny gold"]
	log.Print(b.countBags(bags))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	load()
}
