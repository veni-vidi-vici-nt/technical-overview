package q0001_two_sum_e

// TwoSum returns the indices of the two numbers that add up to the target.
func TwoSum(nums []int, target int) []int {
	/*
		# RE (Read)
		Thoughts:
			Simple problem

		Scope:
			Will our input array ever be empty?
			Will we be dealing with +/- values?
			What output is preferred when there is no result?
	*/

	/*
		# P (Plan)
		Brute:
			"Double For Loop"; Looking at all potential pairs to satisfy target.
			BigO - N^2
			Space - n/a
			// Not Scalable


	*/

	/*
		# E (Execute Brute)
		Compare each element with all the others
		Double for loop; until match
		BigO - N^2
		Space - n/a
	*/

	/*
		# E (Execute Iterative)
		"Single For Loop"; Use a hashmap to store complements and its index
		Big0 - N
		Space - N
		// Essentially
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

	/*
		# R (Retro)
		Unit Tests:

		Conclusion:

	*/
}
