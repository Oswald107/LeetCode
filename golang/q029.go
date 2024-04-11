func divide(dividend int, divisor int) int {
	if dividend == -2147483648 && divisor == -1 {
		return 2147483647
	}

	output := 0
	d1 := abs(dividend)
	d2 := abs(divisor)
	for d1 >= d2 {
		sub := d2
		add := 1
		for d1 >= sub<<1 {
			sub <<= 1
			add <<= 1
		}
		d1 -= sub
		output += add
	}

	if (dividend < 0) != (divisor < 0) {
		return -output
	}
	return output
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}