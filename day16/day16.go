package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const fieldCount = 20

type rule struct {
	name         string
	r1Min, r1Max int64
	r2Min, r2Max int64
}

func (r *rule) parseRange(input string) {
	reg := regexp.MustCompile(` (\d*)-(\d*) or (\d*)-(\d*)`)
	a := reg.FindStringSubmatch(input)
	var err error
	r.r1Min, err = strconv.ParseInt(a[1], 10, 64)
	if err != nil {
		log.Fatal("Beans")
	}
	r.r1Max, err = strconv.ParseInt(a[2], 10, 64)
	if err != nil {
		log.Fatal("Beans")
	}
	r.r2Min, err = strconv.ParseInt(a[3], 10, 64)
	if err != nil {
		log.Fatal("Beans")
	}
	r.r2Max, err = strconv.ParseInt(a[4], 10, 64)
	if err != nil {
		log.Fatal("Beans")
	}
}

func (r *rule) validateNumber(n int64) bool {
	if n >= r.r1Min && n <= r.r1Max {
		return true
	}
	if n >= r.r2Min && n <= r.r2Max {
		return true
	}
	return false
}

type ticket struct {
	nums []int64
}

func (t *ticket) loadNums(input string) {
	ns := strings.Split(input, ",")
	for _, n := range ns {
		i, err := strconv.ParseInt(n, 10, 64)
		if err != nil {
			log.Fatal()
		}
		t.nums = append(t.nums, i)
	}
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

func checkAllAreLenOne(input [fieldCount]map[string]bool) bool {
	for i := 0; i < fieldCount; i++ {
		if len(input[i]) != 1 {
			return false
		}
	}
	return true
}

func load() {
	fileContents := fileToSlice("input")
	var mode int64
	// 0 - rules
	// 1 - my ticket
	// 2 - nearby tickets
	rules := []rule{}
	myTicket := ticket{}
	nearbyTickets := []ticket{}

	for _, line := range fileContents {
		if line == "" {
			continue
		}
		command := strings.Split(line, ":")
		if command[0] == "your ticket" {
			mode = 1
			continue
		} else if command[0] == "nearby tickets" {
			mode = 2
			continue
		}
		switch mode {
		case 0:
			newRule := rule{}
			newRule.name = command[0]
			newRule.parseRange(command[1])
			rules = append(rules, newRule)
		case 1:
			myTicket.loadNums(command[0])
		case 2:
			newTicket := ticket{}
			newTicket.loadNums(command[0])
			nearbyTickets = append(nearbyTickets, newTicket)
		}
		// strconv.ParseInt(xxx, 10, 64)
		// log.Print(line)
	}
	var rate int64
	validTickets := []ticket{}

	for _, t := range nearbyTickets {
		var validTicket bool = true
		for _, n := range t.nums {
			var valid bool = false
			for _, r := range rules {
				if r.validateNumber(n) {
					valid = true
					break
				}
			}
			if !valid {
				rate += n
				validTicket = false
			}
		}
		if validTicket {
			validTickets = append(validTickets, t)
		}
	}
	log.Print(rate)

	poss := [fieldCount]map[string]bool{}
	imPoss := [fieldCount]map[string]bool{}
	for i := 0; i < fieldCount; i++ {
		poss[i] = make(map[string]bool)
		imPoss[i] = make(map[string]bool)

	}

	for _, t := range validTickets {
		for i, v := range t.nums {
			for _, rule := range rules {
				if rule.validateNumber(v) {
					poss[i][rule.name] = true
				} else {
					imPoss[i][rule.name] = true
				}
			}
		}
	}
	// Filter out the impossibles
	for i := 0; i < fieldCount; i++ {
		for k := range imPoss[i] {
			delete(poss[i], k)
		}
	}

	// Now condense the combinations
	for !checkAllAreLenOne(poss) {
		for i := 0; i < fieldCount; i++ {
			if len(poss[i]) == 1 {
				// If this field has only one possibility, it must be right.
				for k := 0; k < fieldCount; k++ {
					if i == k {
						continue
					}
					for field := range poss[i] {
						// This only has one element, but I don't know another way of getting
						// the keys of this map
						delete(poss[k], field)
					}
				}
			}
		}
	}
	var result int64 = 1
	for i := 0; i < fieldCount; i++ {
		for field := range poss[i] {
			if len(field) < len("departure") {
				continue
			}
			if field[:len("departure")] == "departure" {
				log.Print("Field: ", i, ": ", poss[i])
				result *= myTicket.nums[i]
			}
		}
	}
	log.Print(result)
}

func main() {
	load()
}
