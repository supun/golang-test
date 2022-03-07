package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

func main() {
	var param string

	fmt.Print("Enter string: ")
	fmt.Scanf("%s", &param)

	if !TestValidity(param) {
		fmt.Println("The entered string is not in the correct format")
		return
	}

	fmt.Println("Average number ", AverageNumber(param))
	fmt.Println("Whole sentence ", WholeStory(param))
	fmt.Println(StoryStat(param))

}

// DIFFICULTY MEDIUM
// ESTIMATED 2 hr
// COMPLETED 2.5 hr
func TestValidity(str string) bool {
	if len(str) == 0 {
		return false
	}
	match, _ := regexp.MatchString("^[0-9+]+-[a-zA-Z]+(?:-[0-9]+-[a-zA-Z]+)*$", str)
	return match
}

// DIFFICULTY EASY
// ESTIMATED 1 hr
// COMPLETED 0.5 hr
func AverageNumber(str string) float32 {
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
func WholeStory(str string) string {
	// REPLACE NON ALPHABETICAL WITH SPACE
	m := regexp.MustCompile("\\P{L}")
	temp := m.ReplaceAllString(str, " ")

	// REMOVE MULTIPLE SPACES, KEEP ONLY SINGLE SPACE BETWEEN TWO WORDS
	m1 := regexp.MustCompile("[\\s\\p{Zs}]{2,}")
	temp1 := m1.ReplaceAllString(temp, " ")

	// REMOVE LEADING AND TRAILING SPACES
	m2 := regexp.MustCompile("^[\\s\\p{Zs}]+|[\\s\\p{Zs}]+$")
	res := m2.ReplaceAllString(temp1, "")
	return res
}

// DIFFICULTY EASY
// ESTIMATED 1 hr
// COMPLETED 0.5 hr
func StoryStat(str string) (shortest, longest, averagelength string, list []string) {
	// CONVERT STRING TO SENTENCE WITH SPACES
	sentence := WholeStory(str)

	// CONVERT SENTENCE TO STRING ARRAY
	s := strings.Fields(sentence)
	shortest = s[0]
	longest = s[0]
	count := decimal.NewFromInt(int64(len(s)))
	tempstr := strings.ReplaceAll(sentence, " ", "")
	length := decimal.NewFromInt(int64(len(tempstr)))

	average := length.Div(count)
	var listtemp = make([]string, 0)

	// LOOP THROUGH THE STRING TO GET THE SHORTEST AND LONGEST WORDS
	for _, val := range s {
		if len(val) > len(longest) {
			longest = val
		}

		if len(val) < len(shortest) {
			shortest = val
		}

		// USED THIRD PARTY PACKAGE TO HANDLE DECIMAL VALUES WHICH CAN BE USED TO ROUND UP AND ROUND DOWN FUNCTIONS
		if decimal.NewFromInt(int64(len(val))).Equal(average.RoundUp(0)) || decimal.NewFromInt(int64(len(val))).Equal(average.RoundDown(0)) {
			listtemp = append(listtemp, val)
		}
	}

	return shortest, longest, average.String(), listtemp
}
