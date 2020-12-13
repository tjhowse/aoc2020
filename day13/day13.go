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
	busIdStrings := strings.Split(fileContents[1], ",")
	busIDs := []int64{}
	for _, i := range busIdStrings {
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

}
func main() {
	load()
}
