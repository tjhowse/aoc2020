package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

func sumSlice(slice []int64, start int64, end int64) int64 {
	var sum int64
	for i := start; i < end; i++ {
		sum += slice[i]
	}
	return sum
}

func minInSlice(slice []int64) int64 {
	var min int64
	min = int64(math.Pow(2, 32))
	log.Print(min)
	for _, i := range slice {
		if i < min {
			min = i
		}
	}
	return min
}
func maxInSlice(slice []int64) int64 {
	var max int64
	for _, i := range slice {
		if i > max {
			max = i
		}
	}
	return max
}

func load() {

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var xmas []int64
	for scanner.Scan() {
		i, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal("Couldn't parse offset")
		}
		xmas = append(xmas, i)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	var partA int64

	for i := 25; i < len(xmas); i++ {
		var found bool
		for j := 1; j <= 25; j++ {
			for k := 1; k <= 25; k++ {
				if i == k {
					continue
				}
				if xmas[i-j]+xmas[i-k] == xmas[i] {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			partA = xmas[i]
			log.Print(i)
			log.Print(xmas[i])
			break
		}
	}
	for i := 0; i < len(xmas); i++ {
		for j := i + 1; j < len(xmas); j++ {
			sum := sumSlice(xmas, int64(i), int64(j))
			if sum == partA {
				log.Print(minInSlice(xmas[i:j]) + maxInSlice(xmas[i:j]))
				return
			}
			if sum > partA {
				break
			}
		}
	}
}

func main() {
	load()
}
