package main

import (
	"fmt"
	"regexp"
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
