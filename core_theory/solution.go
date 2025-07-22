package core_theory

import (
	"fmt"
	"strconv"
)

func Two_pointers_intro() {
	fmt.Println("~ Set: Two Pointers Intro")
	// WHEN? Searching for a pair+ of items based on some REQ

	// Test Input
	input1 := []int{1, 3, 4, 6, 8, 10, 13}

	fmt.Println("f(x) two_sum")
	// Case: invalid array length
	fmt.Println(two_sum([]int{1}, 13))
	// Case: Valid
	fmt.Println(two_sum(input1, 13))
	fmt.Println(two_sum(input1, 4))

	fmt.Println("f(x) two_sum_unsorted")
	// Case: invalid array length
	fmt.Println(two_sum_unsorted([]int{1}, 13))
	// Case: Valid
	fmt.Println(two_sum_unsorted(input1, 13))
	fmt.Println(two_sum_unsorted(input1, 4))
}

func two_sum(input_array []int, target int) []int {
	// Baseline
	if len(input_array) < 2 {
		return []int{-1, -1}
	}

	// #BruteForce
	// for i := 0; i < len(input_array); i++ { // Look at each
	// 	for j := i + 1; j < len(input_array); j++ { // Look at all options
	// 		if input_array[i]+input_array[j] == target { // Compare
	// 			return []int{i, j}
	// 		}
	// 	}
	// }
	// BigO: O(n^2)

	// #Efficient
	left_pointer := 0                     // p* @ start
	right_pointer := len(input_array) - 1 // p* @ end

	for left_pointer < right_pointer { // Termination Point
		sum := input_array[left_pointer] + input_array[right_pointer]

		if sum == target { // Finish...
			return []int{left_pointer, right_pointer}
		}

		// Continue...
		if sum > target {
			right_pointer--
		} else {
			left_pointer++
		}
	}
	// BigO: O(n)

	return []int{-1, -1}
}

func two_sum_unsorted(input_array []int, target int) []int {
	// #Hashmap

	// Baseline
	if len(input_array) < 2 {
		return []int{0, 0}
	}

	seen := make(map[int]int) // val, index

	for i, item := range input_array {
		complement := target - item
		if j, exists := seen[complement]; exists {
			return []int{j, i}
		}

		seen[item] = i
	}
	// BigO: O(n), Space: O(n)

	return []int{0, 0}
}

// TODO
func three_sum() {
	// Pattern: Two Pointer

}

func Sliding_window_intro() {
	fmt.Println("~ Set: Sliding Window intro")

	// WHEN? Expand/Contract window to maintain a valid condition

	// Test Input
	input1 := "abcabcbb"

	fmt.Println("f(x) longest_substring_without_repeat")
	fmt.Println(longest_substring_without_repeat(input1))
}

func longest_substring_without_repeat(input string) int {
	// #Brute Force
	/*
		Looking at each index and moving forward, calculating max where all character is unique
		Use Set to check unique
		Time: O(n^3) because nested loop and checking
		Space: constant
	*/

	maxLen := 0
	fmt.Println(input)

	for i := 0; i < len(input); i++ {
		fmt.Println("- " + string(input[i]))
		seen := make(map[byte]bool)
		for j := i; j < len(input); j++ {
			fmt.Println("looking at j:" + string(input[j]))
			if seen[input[j]] {
				fmt.Println("break")
				break
			}

			seen[input[j]] = true
			if j-i+1 > maxLen {
				maxLen = j - i + 1
				fmt.Println("IMPT: Updating maxLen - " + strconv.Itoa(maxLen))
				// fmt.Printf("IMPT: Updating maxLen - %d \n", maxLen)
			}
		}
	}
	return maxLen
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

func max_subarray_sum() {
	// #Brute - nest loop

	// #Efficient - Kadane's Alg
}

func reverse_linked_list() {

	// #Brute - using extra space

	// #Efficient - inplace

}

func longest_palindrome() {

	// #Brute -

	// #Efficient -
}

func merge_intervals() {

	// #Brute -

	// #Efficient -
}

func remove_dupe_from_string() {
	// #Brute -

	// #Efficient -
}
