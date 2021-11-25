package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	f []int64
	s []int64
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

func sliceEqual(s []int64, z []int64) bool {
	// Returns true if the two slices match
	if len(s) != len(z) {
		return false
	}
	for i, c := range s {
		if z[i] != c {
			return false
		}
	}
	return true
}

func validate(s []int64, rules map[rulePattern]int64, ruleZero []int64) bool {
	// log.Print(s)
	if sliceEqual(s, ruleZero) {
		return true
	}
	for i := 0; i < len(s)-1; i++ {
		rp := rulePattern{s[i], s[i+1]}
		if rules[rp] != 0 {
			new := []int64{}
			new = append(new, s[:i]...)
			// new += s[:i]
			new = append(new, rules[rp])
			new = append(new, s[i+2:]...)
			if validate(new, rules, ruleZero) {
				return true
			}
			// if !sliceEqual(s[:i], ruleZero[:i]) {
			// 	return false
			// }
		}
		// break
	}

	return false
}

func validate2(s []int64, rules map[rulePattern]int64, ruleZero []int64) bool {
	// for _, c := range ruleZero {

	// }
	return true
}

type rulePattern [2]int64

func stringToRulePattern(in string) (result rulePattern) {
	i := 0
	numbers := strings.Split(in, " ")
	for _, c := range numbers {
		b, err := strconv.ParseInt(c, 10, 64)
		if err != nil {
			log.Fatal("Failed to parse rule pattern: \"", in, "\"")
		}
		result[i] = b
		i++
	}
	return result
}

func load() {
	fileContents := fileToSlice("input")
	var mode bool
	rules := make(map[rulePattern]int64)
	ruleZero := []int64{}
	var aRule, bRule int64
	var total int64
	for _, line := range fileContents {
		if line == "" {
			mode = true
			// for k, v := range rules {
			// 	log.Print(k, ": ", v)
			// }
			log.Print("Processing strings")
			continue
		}
		if !mode {
			// Collecting rules
			split := strings.Split(line, ":")
			ruleNumber, err := strconv.ParseInt(split[0], 10, 64)
			if err != nil {
				log.Fatal()
			}
			if split[1] == " \"a\"" {
				// rules[split[0]] = "a"
				aRule = ruleNumber
			} else if split[1] == " \"b\"" {
				// rules[split[0]] = "b"
				bRule = ruleNumber
			} else {
				if ruleNumber != 0 {
					split2 := strings.Split(split[1], "|")
					for _, c := range split2 {
						c = strings.TrimSpace(c)
						rp := stringToRulePattern(c)
						rules[rp] = ruleNumber
					}
					// rules[split[0]] = split[1][1:]
				} else {
					// Special case for rule zero.
					trimmed := strings.TrimSpace(split[1])
					split2 := strings.Split(trimmed, " ")
					for _, c := range split2 {
						b, err := strconv.ParseInt(c, 10, 64)
						if err != nil {
							log.Fatal("Can't parse rule zero")
						}
						ruleZero = append(ruleZero, b)
					}
				}
			}
		} else {
			// Checking rules
			intLine := []int64{}
			for _, c := range line {
				if c == 'a' {
					intLine = append(intLine, aRule)
				} else {
					intLine = append(intLine, bRule)
				}
			}
			log.Print("Checking ", intLine)
			if validate(intLine, rules, ruleZero) {

				// log.Print("Good: ", intLine)
				log.Print("Yes")
				total++
			} else {
				log.Print("No")
			}
		}
	}
	log.Print(total)

}

func main() {
	load()
}
