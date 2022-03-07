package main

import (
	"fmt"

	"github.com/lucasjones/reggen"
)

func main() {
	var param string

	fmt.Print("Enter flag (true or false): ")
	fmt.Scanf("%s", &param)

	if param == "true" || param == "false" {
		fmt.Println(generateRandomString(param == "true"))
	}
}

// DIFFICULTY EASY
// ESTIMATED 1 hr
// COMPLETED 0.5 hr
func generateRandomString(param bool) string {
	if !param {
		pattern := "^[0-9+][a-zA-Z]+(?:-[0-9]+-[a-zA-Z]+)*$"
		str, _ := reggen.Generate(pattern, 10)
		return str
	}

	pattern := "^[0-9+]-[a-zA-Z]+(?:-[0-9]+-[a-zA-Z]+)*$"
	str, err := reggen.Generate(pattern, 10)

	if err != nil {
		fmt.Println(err)
	}
	return str

}
