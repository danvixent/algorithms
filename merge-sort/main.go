package main

import (
	"fmt"
	"math"
	"sort"

	fuzz "github.com/google/gofuzz"
)

func merge(slice []int64, p, q, r int) {
	n1 := q - p + 1
	n2 := r - q
	left := make([]int64, n1+1)
	right := make([]int64, n2+1)

	for i := 0; i < n1; i++ {
		left[i] = slice[p+i]
	}

	for i := 0; i < n2; i++ {
		right[i] = slice[q+i+1]
	}

	left[n1] = math.MaxInt64
	right[n2] = math.MaxInt64
	i, j := 0, 0
	for ; p <= r; p++ {
		if left[i] <= right[j] {
			slice[p] = left[i]
			i++
		} else {
			slice[p] = right[j]
			j++
		}
	}
}

func mergeSort(slice []int64, p, r int) {
	if p < r {
		q := (p + r) / 2
		mergeSort(slice, p, q)
		mergeSort(slice, q+1, r)
		merge(slice, p, q, r)
	}
}

type int64Slice []int64

func (s int64Slice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s int64Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s int64Slice) Len() int {
	return len(s)
}

func main() {
	var slice = [500000000]int64{} // 500m elements
	fuzz.New().Fuzz(&slice)
	mergeSort(slice[:], 0, len(slice)-1)

	// wrap as int64Slice type, so we can confirm it is sorted.
	fmt.Println(sort.IsSorted(int64Slice(slice[:])))
}
