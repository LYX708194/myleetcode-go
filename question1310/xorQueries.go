package question1310

func xorQueries(arr []int, queries [][]int) []int {
	n := len(arr)
	xors := make([]int, n+1)
	for i := 0; i < n; i++ {
		xors[i+1] = xors[i] ^ arr[i]
	}
	m := len(queries)
	ans := make([]int, m)
	for i := 0; i < m; i++ {
		ans[i] = xors[queries[i][0]] ^ xors[queries[i][1]+1]
	}
	return ans
}
