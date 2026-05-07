package base

func Mono(nums []int) bool {

	isIncrease := true
	isDecrease := true

	for i := 1; i < len(nums); i++ {

		if nums[i-1] < nums[i] {
			isDecrease = false
		}

		if nums[i-1] > nums[i] {
			isIncrease = false
		}

		if !isIncrease && !isDecrease {
			return false
		}
	}

	return isIncrease || isDecrease
}
