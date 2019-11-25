package main

import (
	"fmt"
	"unicode/utf8"
)

func max(sl []string) string {
	str := ""
	for _, n := range sl {
		if utf8.RuneCountInString(n) > utf8.RuneCountInString(str) { 
			str = n
		}
	}
	return str
}

func main() {
	fmt.Println(max([]string{"学", "fo"}))
	fmt.Println(max([]string{"one", "two", "three"}))
}
