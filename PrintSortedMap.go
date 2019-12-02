package main

import (
	"fmt"
	"sort"
)

func printsorted(m map[int]string) {
	sliceOfKeys := make([]int, 0)
	for k := range m {
		sliceOfKeys = append(sliceOfKeys, k)
	}
	sort.Ints(sliceOfKeys)
	for i := range sliceOfKeys {
		 fmt.Println(m[sliceOfKeys[i]])
	}
}

func main() {
	printsorted(map[int]string{2: "a", 0: "b", 1: "c"})
	printsorted(map[int]string{10: "aa", 0: "bb", 500: "cc"})
}
