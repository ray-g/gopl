package minmax

func min(nums ...int) int {
	if len(nums) == 0 {
		panic("No Args")
	}
	m := nums[0]
	for _, n := range nums {
		if n < m {
			m = n
		}
	}
	return m
}

func max(nums ...int) int {
	if len(nums) == 0 {
		panic("No Args")
	}
	m := nums[0]
	for _, n := range nums {
		if n > m {
			m = n
		}
	}
	return m
}

func min2(base int, nums ...int) int {
	m := base
	for _, n := range nums {
		if n < m {
			m = n
		}
	}
	return m
}

func max2(base int, nums ...int) int {
	m := base
	for _, n := range nums {
		if n > m {
			m = n
		}
	}
	return m
}
