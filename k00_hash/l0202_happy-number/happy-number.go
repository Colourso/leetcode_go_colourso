package l0202_happy_number

func isHappy(n int) bool {
	numMap := make(map[int]int)
	numMap[n] = 1

	for {
		num := calculateHappyNum(n)
		if num == 1 {
			return true
		} else {
			if numMap[num] > 0 {
				return false
			} else {
				numMap[num] = 1
			}
		}
		n = num
	}
}

func calculateHappyNum(n int) int {
	sum := 0
	for n > 9 {
		num := n % 10 // 取最后一位
		sum += num * num
		n = n / 10 // 舍去各位数
	}
	sum += n * n
	return sum
}
