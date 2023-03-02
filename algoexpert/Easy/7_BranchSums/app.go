package main

import (
    "fmt"
)

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

func main() {
    tree := &BinaryTree{
        Value: 1,
        Left: &BinaryTree{
            Value: 2,
            Left: &BinaryTree{
                Value: 4,
                Left: &BinaryTree{
                    Value: 8,
                },
                Right: &BinaryTree{
                    Value: 9,
                },
            },
            Right: &BinaryTree{
                Value: 5,
                Left: &BinaryTree{
                    Value: 10,
                },
            },
        },
        Right: &BinaryTree{
            Value: 3,
            Left: &BinaryTree{
                Value: 6,
            },
            Right: &BinaryTree{
                Value: 7,
            },
        },
    }

    sums := BranchSums(tree)
    fmt.Println(sums) // Output: [15 16 18 10 11]
}
