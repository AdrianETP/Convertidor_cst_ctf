package main

import (
	"os"
	"strings"
)

func main() {
	entrada := os.Args[1]
	var flag string

	changed_text := strings.ToLower(entrada)
	changed_text = strings.ReplaceAll(changed_text, "a", "4")
	changed_text = strings.ReplaceAll(changed_text, "e", "3")
	changed_text = strings.ReplaceAll(changed_text, "i", "1")
	changed_text = strings.ReplaceAll(changed_text, "o", "0")
	changed_text = strings.ReplaceAll(changed_text, " ", "_")
	changed_text = strings.ReplaceAll(changed_text, "\n", "_")
	flag = "CST{" + changed_text + "}"
	print(flag)
}
