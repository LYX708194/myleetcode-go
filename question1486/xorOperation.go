package question1486

func xorOperation(n int, start int) int {
	ans := start
	for i := 1 ; i < n; i++{
		start += 2
		ans ^= start
	}
	return ans
}
