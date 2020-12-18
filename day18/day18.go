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

func doMulti(expr string) bool {
	// Returns true if there are no additions or brackets in the string
	for _, c := range expr {
		if c == '(' || c == ')' || c == '+' {
			return false
		}
	}
	return true
}

func eval(expr string) (result int64) {
	var acc int64
	var oper rune
	// log.Print("Evaluating: ", expr)
	// expr = evalPlus(expr)
	// log.Print("After: ", expr)
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
					n = eval(expr[i+1 : j])
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
				// expr = expr[:i]
			} else if oper == '+' {
				acc += n
			} else {
				acc = n
			}
		}
	}
	result = acc
	// log.Print("Result: ", result)
	return result
}

func isInt(r byte) bool {
	_, err := strconv.ParseInt(string(r), 10, 64)
	return err == nil
}

// func evalPlus(expr string) (result string) {
// 	var acc int64
// 	var oper rune
// 	log.Print("Evaluating: ", expr)
// 	// expr = evalPlus(expr)
// 	// log.Print("After: ", expr)
// 	// for i, c := range expr {
// 	for i := 0; i < len(expr); i++ {
// 		c := rune(expr[i])
// 		n := int64(0)
// 		// log.Print("Rune: ", string(c))
// 		switch c {
// 		case ' ':
// 			continue
// 		case '*':
// 			oper = c
// 		case '+':
// 			oper = c
// 		case '(':
// 			var depth int64 = 1
// 			for j := i + 1; j < len(expr); j++ {
// 				if expr[j] == '(' {
// 					depth++
// 				} else if expr[j] == ')' {
// 					depth--
// 				}
// 				// log.Print("Building: ", expr[i+1:j])
// 				// log.Print("Depth: ", depth)
// 				if depth == 0 {
// 					n = eval(expr[i+1 : j])
// 					i = j + 1
// 					break
// 				}
// 			}
// 		case ')':
// 			continue
// 		default:
// 			// A number
// 			num, err := strconv.ParseInt(string(c), 10, 64)
// 			if err != nil {
// 				log.Fatal()
// 			}
// 			n = num
// 		}
// 		if n != 0 {
// 			// log.Print("Doing maths with: ", n)
// 			if oper == '*' {
// 				acc *= n
// 				// expr = expr[:i]
// 			} else {
// 				acc += n
// 			}
// 		}
// 	}
// 	result = acc
// 	log.Print("Result: ", result)
// 	return result
// }

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
		a := eval(line)
		// log.Print(a)
		total += a

	}
	log.Print(total)

}

func main() {
	load()
}
