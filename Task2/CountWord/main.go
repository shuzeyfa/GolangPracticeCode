package main

import (
	"fmt"
	"strings"
	"unicode"
)

func removingPunctuationMark(l []string) []string {

	container := []string{}

	for _, value := range l {

		temp := []string{}
		for _, char := range value {
			if unicode.IsPunct(char) {
				continue
			}

			temp = append(temp, string(char))
		}
		container = append(container, strings.Join(temp, ""))
	}

	fmt.Println(container)
	return container

}

func CountWord(s string) map[string]int {

	lowerCaseString := strings.ToLower(s)

	sliced := strings.Fields(lowerCaseString)

	filteredString := removingPunctuationMark(sliced)

	countedString := map[string]int{}

	for _, val := range filteredString {
		countedString[string(val)] += 1
	}

	return countedString

}

func main() {

	test1 := "This One Is The First Test Case"
	test2 := "This One Is The Second Test Case!"
	test3 := "This one is the final test case, right? "

	Count := CountWord(test1)
	count2 := CountWord(test2)
	count3 := CountWord(test3)

	for key, val := range Count {
		fmt.Println(key, val)
	}
	println()

	for key, val := range count2 {
		fmt.Println(key, val)
	}
	println()

	for key, val := range count3 {
		fmt.Println(key, val)
	}
}
