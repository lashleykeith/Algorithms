// O(n) time | O(n) space - where n is the number of nodes in the Binary Tree
package main

type BinaryTree struct {
    Value int
    Left *BinaryTree
    Right *BinaryTree
}

func BranchSums(root *BinaryTree) []int{
    sums := []int{}
    calculateBranchSums(root, 0, &sums)
    return sums
}

func calculateBranchSums(node *BinaryTree, runningSum int, sums* []int){
    if node == nil{
        return
    }

    runningSum += node.Value
    if node.Left == nil && node.Right == nil{
        *sums = append(*sums, runningSum)
        return
    }
    calculateBranchSums(node.Left, runningSum, sums)
    calculateBranchSums(node.Right, runningSum, sums)
}

/* 
{
  "tree": {
    "nodes": [
      {"id": "1", "left": "2", "right": "3", "value": 1},
      {"id": "2", "left": "4", "right": "5", "value": 2},
      {"id": "3", "left": "6", "right": "7", "value": 3},
      {"id": "4", "left": "8", "right": "9", "value": 4},
      {"id": "5", "left": "10", "right": null, "value": 5},
      {"id": "6", "left": null, "right": null, "value": 6},
      {"id": "7", "left": null, "right": null, "value": 7},
      {"id": "8", "left": null, "right": null, "value": 8},
      {"id": "9", "left": null, "right": null, "value": 9},
      {"id": "10", "left": null, "right": null, "value": 10}
    ],
    "root": "1"
  }
}

*/