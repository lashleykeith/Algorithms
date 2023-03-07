// Not working
package main

import "fmt"

type BinaryTree struct {
    Value int
    Left  *BinaryTree
    Right *BinaryTree
}

func NodeDepths(root *BinaryTree) (int, int) {
    return NodeDepthsHelper(root, 0, 0)
}

func NodeDepthsHelper(root *BinaryTree, depth int, sum int) (int, int) {
    if root == nil {
        return 0, sum
    }
    leftDepth, sum := NodeDepthsHelper(root.Left, depth+1, sum)
    rightDepth, sum := NodeDepthsHelper(root.Right, depth+1, sum)
    nodeDepth := depth + leftDepth + rightDepth
    sum += root.Value
    fmt.Printf("Node value: %v, depth: %v\n", root.Value, nodeDepth)
    return nodeDepth, sum
}

func main() {
    // Define the binary tree
    root := &BinaryTree{
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

    // Compute and print the sum of all nodes in the binary tree
    _, sum := NodeDepths(root)
    fmt.Println("Sum of all nodes:", sum)
}
