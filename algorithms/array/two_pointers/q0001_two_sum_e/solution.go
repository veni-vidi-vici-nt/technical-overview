package q0001_two_sum_e

// TwoSum returns the indices of the two numbers that add up to the target.
func TwoSum(nums []int, target int) []int {
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
