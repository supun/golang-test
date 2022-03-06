package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {

}

// DIFFICULTY MEDIUM
// ESTIMATED 2 hr
// COMPLETED 2.5 hr
func testValidity(str string) bool {
	match, _ := regexp.MatchString("^[0-9+]+-[a-zA-Z]+$", str)
	fmt.Println(match)
	return match
}

// DIFFICULTY EASY
// ESTIMATED 1 hr
// COMPLETED 0.5 hr
func averageNumber(str string) float32 {
	var count, total int
	s := strings.Split(str, "-")
	for _, val := range s {
		if valnum, err := strconv.Atoi(val); err == nil {
			count++
			total += valnum
		}
	}

	// THIS IS TO GET RID FROM DIVISION BY ZERO
	if count == 0 {
		return 0
	}
	return float32(total / count)
}

// DIFFICULTY EASY
// ESTIMATED 1 hr
// COMPLETED 0.5 hr
func wholeStory(str string) string {
	// REPLACE NON ALPHABETICAL WITH SPACE
	m := regexp.MustCompile("\\P{L}")
	temp := m.ReplaceAllString(str, " ")

	// REMOVE MULTIPLE SPACES, KEEP ONLY SINGLE SPACE BETWEEN TWO WORDS
	m1 := regexp.MustCompile(" +")
	res := m1.ReplaceAllString(temp, " ")
	//fmt.Println(strings.TrimSpace(res))
	return res
}
