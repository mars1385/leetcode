package binaryearch

func Search(nums []int, target int) int {

	index := len(nums) / 2
	num := nums[index]
	isLow := num > target
	for num != target && index >= 0 && index < len(nums) {

		if isLow {
			index--
		} else {
			index++
		}

		if index >= 0 && index < len(nums) {
			num = nums[index]
		}
	}

	if num != target {
		return -1
	}

	return index
}
