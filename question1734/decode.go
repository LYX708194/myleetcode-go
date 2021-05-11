package question1734

func decode(encoded []int) []int {
	n := len(encoded)
	//前n个异或
	a := 0
	for i := 1; i <= n+1; i++ {
		a ^= i
	}
	//除了第一个的异或
	b := 0
	for i := 1; i < n; i += 2 {
		b ^= encoded[i]
	}
	ans := make([]int, n+1)
	//第一位
	ans[0] = a ^ b
	for i, v := range encoded {
		ans[i+1] = ans[i] ^ v
	}
	return ans
}
