package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

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

	depart, err := strconv.ParseInt(fileContents[0], 10, 64)
	if err != nil {
		log.Fatal("Cannot parse thing")
	}
	busIDStrings := strings.Split(fileContents[1], ",")
	busIDs := []int64{}
	for _, i := range busIDStrings {
		if i != "x" {
			j, err := strconv.ParseInt(i, 10, 64)
			if err != nil {
				log.Fatal("Couldn't parse id")
			}
			busIDs = append(busIDs, j)
		}
	}
	log.Print(depart)
	log.Print(busIDs)
	var max float64
	var minID int64
	// min = math.MaxInt64
	for _, id := range busIDs {
		// for k := 1; k*
		// j := depart % i
		k := float64(depart) / float64(id)
		j := k - math.Floor(k)
		if j > max {
			max = j
			minID = id
		}
	}
	a := minID - int64(math.Round(max*float64(minID)))
	log.Print(a)
	log.Print(minID)
	log.Print(minID * a)

	busIDs = []int64{}
	for _, i := range busIDStrings {
		if i != "x" {
			j, err := strconv.ParseInt(i, 10, 64)
			if err != nil {
				log.Fatal("Couldn't parse id")
			}
			busIDs = append(busIDs, j)
		} else {
			busIDs = append(busIDs, 1)
		}
	}

}

func load2() {
	// This approach works, but takes literally forever.
	fileContents := fileToSlice("input")

	busIDStrings := strings.Split(fileContents[1], ",")
	busIDs := []int64{}

	for _, i := range busIDStrings {
		if i != "x" {
			j, err := strconv.ParseInt(i, 10, 64)
			if err != nil {
				log.Fatal("Couldn't parse id")
			}
			busIDs = append(busIDs, j)
		} else {
			busIDs = append(busIDs, 1)
		}
	}
	log.Print(busIDs)
	for j := busIDs[0]; j < math.MaxInt64; j += busIDs[0] {
		for i, b := range busIDs {
			if (j+int64(i))%b != 0 {
				break
			}
			if i == (len(busIDs) - 1) {
				log.Print(j)
				return
			}
		}
	}
}
func woo(start, inc, max int64, busIDs []int64) int64 {
	for j := int64(start); j < math.MaxInt64; j += inc {
		for i, b := range busIDs[:int64(max)] {
			if (j+int64(i))%b != 0 {
				break
			}
			// if i == (len(busIDs) - 1) {
			if int64(i) == max-1 {
				log.Print(j)
				return j
			}
		}
	}
	return 0
}

func load3() {
	// This approach works, but takes literally forever.
	fileContents := fileToSlice("input")

	busIDStrings := strings.Split(fileContents[1], ",")
	busIDs := []int64{}

	for _, i := range busIDStrings {
		if i != "x" {
			j, err := strconv.ParseInt(i, 10, 64)
			if err != nil {
				log.Fatal("Couldn't parse id")
			}
			busIDs = append(busIDs, j)
		} else {
			busIDs = append(busIDs, 1)
		}
	}
	log.Print(busIDs)
	for k := 2; k <= len(busIDs); k++ {
		log.Print("loopan: ", k)
		var inc int64 = 1
		for _, b := range busIDs[:k-1] {
			inc *= b
		}
		log.Print("inc: ", inc)
	out:
		for j := busIDs[0]; j < math.MaxInt64; j += inc {
			for i, b := range busIDs[:k] {
				if (j+int64(i))%b != 0 {
					break
				}
				// if i == (len(busIDs) - 1) {
				if i == k-1 {
					log.Print(j)
					break out
				}
			}
		}
	}
}

func main() {
	load2()
	load3()
}
