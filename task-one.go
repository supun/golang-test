package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
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

// DIFFICULTY EASY
// ESTIMATED 1 hr
// COMPLETED 0.5 hr
func storyStat(str string) (shortest, longest, averagelength string, list []string) {
	// CONVERT STRING TO SENTENCE WITH SPACES
	sentence := wholeStory(str)

	s := strings.Fields(sentence)
	shortest = s[0]
	longest = s[0]
	count := decimal.NewFromInt(int64(len(s)))
	tempstr := strings.ReplaceAll(sentence, " ", "")
	length := decimal.NewFromInt(int64(len(tempstr)))

	average := length.Div(count)
	var listtemp = make([]string, 0)
	for _, val := range s {
		if len(val) > len(longest) {
			longest = val
		}

		if len(val) < len(shortest) {
			shortest = val
		}

		if decimal.NewFromInt(int64(len(val))).Equal(average.RoundUp(0)) || decimal.NewFromInt(int64(len(val))).Equal(average.RoundDown(0)) {
			listtemp = append(listtemp, val)
		}
	}

	return shortest, longest, average.String(), listtemp
}
