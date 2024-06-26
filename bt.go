package BinTree

import (
	"container/list"
	"fmt"
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

func (o TreeNode) String() string {
	l, r := '*', '*'
	if o.Left == nil {
		l = '-'
	}
	if o.Right == nil {
		r = '-'
	}
	return fmt.Sprintf("{%c %d %c}", l, o.Val, r)
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

// 1038m Binary Search Tree to Greater Sum Tree
func bstToGst(root *TreeNode) *TreeNode {
	nVal := 0
	var RInOrder func(*TreeNode)
	RInOrder = func(n *TreeNode) {
		if n == nil {
			return
		}

		RInOrder(n.Right)
		nVal += n.Val
		n.Val = nVal
		RInOrder(n.Left)
	}

	RInOrder(root)
	return root
}

// 1382m Balance a Binary Search Tree
func balanceBST(root *TreeNode) *TreeNode {
	V := []int{}

	var InOrder func(*TreeNode)
	InOrder = func(n *TreeNode) {
		if n == nil {
			return
		}
		InOrder(n.Left)
		V = append(V, n.Val)
		InOrder(n.Right)
	}

	InOrder(root)

	var Build func(l, r int) *TreeNode
	Build = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}

		m := l + (r-l)>>1
		return &TreeNode{Val: V[m], Left: Build(l, m-1), Right: Build(m+1, r)}
	}

	return Build(0, len(V)-1)
}

// 2385m Amount of Time for Binary Tree to Be Infected
func amountOfTime(root *TreeNode, start int) int {
	lsAdj := map[int][]int{}

	var Walk func(*TreeNode, int)
	Walk = func(n *TreeNode, pVal int) {
		if n == nil {
			return
		}
		lsAdj[pVal] = append(lsAdj[pVal], n.Val)
		lsAdj[n.Val] = append(lsAdj[n.Val], pVal)

		Walk(n.Left, n.Val)
		Walk(n.Right, n.Val)
	}
	Walk(root.Left, root.Val)
	Walk(root.Right, root.Val)

	log.Print("BT -> G :: ", lsAdj)

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
