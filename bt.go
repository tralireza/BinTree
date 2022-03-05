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

// 979m Distribute Coins in Binary Tree
func distributeCoins(root *TreeNode) int {
	abs := func(v int) int {
		if v < 0 {
			return -v
		}
		return v
	}

	moves := 0
	var Walk func(*TreeNode) int
	Walk = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		l, r := Walk(n.Left), Walk(n.Right)
		moves += abs(l) + abs(r)
		return n.Val - 1 + l + r
	}

	Walk(root)
	return moves
}
