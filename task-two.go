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

func generateRandomString(param bool) string {
	if !param {
		return "abc"
	}

	pattern := "^[0-9+]-[a-zA-Z]+(?:-[0-9]+-[a-zA-Z]+)*$"
	str, err := reggen.Generate(pattern, 10)

	if err != nil {
		fmt.Println(err)
	}
	return str

}
