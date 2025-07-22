package q0001_two_sum_e

// TwoSum returns the indices of the two numbers that add up to the target.
func TwoSum(nums []int, target int) []int {
	// PreDiscussion
	/*
		What should we do if we cannot find an appropriate result?
	*/

	// Brute Solution
	/*
		Compare each element with all the others
		Double for loop; until match
		BigO - N^2
		Space - n/a
	*/

	// Efficient Solution
	/*
		Use a hashmap to store complements and its index
		Big0 - N
		Space - N
	*/
	seen := make(map[int]int) // Store number → index

	for i, num := range nums {
		complement := target - num
		if j, exists := seen[complement]; exists {
			return []int{j, i}
		}
		seen[num] = i
	}
	return nil
}
