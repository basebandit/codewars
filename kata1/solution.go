package main

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

func main() {
	s1 := time.Now()
	str := toWeirdCase("Weird string case")
	t1 := time.Since(s1)
	fmt.Println(str)
	fmt.Printf("toWeirdCase1 took %s\n", t1)

	s2 := time.Now()
	str = toWeirdCase2("Weird string case")
	t2 := time.Since(s2)
	fmt.Println(str)
	fmt.Printf("toWeirdCase2 took %s\n", t2)
}

func toWeirdCase(s string) string {
	var buf strings.Builder
	if strings.Contains(s, " ") {
		words := strings.Fields(s)   //substring containing words without spaces e.g ["foo" "bar" "baz"]
		for _, word := range words { //iterate over each word
			for i, v := range word { //iterate over each rune in each word
				if i%2 == 0 { //uppercase even indexed runes
					buf.WriteRune(unicode.ToUpper(v))
				} else {
					buf.WriteRune(unicode.ToLower(v))
				}
			}

			buf.WriteString(" ") //Append back the space after each word. Note this appends a trailing whitespace which we remove at the end before returning the new string.

		}
	} else { //for one word strings.
		for i, v := range s {
			if i%2 == 0 { //uppercase even indexed runes
				buf.WriteRune(unicode.ToUpper(v))
			} else { //lowercase odd indexed runes
				buf.WriteRune(unicode.ToLower(v))
			}
		}
	}

	return strings.TrimSpace(buf.String()) //To remove the trailing we added above.
}

func toWeirdCase2(s string) string {
	var buf strings.Builder
	if strings.Contains(s, " ") {
		words := strings.Fields(s)     //substring containing words without spaces e.g ["foo" "bar" "baz"]
		for idx, word := range words { //iterate over each word

			if idx > 0 {
				buf.WriteString(" ") //Append space after each word
			}

			for i, v := range word { //iterate over each rune in each word
				if i%2 == 0 { //uppercase even indexed runes
					buf.WriteRune(unicode.ToUpper(v))
				} else { //lowercase odd indexed runes
					buf.WriteRune(unicode.ToLower(v))
				}
			}
		}
	} else { //one word strings
		for i, v := range s {
			if i%2 == 0 { //uppercase even indexed runes
				buf.WriteRune(unicode.ToUpper(v))
			} else { //lowercase odd indexed runes
				buf.WriteRune(unicode.ToLower(v))
			}
		}
	}

	return buf.String()
}
