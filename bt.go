package BinTree

import "log"

func init() {
	log.SetFlags(0)
	log.SetPrefix("")
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 1325m Delete Leaves With a Given Value
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = removeLeafNodes(root.Left, target)
	root.Right = removeLeafNodes(root.Right, target)
	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}
	return root
}
