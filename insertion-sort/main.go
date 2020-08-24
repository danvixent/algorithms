package main

func main() {

}

func sort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}

	i, j := 0, 1
	for {
		key := slice[j]
		for i > -1 && slice[i] > key {
			slice[i+1] = slice[i]
		}
		// slice
		j++
	}
}
