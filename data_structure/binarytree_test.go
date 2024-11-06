package data_structure

import (
	"fmt"
	"testing"
)

var treeNode *TreeNode = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 4,
			Right: &TreeNode{
				Val:   6,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 8},
			},
		},
	},
	Right: &TreeNode{
		Val:   3,
		Right: &TreeNode{Val: 5},
	},
}

func TestPreorderTraversalRecursion(t *testing.T) {
	preorderTraversalRecursion(treeNode)
}

func TestPreorderTraversalNotRecursion(t *testing.T) {
	result := preorderTraversalNotRecursion(treeNode)
	fmt.Println(result)
}

func TestInorderTraversalNotRecursion(t *testing.T) {
	result := inorderTraversalNotRecursion(treeNode)
	fmt.Println(result)
}

func TestPostorderTraversalNotRecursion(t *testing.T) {
	result := postorderTraversalNotRecursion(treeNode)
	fmt.Println(result)
}

func TestPreorderTraversalDFS(t *testing.T) {
	result := preorderTraversalDFS(treeNode)
	fmt.Println(result)
}

func TestPreorderTraversalDFSDivide(t *testing.T) {
	result := preorderTraversalDFSDivide(treeNode)
	fmt.Println(result)
}

func TestLevelOrderBFS(t *testing.T) {
	result := levelOrderBFS(treeNode)
	fmt.Println(result)
}

func TestMergeSort(t *testing.T) {
	nums := []int{3, 2, 9, 4, 6, 8, 1, 5}
	result := MergeSort(nums)
	fmt.Println(result)
}

func TestQuickSort(t *testing.T) {
	nums := []int{3, 2, 9, 4, 6, 8, 1, 5}
	result := QuickSort(nums)
	fmt.Println(result)
}
