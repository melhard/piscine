package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	right := BTreeLevelCount(root.Right)
	lift := BTreeLevelCount(root.Left)
	if right > lift {
		return right + 1
	}
	return lift + 1
}
