package main

import "fmt"

func reverse(sl []int64) []int64 {

	lenSl := len(sl)
	reversedSlice := make([]int64, 0)
	for i := lenSl - 1; i >= 0; i-- {
		reversedSlice = append(reversedSlice, sl[i])
	}
	return reversedSlice

}

func main() {
	slice := []int64{1, 2, 5, 15}
	fmt.Println(reverse(slice))
	fmt.Println(slice)
}
