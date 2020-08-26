package main

import "fmt"

func main() {
	fmt.Println(sort([]int{5, 4, 3, 6, 2, 1, -2, -8, 1002, 99, 22, 32, 435, 5546, 24311, 4442, 53932, 4245, 242421}))
}

func sort(slice []int) []int {
	for j := 1; j < len(slice); j++ {
		i := j - 1
		key := slice[j]
		for i > -1 && slice[i] > key {
			slice[i+1] = slice[i]
			i--
		}
		slice[i+1] = key
	}
	return slice
}
