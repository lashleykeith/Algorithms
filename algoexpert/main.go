// Solution 1
// O(n) time | O(h) space - where n is the number of nodes in
// the Binary Tree and h is the height of the Binary Tree

package main


type BinaryTree struct{
    Value            int
    Left, Right      *BinaryTree
}

type Level struct{
    Root        *BinaryTree
    Depth       int
}

func NodeDepths(root, *BinaryTree) int{
    sumofDepth :=  0
    stack := {{Root:root, Depth:0}} 
}

/*
Make Binary Tree struct with Left and Right places

Make Level Struct with a Root that inherits from the Binary Tree and and can measure the Depth of our tree

Make a function NodeDepths that has a root and points to the BinaryTree struct

We need to initialize the sum of Depths and we also want to initialize a stack.  This stack will inherit from the Level Struct and will make key-value pairs.  The Root at root and the Depth will start at 0
*/