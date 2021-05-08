package question1723

func minimumTimeRequired(jobs []int, k int) int {
	ans := 0x3f3f3f3f
	sum := make([]int, k)
	n := len(jobs)
	/**
	index:当前遍历任务位置
	used:已经被使用的工人
	max:最大工作时间
	sum:工人负责的情况
	*/
	var dfs func(int, int, int, []int)
	dfs = func(index int, used int, max int, sum []int) {
		if max >= ans {
			return
		}
		if index == n {
			ans = max
			return
		}
		//优先分配还没有任务的
		if used < k {
			sum[used] = jobs[index]
			dfs(index+1, used+1, maxNum(max, sum[used]), sum)
			sum[used] = 0
		}
		for i := 0; i < used; i++ {
			sum[i] += jobs[index]
			dfs(index+1, used, maxNum(max, sum[i]), sum)
			sum[i] -= jobs[index]
		}
	}
	dfs(0, 0, 0, sum)
	return ans
}
func maxNum(a, b int) int {
	if a < b {
		return b
	}
	return a
}
