package main

import "fmt"

func average(sl []int) float64 {
	if len(sl) == 0 {
		return 0
	}
	sum := 0
	for _, n := range sl {
		sum += n
	}
	return float64(sum) / float64(len(sl))
}

func main() {
	fmt.Println(average([]int{1, 2, 3, 4, 5, 6}))
}
