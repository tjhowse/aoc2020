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

func validate(s []int64, rules map[string]string) bool {
	log.Print(s)

	for i := int64(0); i < int64(len(s)); {
		// i += validate(s, rules)

	}

	return true
}

func load() {
	fileContents := fileToSlice("input")
	var mode bool
	rules := make(map[string]string)
	var aRule, bRule int64
	for _, line := range fileContents {
		if line == "" {
			mode = true
			continue
		}
		if !mode {
			// Collecting rules
			split := strings.Split(line, ":")
			if split[1] == " \"a\"" {
				rules[split[0]] = "a"
				aRule, _ = strconv.ParseInt(split[0], 10, 64)
			} else if split[1] == " \"b\"" {
				rules[split[0]] = "b"
				bRule, _ = strconv.ParseInt(split[0], 10, 64)
			} else {
				rules[split[0]] = split[1][1:]
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
			validate(intLine, rules)
		}
	}
	for k, v := range rules {
		log.Print(k, ": ", v)
	}

}

func main() {
	load()
}
