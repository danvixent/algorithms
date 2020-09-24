package main

import (
	"fmt"
	"sort"

	fuzz "github.com/google/gofuzz"
)

// merge combines element in slice[p:q] &
// slice[q+1:r] back in slice[p:r] but in
// sorted order
func merge(slice []int64, p, q, r int) {
	// n1 & n2, contain the number of elements
	// slice[p:q] & slice[q+1:r] respectively
	n1 := q - p + 1
	n2 := r - q

	// left & right are temporary slices to hold
	// left => slice[p:q], right => slice[q+1:r]
	// The extra +1 makes room for math.MaxInt64
	// Which is the max integer so all other are
	// always less than it.
	left, right := make([]int64, n1), make([]int64, n2)

	// copy slice[p:q] into left
	for i := 0; i < n1; i++ {
		left[i] = slice[p+i]
	}

	// copy slice[q+1:r] into right
	for i := 0; i < n2; i++ {
		right[i] = slice[q+i+1]
	}

	i, j := 0, 0 // index counters for left & right respectively

	// range from p to r
	// comparing left[i] to right[j],
	// when left[i] is less than or equal
	// to right[j], left[i] will be copied
	// into slice[p] else, right[j] will be copied
	// into slice[p].

	n2-- // set n2 to the last index value of right
	n1-- // set n1 to the last index value of left
	for ; p <= r; p++ {
		if left[i] <= right[j] {
			slice[p] = left[i]
			if i != n1 { // only increment i if n1 hasn't been reached
				i++
			}
		} else {
			slice[p] = right[j]
			if j != n2 { // only increment j if n2 hasn't been reached
				j++
			}
		}
	}
}

// mergeSort sorts slice using the merge sort algo
func mergeSort(slice []int64, p, r int) {
	// As long as p < r keep making recursive calls
	if p < r {
		// compute q such that
		// slice[p:q] contains len(slice)/2 elements
		q := (p + r) / 2
		// call mergeSort on slice[p:q]
		mergeSort(slice, p, q)
		// call mergeSort on slice[q+1:r]
		mergeSort(slice, q+1, r)
		// merge slice[p:q] & slice[q+1:r]
		merge(slice, p, q, r)
	}
}

// int64Slice exists so we can use sort.IsSorted
// to confirm we have sorted correctly
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
	var slice = [2000000]int64{} // 2m elements
	fuzz.New().Fuzz(&slice)
	mergeSort(slice[:], 0, len(slice)-1)

	// wrap as int64Slice type, so we can confirm it is sorted.
	fmt.Println("was sorted correctly:", sort.IsSorted(int64Slice(slice[:])))
}
