package main

import (
	"log"
	"strings"
)

type passportLax struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
	raw string
}

func (d *passportLax) parse(input string) bool {
	for _, kv := range strings.Split(input, " ") {
		p := strings.Split(kv, ":")
		var err error
		switch p[0] {
		case "byr":
			d.byr = p[1]
		case "iyr":
			d.iyr = p[1]
		case "eyr":
			d.eyr = p[1]
		case "hgt":
			d.hgt = p[1]
		case "hcl":
			d.hcl = p[1]
		case "ecl":
			d.ecl = p[1]
		case "pid":
			d.pid = p[1]
		case "cid":
			d.cid = p[1]
		}
		if err != nil {
			log.Print("Failed to parse field from entry")
			log.Fatal(p)
		}
	}
	d.raw = input
	return true
}

func (d *passportLax) validate() bool {
	if d.byr == "" {
		return false
	}
	if d.iyr == "" {
		return false
	}
	if d.eyr == "" {
		return false
	}
	if d.hgt == "" {
		return false
	}
	if d.hcl == "" {
		return false
	}
	if d.ecl == "" {
		return false
	}
	if d.pid == "" {
		return false
	}
	// if d.cid == 0 {
	// 	return false
	// }
	return true
}
