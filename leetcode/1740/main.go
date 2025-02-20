package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDistance(root *TreeNode, p, q int) int {
	var found bool
	var res int

	var dfs func(node *TreeNode) int

	dfs = func(node *TreeNode) int {
		if found {
			return 0
		}

		if node == nil {
			return 0
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		if left == 0 && right == 0 {
			if node.Val == p || node.Val == q {
				return 1
			}

			return 0
		}

		if left != 0 && right != 0 && !found {
			res = left + right
			found = true
			return 0
		}

		var branch int
		if left == 0 {
			branch = right
		} else {
			branch = left
		}

		if node.Val == p || node.Val == q {
			res = branch
			return 0
		}

		return branch + 1
	}

	dfs(root)
	return res
}

func main() {
	//test

	//test case 1528
	//expected 3
	//         3
	//        / \
	//       5   1528
	//      / \ / \
	//     6  2 0  8
	//       / \
	//      7   4

	//p = 5, q = 1528
	//distance = 3
	//5 -> 2 -> 1528
	//5 -> 3 -> 1528

	//create tree
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 8}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 4}

	fmt.Println(findDistance(root, 7, 8))

}
