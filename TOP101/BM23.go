package main
import . "nc_tools"
/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 
 * @param root TreeNode类 
 * @return int整型一维数组
*/
func preorderTraversal(root *TreeNode) []int {
	ret := []int{}
	preorder(&ret, root)
	return ret
}

func preorder(ret *[]int, node *TreeNode) {
	if node == nil {
		return
	}
	*ret = append(*ret, node.Val)
	preorder(ret, node.Left)
	preorder(ret, node.Right)
}
