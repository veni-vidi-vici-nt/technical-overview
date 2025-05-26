package core_theory

import (
	"fmt"
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
