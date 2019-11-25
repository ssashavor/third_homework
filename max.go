package main

import (
	"fmt"
	"unicode/utf8"
)

func max(sl []string) string {
	m := 0
	str := ""
	for _, n := range sl {
		if utf8.RuneCountInString(n) > m {
			m = utf8.RuneCountInString(n)
			str = n
		}
	}
	return str
}

func main() {
	fmt.Println(max([]string{"学", "fo"}))
	fmt.Println(max([]string{"one", "two", "three"}))
}
