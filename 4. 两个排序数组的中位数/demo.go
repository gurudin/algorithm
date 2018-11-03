package main

import (
	"fmt"
	"sort"
)

type intSlice []int

func (s intSlice) Len() int           { return len(s) }
func (s intSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s intSlice) Less(i, j int) bool { return s[i] < s[j] }

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var result float64
	nums := make([]int, len(nums1)+len(nums2))
	copy(nums, nums1)
	copy(nums[len(nums1):], nums2)
	sort.Sort(intSlice(nums))

	length := len(nums)
	if length%2 == 1 {
		result = float64(nums[length/2])
	} else {
		result = (float64(nums[length/2]) + float64(nums[length/2-1])) / 2
	}

	return result
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	fmt.Println(findMedianSortedArrays([]int{1, 3, 5, 100}, []int{2}))
}
