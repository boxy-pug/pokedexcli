package main

import (
	"fmt"
	"strings"
)

func cleanInput(s string) []string {
	clean := strings.Fields(strings.ToLower(s))
	return clean
}

func main() {

	cln := cleanInput("  hello    wORLD    YEDS sir")
	fmt.Println(cln[1])
}
