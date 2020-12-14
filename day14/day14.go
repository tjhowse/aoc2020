package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type a struct {
	a int
}

const fieldLen = 36

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
	fileContents := fileToSlice("input")

	// depart, err := strconv.ParseInt(fileContents[0], 10, 64)
	// if err != nil {
	// 	log.Fatal("Cannot parse thing")
	// }

	r := make(map[int64]int64)
	var mask string
	var addr int64
	var val int64
	var err error
	for _, line := range fileContents {
		command := strings.Split(line, " ")
		if command[0] == "mask" {
			mask = command[len(command)-1]
		} else {
			addr, err = strconv.ParseInt(command[0][4:len(command[0])-1], 10, 64)
			if err != nil {
				log.Fatal("Poo")
			}
			val, err = strconv.ParseInt(command[len(command)-1], 10, 64)
			if err != nil {
				log.Fatal("Poo")
			}
			// log.Print("Val: ", val)
			// fmt.Println(strconv.FormatInt(val, 2))
			for i, c := range mask {
				switch c {
				case rune('X'):
				case rune('1'):
					// log.Print("Setting bit ", (fieldLen - i - 1))
					val = val | (1 << (fieldLen - i - 1))
				case rune('0'):
					// log.Print("Clearing bit ", (fieldLen - i - 1))
					val = val &^ (1 << (fieldLen - i - 1))
				}
			}
			// log.Print("Val: ", val)
			// fmt.Println(strconv.FormatInt(val, 2))
			r[addr] = val
		}

		// log.Print(mask)
		// log.Print(addr)
		// log.Print(val)
		// log.Print(r)

	}
	sum := 0
	for _, i := range r {
		sum += int(i)
	}
	log.Print(sum)
}

func load2() {
	fileContents := fileToSlice("input")

	// depart, err := strconv.ParseInt(fileContents[0], 10, 64)
	// if err != nil {
	// 	log.Fatal("Cannot parse thing")
	// }

	r := make(map[int64]int64)
	var mask string
	var addr int64
	var val int64
	var err error
	for _, line := range fileContents {
		command := strings.Split(line, " ")
		if command[0] == "mask" {
			mask = command[len(command)-1]
		} else {
			addr, err = strconv.ParseInt(command[0][4:len(command[0])-1], 10, 64)
			if err != nil {
				log.Fatal("Poo")
			}
			val, err = strconv.ParseInt(command[len(command)-1], 10, 64)
			if err != nil {
				log.Fatal("Poo")
			}
			// log.Print("Val: ", val)
			// fmt.Println(strconv.FormatInt(addrMask, 2))
			addrs := map[int64]bool{}
			// addrs := []int64{}
			// addrMask := int64(0)
			for i, c := range mask {
				switch c {
				case '0':
				case '1':
					addr = addr | (1 << (fieldLen - i - 1))
				}
			}
			// addrs = append(addrs, addr)
			addrs[addr] = true
			for i, c := range mask {
				if c == 'X' {
					for addr := range addrs {
						a1 := addr | (1 << (fieldLen - i - 1))
						a2 := addr &^ (1 << (fieldLen - i - 1))
						// fmt.Println(strconv.FormatInt(a1, 2))
						// fmt.Println(strconv.FormatInt(a2, 2))
						addrs[a1] = true
						addrs[a2] = true
					}
				}
			}
			// fmt.Println("Mask: ", strconv.FormatInt(addrMask, 2))
			// log.Print(addrs)
			for addr := range addrs {
				r[addr] = val
			}
		}

		// log.Print(mask)
		// log.Print(addr)
		// log.Print(val)
		// log.Print(r)

	}
	sum := 0
	for _, i := range r {
		sum += int(i)
	}
	log.Print(sum)
}
func main() {
	// load()
	load2()
}
