package piscine

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyInorder(root.Left, f)
	if _, err := f(root.Data); err != nil {
		return
	}
	BTreeApplyInorder(root.Right, f)
}
