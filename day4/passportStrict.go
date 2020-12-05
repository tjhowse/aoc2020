package main

import (
	"log"
	"strconv"
	"strings"
)

func validateEyeColour(c string) bool {
	switch c {
	case "amb":
		return true
	case "blu":
		return true
	case "brn":
		return true
	case "gry":
		return true
	case "grn":
		return true
	case "hzl":
		return true
	case "oth":
		return true
	}
	return false
}
func validateHairColour(c string) bool {
	if len(c) != 7 || c[0] != '#' {
		return false
	}
	_, err := strconv.ParseInt(c[1:len(c)], 16, 64)
	if err != nil {
		return false
	}
	return true
}

type passportStrict struct {
	byr int64
	iyr int64
	eyr int64
	hgt string
	hcl string
	ecl string
	pid string
	cid int64
	raw string
}

func (d *passportStrict) parse(input string) bool {
	for _, kv := range strings.Split(input, " ") {
		p := strings.Split(kv, ":")
		var err error
		switch p[0] {
		case "byr":
			d.byr, err = strconv.ParseInt(p[1], 10, 64)
		case "iyr":
			d.iyr, err = strconv.ParseInt(p[1], 10, 64)
		case "eyr":
			d.eyr, err = strconv.ParseInt(p[1], 10, 64)
		case "hgt":
			d.hgt = p[1]
		case "hcl":
			d.hcl = p[1]
		case "ecl":
			d.ecl = p[1]
		case "pid":
			d.pid = p[1]
		case "cid":
			d.cid, err = strconv.ParseInt(p[1], 10, 64)
		}
		if err != nil {
			log.Print("Failed to parse field from entry")
			log.Fatal(p)
		}
	}
	d.raw = input
	return true
}

func (d *passportStrict) validate() bool {
	if d.byr < 1920 || d.byr > 2002 {
		return false
	}
	if d.iyr < 2010 || d.iyr > 2020 {
		return false
	}
	if d.eyr < 2020 || d.eyr > 2030 {
		return false
	}
	if len(d.hgt) < 4 {
		return false
	}
	if d.hgt[len(d.hgt)-2:] == "in" {
		h, err := strconv.ParseInt(d.hgt[:2], 10, 64)
		// log.Print(h)
		if h < 59 || h > 76 || err != nil {
			return false
		}
	} else {
		h, err := strconv.ParseInt(d.hgt[:3], 10, 64)
		if h < 150 || h > 193 || err != nil {
			return false
		}
	}
	if !validateHairColour(d.hcl) {
		return false
	}
	if !validateEyeColour(d.ecl) {
		return false
	}
	if len(d.pid) != 9 {
		return false
	}
	// if d.cid == 0 {
	// 	return false
	// }
	return true
}
