package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type ins struct {
	ins    string
	offset int64
}

func load() {

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var program []ins
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		newIns := ins{}
		newIns.ins = s[0]
		newIns.offset, err = strconv.ParseInt(s[1], 10, 64)
		if err != nil {
			log.Fatal("Couldn't parse offset")
		}
		program = append(program, newIns)
	}
	// log.Print(program)

	var acc int64
	runLog := make(map[int]bool)

	for i := 0; i < len(program); {
		if runLog[i] {
			break
		}
		runLog[i] = true
		switch program[i].ins {
		case "nop":
			// log.Print("NOP")
			i++
		case "acc":
			acc += program[i].offset
			i++
		case "jmp":
			i += int(program[i].offset)
		}
	}
	log.Print("Part 1")
	log.Print(acc)

	max := int64(10000)

	for j := 0; j < len(program); j++ {
		runLog := make(map[int]bool)
		var count int64
		acc = 0
		broke := false
		for i := 0; i < len(program) && count < max; {
			if runLog[i] {
				broke = true
				break
			}
			runLog[i] = true
			ins := program[i].ins
			if i == j {
				if ins == "nop" {
					ins = "jmp"
				} else if ins == "jmp" {
					ins = "nop"
				}
			}
			switch ins {
			case "nop":
				// log.Print("NOP")
				i++
			case "acc":
				acc += program[i].offset
				i++
			case "jmp":
				i += int(program[i].offset)
			}
		}
		if !broke {
			log.Print(acc)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	load()
}
