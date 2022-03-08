package BinTree

import (
	"container/list"
	"fmt"
	"log"
	"testing"
)

func init() {
	log.Print("> Binary Tree")
}

// 1325m Delete Leaves With a Given Value
func Test1325(t *testing.T) {
	Iterative := func(root *TreeNode, target int) *TreeNode {
		Q := list.New()

		n := root
		var lVis *TreeNode
		for Q.Len() > 0 || n != nil {
			for n != nil {
				Q.PushBack(n)
				n = n.Left
			}
			pk := Q.Back().Value.(*TreeNode)
			if pk.Right != nil && pk.Right != lVis {
				n = pk.Right
			} else {
				// postOrder Visit
				lVis = Q.Remove(Q.Back()).(*TreeNode)

				if lVis.Left == nil && lVis.Right == nil && lVis.Val == target {
					if Q.Len() == 0 {
						return nil
					}

					p := Q.Back().Value.(*TreeNode)
					if p.Left == lVis {
						p.Left = nil
					} else if p.Right == lVis {
						p.Right = nil
					}
				}
			}
		}
		return root
	}

	var Draw func(*TreeNode)
	Draw = func(n *TreeNode) {
		if n != nil {
			Draw(n.Left)
			l, r := '*', '*'
			if n.Left == nil {
				l = '-'
			}
			if n.Right == nil {
				r = '-'
			}
			fmt.Printf("{%c %d %c}", l, n.Val, r)
			Draw(n.Right)
		}
	}

	type T = TreeNode
	for _, f := range []func(*TreeNode, int) *TreeNode{removeLeafNodes, Iterative} {
		r := &T{1, &T{2, &T{Val: 2}, nil}, &T{3, &T{Val: 2}, &T{Val: 4}}}
		Draw(r)
		fmt.Print("  ->  ")
		Draw(f(r, 2))
		fmt.Println()
	}
}

// 979m Distribute Coins in Binary Tree
func Test979(t *testing.T) {
	type T = TreeNode
	log.Print("2 ?= ", distributeCoins(&T{3, &T{Val: 0}, &T{Val: 0}}))
	log.Print("3 ?= ", distributeCoins(&T{0, &T{Val: 3}, &T{Val: 0}}))
}

// 1026m Maximum Difference Between Node and Ancestor
func Test1026(t *testing.T) {
	type T = TreeNode

	log.Print("7 ?= ", maxAncestorDiff(&T{8, &T{3, &T{Val: 1}, &T{6, &T{Val: 4}, &T{Val: 7}}}, &T{10, nil, &T{14, &T{Val: 13}, nil}}}))
}

// 2385m Amount of Time for Binary Tree to Be Infected
func Test2385(t *testing.T) {
	OnPass := func(root *TreeNode, start int) int {
		t := 0

		var Walk func(*TreeNode) int
		Walk = func(n *TreeNode) int {
			if n == nil {
				return 0
			}

			l, r := Walk(n.Left), Walk(n.Right)

			log.Print(l, r, n)

			if n.Val == start {
				t = max(l, r)
				return -1
			} else if l >= 0 && r >= 0 {
				return max(l, r) + 1
			} else {
				if l < 0 {
					t = max(t, r-l)
					return l - 1
				} else {
					t = max(t, l-r)
					return r - 1
				}
			}
		}

		log.Print(" -> ", Walk(root))
		return t
	}

	type T = TreeNode

	for _, f := range []func(*TreeNode, int) int{amountOfTime, OnPass} {
		log.Print("4 ?= ", f(&T{1, &T{5, nil, &T{4, &T{Val: 9}, &T{Val: 2}}}, &T{3, &T{Val: 10}, &T{Val: 6}}}, 3))
		log.Print("0 ?= ", f(&T{Val: 1}, 1))
	}
}
