package BinTree

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	log.Print("> Binary Tree")
}

// 1325m Delete Leaves With a Given Value
func Test1325(t *testing.T) {
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
	r := &T{1, &T{2, &T{Val: 2}, nil}, &T{3, &T{Val: 2}, &T{Val: 4}}}
	Draw(r)
	fmt.Print("  ->  ")
	Draw(removeLeafNodes(r, 2))
	fmt.Println()
}
