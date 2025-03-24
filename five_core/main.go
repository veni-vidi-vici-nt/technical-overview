package main

import (
	"fmt"
)

func main() {
	fmt.Println("[Five Core Questions]")
	pattern_two_pointers_intro()

	fmt.Println("[Current Workspace]")
	result := two_sum.TwoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(result) // Expected output: [0, 1]

}

func pattern_two_pointers_intro() {
	fmt.Println("~ Pattern: Two Pointers Intro")

	// Two Sum
	input1 := []int{1, 3, 4, 6, 8, 10, 13}

	println(two_sum(input1, 13))
	println(two_sum(input1, 6))

	// When to use? Searching for a pair+ of items based on REQ

}

func two_sum(input_array []int, target int) bool {
	// Baseline Checks
	if len(input_array) < 2 {
		return false
	}

	// #Brute Force Approach
	// for i := 0; i < len(input_array); i++ { // Look at each
	// 	for j := i + 1; j < len(input_array); j++ { // Look at each @ index +1
	// 		if input_array[i]+input_array[j] == target { // Compare
	// 			fmt.Printf("TRUE: %d %d \n", i, j)
	// 			return true
	// 		}
	// 	}
	// }

	// #Hashmap Approach
	// Note: Unsorted
	// TODO: This returns indexes
	// seen := make(map[int]int) // val, index

	// for i, num := range input_array {
	// 	complement := target - num
	// 	if j, exists := seen[complement]; exists {
	// 		return []int{j, i}
	// 	}
	// 	seen[num] = i
	// }
	// return nil

	// Efficient Two Pointer Approach
	// Constraint: Sorted Input
	left_pointer := 0                     // p* @ start
	right_pointer := len(input_array) - 1 // p* @ end

	for left_pointer < right_pointer { // termination point
		sum := input_array[left_pointer] + input_array[right_pointer]

		// TRUE scenario
		if sum == target {
			return true
		}

		// Continue Onwards
		if sum > target {
			right_pointer--
		} else {
			left_pointer++
		}

	}

	return false
}

func three_sum() {
	// Pattern: Two Pointer

}

func longest_substring_without_repeat() {
	// Pattern: Sliding Window

}

func diameter_of_binary_tree() {
	// Pattern: DFS, Recursion

}

func kth_largest_elements() {
	// Pattern: Heap

}

func number_of_islands() {
	// Pattern: DFS, BFS

}
