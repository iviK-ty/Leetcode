func moveZeroes(nums []int) {
	write := 0

	for i := range len(nums) {
		if nums[i] != 0 {
			nums[write] = nums[i]
			write++

		}
	}
	for i := write; i < len(nums); i++ {
		nums[i] = 0
	}
	fmt.Println(nums)
}