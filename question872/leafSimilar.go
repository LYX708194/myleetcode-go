package question872

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	vals1 := make([]int, 0)
	vals2 := make([]int, 0)
	dfs(root1, &vals1)
	dfs(root2, &vals2)
	if len(vals1) != len(vals2) {
		return false
	}
	for i, v := range vals1 {
		if v != vals2[i] {
			return false
		}
	}
	return true
}

func dfs(node *TreeNode, vals *[]int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		*vals = append(*vals, node.Val)
		return
	}
	dfs(node.Left, vals)
	dfs(node.Right, vals)
}
