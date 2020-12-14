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
				return j
			}
		}
	}
	return 0
}

// This solution thanks to cinphart. I reached many dead ends.
func load3() {
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
	// Step through the list of IDs gradually. Do:
	// 0 -> 1, get an answer, then
	// 0 -> 2, get an answer... etc
	var prevAnswer int64
	for k := 2; k <= len(busIDs); k++ {
		var inc int64 = 1
		// Incrementing from the starting point using the
		// product of all prior IDs. This is because
		// a valid departure time is v + n*p, where v
		// is the -first- valid departure time, and n is
		// any integer and p is the product of all bus IDs
		// for which that is a valid departure time.
		// I don't fully understand why yet.
		for _, b := range busIDs[:k-1] {
			inc *= b
		}
		// The starting point for the search for the next valid answer
		// for this 0 -> k range is the previous answer.
		prevAnswer = woo(prevAnswer, inc, int64(k), busIDs)
	}
	log.Print("prevAnswer: ", prevAnswer)
}

func main() {
	// load2()
	load3()
}
