package main

import (
	"fmt"
	"strings"
	"unicode"
)

func getFilteredString(s string) string {

	stringList := []string{}

	for _, val := range s {
		if val == ' ' || unicode.IsPunct(val) {
			continue
		}
		stringList = append(stringList, strings.ToLower(string(val)))
	}

	return strings.Join(stringList, "")
}

func check(s string) bool {

	index := len(s) - 1

	for i := 0; i < index; i++ {
		if s[index] != s[i] {
			return false
		}
		index -= 1
	}

	return true
}

func main() {

	test1 := "I am I"
	test2 := "abcdcba ? "
	test3 := ""
	test4 := "A"
	test5 := "ABCCBA"
	test6 := "ABCDEF"

	first := getFilteredString(test1)
	second := getFilteredString(test2)
	third := getFilteredString(test3)
	fourth := getFilteredString(test4)
	fifth := getFilteredString(test5)
	sixth := getFilteredString(test6)

	fmt.Println(check(first))
	fmt.Println(check(second))
	fmt.Println(check(third))
	fmt.Println(check(fourth))
	fmt.Println(check(fifth))
	fmt.Println(check(sixth))

}
