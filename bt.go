package BinTree

import (
	"container/list"
	"log"
)

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

// 1026m Maximum Difference Between Node and Ancestor
func maxAncestorDiff(root *TreeNode) int {
	x := 0

	var Diff func(n *TreeNode, mn, mx int)
	Diff = func(n *TreeNode, mn, mx int) {
		x = max(x, mx-mn)

		if n.Left != nil {
			ln, lx := min(n.Left.Val, mn), max(n.Left.Val, mx)
			Diff(n.Left, ln, lx)
		}
		if n.Right != nil {
			rn, rx := min(n.Right.Val, mn), max(n.Right.Val, mx)
			Diff(n.Right, rn, rx)
		}
	}

	Diff(root, root.Val, root.Val)
	return x
}

// 2385m Amount of Time for Binary Tree to Be Infected
func amountOfTime(root *TreeNode, start int) int {
	lsAdj := map[int][]int{}

	var Walk func(*TreeNode, int)
	Walk = func(n *TreeNode, pVal int) {
		if n == nil {
			return
		}

		if pVal != -1 {
			lsAdj[pVal] = append(lsAdj[pVal], n.Val)
			lsAdj[n.Val] = append(lsAdj[n.Val], pVal)
		}
		Walk(n.Left, n.Val)
		Walk(n.Right, n.Val)
	}
	Walk(root, -1)
	log.Print("BT -> Graph :: ", lsAdj)

	Q := list.New()
	Q.PushBack(start)
	Vis := map[int]bool{start: true}
	t := 0
	for Q.Len() > 0 {
		for range Q.Len() {
			v := Q.Remove(Q.Front()).(int)
			for _, u := range lsAdj[v] {
				if !Vis[u] {
					Vis[u] = true
					Q.PushBack(u)
				}
			}
		}
		t++
	}
	return t - 1
}
