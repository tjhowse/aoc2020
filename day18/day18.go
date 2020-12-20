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
var reMulAdd = regexp.MustCompile(`[\d\+* ]+`)
var reSubExpr = regexp.MustCompile(`\((?P<exp>[\d +*]+)\)`)
var reNoOp = regexp.MustCompile(`\((?P<exp>[\d ]+)\)`)

// var reFooBar = regexp.MustCompile(`[\w+:]{1,}(?P<fooNum>\d+),[\w+\s:]{1,}(?P<barNum>\d+)`)
func reSubMatchMap(r *regexp.Regexp, str string) map[string]string {
	match := r.FindStringSubmatch(str)
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

type operation func(int64, int64) int64

func add(a, b int64) int64 {
	return a + b
}
func mul(a, b int64) int64 {
	return a * b
}

func evalRegexPlus(expr string) (result string, changed bool) {
	orig := expr
	result = evalRegex(expr, reAdd, add, "add")
	if orig != result {
		changed = true
	}
	return result, changed
}

func evalRegexMul(expr string) (result string, changed bool) {
	orig := expr
	// // Find any instances of mixtures of + and * at the same level and strip them
	// loc := reMulAdd.FindStringIndex(expr)
	result = evalRegex(expr, reMul, mul, "mul")
	if orig != result {
		changed = true
	}
	return result, changed
}

func evalRegex(expr string, reg *regexp.Regexp, op operation, t string) (result string) {
	var changed bool = true
	for changed {
		changed = false
		adds := reSubMatchMap(reg, expr)
		if len(adds) > 0 {
			loc := reg.FindStringIndex(expr)
			if t == "mul" {
				// We need to prevent this operating on expressions that contain + and * at the
				// same level
			}
			a, _ := strconv.ParseInt(adds["a"], 10, 64)
			b, _ := strconv.ParseInt(adds["b"], 10, 64)
			expr = expr[:loc[0]] + strconv.Itoa(int(op(a, b))) + expr[loc[1]:]
			changed = true
		}
		b := reSubMatchMap(reNoOp, expr)
		if len(b) > 0 {
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

func findEndOfSubExpression(expr string, start int64) (end int64) {
	var depth int64 = 0
	for j := int64(start); j < int64(len(expr)); j++ {
		if expr[j] == '(' {
			depth++
		} else if expr[j] == ')' {
			depth--
		}
		if depth == 0 {
			return j
		}
	}
	return start
}

func findStartOfSubExpression(expr string, end int64) (start int64) {
	var depth int64 = 0
	for j := int64(end); j >= 0; j-- {
		if expr[j] == ')' {
			depth++
		} else if expr[j] == '(' {
			depth--
		}
		if depth == 0 {
			return j
		}
	}
	return end
}

func addBracketsToPlus(expr string) (result string) {
	// findEndOfSubExpression(expr, start)
	// for i, c := range expr {
	for i := 0; i < len(expr); i++ {
		c := rune(expr[i])
		if c == '+' {
			start := findStartOfSubExpression(expr, int64(i-2))
			end := findEndOfSubExpression(expr, int64(i+2)) + 1
			// log.Print("Wrapping \"", expr[start:end], "\" in brackets")
			expr = expr[:start] + "(" + expr[start:end] + ")" + expr[end:]
			i += 2
			// fmt.Println(expr)
			// for k := 0; k < i; k++ {
			// 	fmt.Print(" ")
			// }
			// fmt.Println("^")
			// log.Print("i: ", i, " expr[i]: ", string(expr[i]))
			// log.Print(result)

		}
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
	// for _, line := range fileContents {
	// 	for {
	// 		var changedPlus, changedMul bool
	// 		line, changedPlus = evalRegexPlus(line)
	// 		line, changedMul = evalRegexMul(line)
	// 		// if line == "46" {
	// 		// 	break
	// 		// }
	// 		if !changedPlus && !changedMul {
	// 			break
	// 			// log.Print()
	// 		}
	// 	}
	// 	log.Print(line)
	// }
	// test := "1 + 2 + 3"
	// test = addBracketsToPlus(test)
	// log.Print(test)

	// log.Print(findEndOfSubExpression(test, 0))

	// log.Print(evalPlus(line))
	// log.Print("---------")

	fileContents := fileToSlice("input")
	var total int64
	for _, line := range fileContents {
		// log.Print(line)
		line = addBracketsToPlus(line)
		a := eval(line)
		// log.Print(line)
		// log.Print(a)
		total += a
	}
	log.Print(total)

	// test := "2 + (3 * 3 + 4) + 5 * 2"
	// test := "2 + 2 + 3 * (3 + 55 * (2 * 3))"
	// log.Print(test)
	// test = evalRegexPlus(test)
	// test = evalRegexMul(test)
	// log.Print(test)
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
