package main

import (
	"fmt"
	"strings"
)


func countString(s string) map[string]int {

	var count = make(map[string]int)

	for _, c := range(s) {
		c_str := string(c)
		x, ok := count[c_str]
		if ok {
			count[c_str] = x + 1
		} else {
			count[c_str] = 1
		}
	}

	return count

}

func swapCase(s string) string {

	upper := strings.ToUpper(s)

	if s == upper {
		return strings.ToLower(s)
	} else {
		return upper
	}

}

func consume(msg map[string]int, news map[string]int) (int, int) {

	var right int = 0
	var wrong int = 0

	// consume right case
	for letter, _ := range(msg) {
		_, ok := news[letter]
		if ok {
			if news[letter] >= msg[letter] {
				right += msg[letter]
				news[letter] -= msg[letter]
				msg[letter] = 0
			} else {
				right += news[letter]
				msg[letter] -= news[letter]
				news[letter] = 0
			}
		}
	}

	// consume the swapped case if needed
	for letter, _ := range(msg) {
		if msg[letter] != 0 {
			swapped := swapCase(letter)
			_, ok := news[swapped]
			if ok {
				if news[swapped] >= msg[letter] {
					wrong += msg[letter]
					news[swapped] -= msg[letter]
					msg[letter] = 0
				} else {
					wrong += news[swapped]
					msg[letter] -= news[swapped]
					news[swapped] = 0
				}
			}
		}
	}

	return right, wrong

}


func main() {

	msg := "AaAaAa"
	news := "Aaaaaaaaaa"

	countMsg := countString(msg)
	countNews := countString(news)

	//fmt.Printf("%v\n", countMsg)
	//fmt.Printf("%v\n", countNews)

	right, wrong := consume(countMsg, countNews)

	//fmt.Printf("%v\n", countMsg)
	//fmt.Printf("%v\n", countNews)
	fmt.Printf("%d %d\n", right, wrong)

}
