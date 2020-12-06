package main

type groupAnswers struct {
	answers map[rune]bool
}

func (d *groupAnswers) init() {
	d.answers = make(map[rune]bool)
}

func (d *groupAnswers) parse(input string) {
	for _, a := range input {
		d.answers[a] = true
	}
}

func (d *groupAnswers) count() int {
	return len(d.answers)
}
