package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func reSubMatchMap(r *regexp.Regexp, str string) map[string]string {
	match := r.FindStringSubmatch(str)
	subMatchMap := make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i != 0 {
			subMatchMap[name] = match[i]
		}
	}

	return subMatchMap
}

func eval(expr string) (result int64) {
	var acc int64
	var oper rune
	// log.Print("Evaluating: ", expr)
	// for i, c := range expr {
	for i := 0; i < len(expr); i++ {
		c := rune(expr[i])
		n := int64(0)
		// log.Print("Rune: ", string(c))
		switch c {
		case ' ':
			continue
		case '*':
			oper = c
		case '+':
			oper = c
		case '(':
			var depth int64 = 1
			for j := i + 1; j < len(expr); j++ {
				if expr[j] == '(' {
					depth++
				} else if expr[j] == ')' {
					depth--
				}
				// log.Print("Building: ", expr[i+1:j])
				// log.Print("Depth: ", depth)
				if depth == 0 {
					n = eval(expr[i+1 : j+1])
					i = j + 1
					break
				}
			}
		case ')':
			continue
		default:
			// A number
			num, err := strconv.ParseInt(string(c), 10, 64)
			if err != nil {
				log.Fatal()
			}
			n = num
		}
		if n != 0 {
			// log.Print("Doing maths with: ", n)
			if oper == '*' {
				acc *= n
			} else {
				acc += n
			}
		}
	}
	result = acc
	// log.Print("Result: ", result)
	return result
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

func load() {
	// log.Print(-21 % 5)
	fileContents := fileToSlice("input")
	var total int64
	for _, line := range fileContents {

		// log.Print(eval(line))
		// log.Print("---------")
		total += eval(line)

	}
	log.Print(total)

}

func main() {
	load()
}
