package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

var reMaths = regexp.MustCompile(`(?P<a>\d+) (?P<o>[\+*]) (?P<b>\d+)`)
var reAdd = regexp.MustCompile(`(?P<a>\d+) (?P<o>[\+]) (?P<b>\d+)`)
var reMul = regexp.MustCompile(`(?P<a>\d+) (?P<o>[*]) (?P<b>\d+)`)
var reSubExpr = regexp.MustCompile(`\((?P<exp>[\d +*]+)\)`)
var reNoOp = regexp.MustCompile(`\((?P<exp>[\d ]+)\)`)

// var reFooBar = regexp.MustCompile(`[\w+:]{1,}(?P<fooNum>\d+),[\w+\s:]{1,}(?P<barNum>\d+)`)
func reSubMatchMap(r *regexp.Regexp, str string) map[string]string {
	match := r.FindStringSubmatch(str)
	log.Print(match)
	subMatchMap := make(map[string]string)
	if len(match) == 0 {
		return subMatchMap
	}
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

// 1 + 2
// 21012
func evalPiecewise(expr *string, add bool) (result string, changed bool) {
	log.Print("Evaluating: \"", (*expr), "\"")
	for i := 2; i < len(*expr); i++ {
		var a, b int64
		var err error
		var oper rune
		oper = rune((*expr)[i])
		if oper != '+' && oper != '*' {
			continue
		}
		a, err = strconv.ParseInt(string((*expr)[i-2]), 10, 64)
		if err != nil {
			continue
		}
		b, err = strconv.ParseInt(string((*expr)[i+2]), 10, 64)
		if err != nil {
			continue
		}
		if add && oper == '+' {
			changed = true
			// (*expr)[i-2 : i+2] = "     "
			// for j := i - 2; j <= i+2; j++ {
			// 	// (*expr)[j] = ' '
			// }
			// result += (*expr)[:i-2] + strconv.Itoa(int(a+b)) + " " + (*expr)[int(math.Min(float64(i)+3, float64(len(*expr)))):]
			result += (*expr)[:i-2] + strconv.Itoa(int(a+b)) + (*expr)[i+3:]
			return result, changed
		}
	}
	return result, changed
}

// func evalRegex(expr *string, add bool) (result string, changed bool) {

// }

func evalRegexPlus(expr string) (result string) {
	log.Print(expr)
	var changed bool = true
	for changed {
		changed = false
		adds := reSubMatchMap(reAdd, expr)
		if len(adds) > 0 {
			loc := reAdd.FindStringIndex(expr)
			a, _ := strconv.ParseInt(adds["a"], 10, 64)
			b, _ := strconv.ParseInt(adds["b"], 10, 64)
			expr = expr[:loc[0]] + strconv.Itoa(int(a+b)) + expr[loc[1]:]
			changed = true
		}
		log.Print(expr)
		b := reSubMatchMap(reNoOp, expr)
		if len(b) > 0 {
			log.Print("Sad: ", b)
			// We have things like (55)loc := reAdd.FindStringIndex(expr)
			loc := reNoOp.FindStringIndex(expr)
			a, _ := strconv.ParseInt(b["exp"], 10, 64)
			expr = expr[:loc[0]] + strconv.Itoa(int(a)) + expr[loc[1]:]
			changed = true
		}
		log.Print(expr)
	}
	return expr
}

func evalRegexMul(expr string) (result string) {
	return evalRegex(expr, reMul)
}

func evalRegex(expr string, reg *regexp.Regexp) (result string) {
	log.Print(expr)
	var changed bool = true
	for changed {
		changed = false
		adds := reSubMatchMap(reg, expr)
		if len(adds) > 0 {
			loc := reg.FindStringIndex(expr)
			a, _ := strconv.ParseInt(adds["a"], 10, 64)
			b, _ := strconv.ParseInt(adds["b"], 10, 64)
			expr = expr[:loc[0]] + strconv.Itoa(int(a+b)) + expr[loc[1]:]
			changed = true
		}
		log.Print(expr)
		b := reSubMatchMap(reNoOp, expr)
		if len(b) > 0 {
			log.Print("Sad: ", b)
			// We have things like (55)
			loc := reNoOp.FindStringIndex(expr)
			a, _ := strconv.ParseInt(b["exp"], 10, 64)
			expr = expr[:loc[0]] + strconv.Itoa(int(a)) + expr[loc[1]:]
			changed = true
		}
		log.Print(expr)
	}
	return expr
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
	// fileContents := fileToSlice("input")
	// var total int64
	// for _, line := range fileContents {

	// log.Print(evalPlus(line))
	// log.Print("---------")
	// a := eval(line)
	// log.Print(a)
	// total += a

	// }
	// test := "2 + (3 * 3 + 4) + 5 * 2"
	test := "2 + 2 + 3 * (3 + 55 * (2 * 3))"
	log.Print(test)
	test = evalRegexPlus(test)
	log.Print(test)
	// log.Print(reSubMatchMap(reMaths, test))
	// log.Print(reSubMatchMap(reMaths, test))
	// log.Print(reSubMatchMap(reSubExpr, test))
	// reAdd
	// evalPiecewise(test, true)
	// test, _ = evalPiecewise(&test, true)
	// log.Print("\"", test, "\"")
	// var changed bool = true
	// for changed {
	// 	test, changed = evalPiecewise(&test, true)
	// }
	// a, err := strconv.ParseInt("(", 10, 64)
	// if err != nil {
	// 	log.Fatal("Ohfuk")
	// }
	// log.Print(a)
}

func main() {
	load()
}
