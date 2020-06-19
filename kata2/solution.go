package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	s1 := time.Now()
	dups := dupCount("Indivisibilities")
	t1 := time.Since(s1)
	fmt.Println(dups)
	fmt.Printf("dupCount1 took %s", t1)
	s2 := time.Now()
	dup2 := dupCount2("Indivisibilities")
	t2 := time.Since(s2)
	fmt.Println(dup2)
	fmt.Printf("dupCount2 took %s", t2)

}

//my solution.
func dupCount(word string) int {
	str := strings.ToLower(word)

	counts := make(map[string]int)

	for _, s := range str {
		counts[string(s)]++
	}

	var count int

	for _, v := range counts {
		if v > 1 { //appears more than once
			count++
		}
	}

	return count
}

// slightly modified version of this solution
//https://www.codewars.com/kata/reviews/59afce199ba64cdd16000ce2/groups/59c1a84d67f80c128b00094e
func dupCount2(word string) (c int) {
	h := map[rune]int{}
	str := strings.ToLower(word)
	for _, r := range str {
		if h[r]++; h[r] == 2 {
			c++
		}
	}
	return
}
