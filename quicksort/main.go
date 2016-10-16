package main

import "fmt"

// Swap values in a slice
func swap(s []int, i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Partition is the meat of ths sorting. Find a midpoint and arrange values relative to it
func partition(s []int, begin, end int) int {
	pivotVal := s[begin]
	left := begin
	right := end

	for left < right {
		// Deal with duplicate values. Simply move ahead left (or right)
		if s[left] == s[right] {
			left++
		}
		for s[left] < pivotVal {
			left++
		}

		for s[right] > pivotVal {
			right--
		}
		swap(s, left, right)
	}
	return right
}

// Sort find the split point and recursively calls itself over again until the whole thing is sorted
func sort(s []int, begin, end int) {
	if begin < end {
		splitPoint := partition(s, begin, end)
		sort(s, begin, splitPoint-1)
		sort(s, splitPoint+1, end)
	}
}

func main() {
	//unorderedList := []int{901, 73, 73, 278, 24, 12, 35, 23, 235, 25, 134, 123, 52, 634, 6723, 2, 82, 937, 903, 17, 10}
	unorderedList := []int{1, 2, 2, 5, 2, 4}
	fmt.Printf("Sort the following: \n%v\n", unorderedList)

	sort(unorderedList, 0, len(unorderedList)-1)

	fmt.Printf("%v\n", unorderedList)
}
