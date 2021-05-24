package leetcode_go

func xorGame(nums []int) bool {
	var sum int
	for i := range nums {
		sum ^= nums[i]
	}
	return sum == 0 || len(nums)%2 == 0
}
