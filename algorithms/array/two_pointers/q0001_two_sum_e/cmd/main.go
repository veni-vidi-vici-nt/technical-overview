package main

import (
	"fmt"
	"technical-overview/algorithms/array/two_pointers/q0001_two_sum_e"
)

// https://leetcode.com/problems/two-sum/description/
func main() {
	fmt.Println("Pattern: two pointers")
	fmt.Println("Question: two sum 0001 e")

	result := q0001_two_sum_e.TwoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println("Result:", result)

}
