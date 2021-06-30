package main

import "fmt"

var (
	m, n int
)

func mergeArray(nums1 []int, nums2 []int) []int {
	var res []int
	m = len(nums1)
	n = len(nums2)
	i := 0
	j := 0

	for i < m && j < n {
		if nums1[i] < nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}
	for i < m {
		res = append(res, nums1[i])
		i++
	}
	for j < n {
		res = append(res, nums2[j])
		j++
	}
	return res
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var sum, median float64
	mergedArray := mergeArray(nums1, nums2)
	total := len(mergedArray)
	if total%2 == 0 { //Even number of elements
		i := total / 2
		sum = float64(mergedArray[i-1] + mergedArray[i])
		median = float64(sum / 2)
	} else {
		i := total / 2
		median = float64(mergedArray[i])
	}
	fmt.Printf("%2f\n", median)
	return median
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	m = len(nums1)
	n = len(nums2)
	findMedianSortedArrays(nums1, nums2)
}
