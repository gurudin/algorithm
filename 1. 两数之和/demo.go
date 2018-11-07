package main

import (
	"fmt"
	"strconv"
)

func twoSum(nums []int, target int) []int {
	result := []int{0, 0}
	for i := 0; i < len(nums); i++ {
		for j := (i + 1); j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result[0] = i
				result[1] = j

				return result
			}
		}
	}

	return result
}

func twoSumB(nums []int, target int) []int {
	result := []int{0, 0}
	tmp := make(map[string]int)
	for i := 0; i < len(nums); i++ {
		tmp[strconv.Itoa(nums[i])] = i
	}

	for i := 0; i < len(nums); i++ {
		_, ok := tmp[strconv.Itoa(target-nums[i])]
		if ok && tmp[strconv.Itoa(target-nums[i])] != i {
			result[0] = i
			result[1] = tmp[strconv.Itoa(target-nums[i])]

			return result
		}
	}

	return result
}

func twoSumC(nums []int, target int) []int {
	result := []int{0, 0}
	tmp := make(map[string]int)

	for i := 0; i < len(nums); i++ {
		num := string(nums[i])

		if _, ok := tmp[num]; ok {
			result[0] = tmp[num]
			result[1] = i

			return result
		} else {
			tmp[string(target-nums[i])] = i
		}
	}

	return result
}

func main() {
	nums := []int{3, 2, 95, 4, -3}
	target := 92
	fmt.Println(twoSum(nums, target))
	fmt.Println(twoSumB(nums, target))
	fmt.Println(twoSumC(nums, target))
}
