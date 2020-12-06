package main

type groupAnswers struct {
	answers     map[rune]int
	peopleCount int
}

func (d *groupAnswers) init() {
	d.answers = make(map[rune]int)
}

func (d *groupAnswers) parse(input string) {
	d.peopleCount++
	for _, a := range input {
		d.answers[a]++
	}
}

func (d *groupAnswers) countPart2() int {
	var result int
	for _, count := range d.answers {
		if count == d.peopleCount {
			result++
		}
	}
	return result
}

func (d *groupAnswers) countPart1() int {
	return len(d.answers)
}
