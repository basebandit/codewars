package main

import (
	"fmt"
	"strings"
)

func main() {
	count := vowelCount("abracadabra")
	fmt.Println(count)
}

func vowelCount(word string) (count int) {
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}

	str := strings.ToLower(word)

	for _, s := range str {
		for _, v := range vowels {
			if s == v {
				count++
			}
		}
	}
	return
}
