// Solution 2
// Average: O(log(n)) time | O(1) space
// Worst: O(n) time | O(1) spacepackage main
package main

import "fmt"

type BST struct {
    Value int
    Left  *BST
    Right *BST
}

func (tree *BST) FindClosestValue(target int) int {
    return tree.findClosestValue(target, tree.Value)
}

func (tree *BST) findClosestValue(target, closest int) int {
    if tree == nil {
        return closest
    }
    if absdiff(target, closest) > absdiff(target, tree.Value) {
        closest = tree.Value
    }
    if target < tree.Value {
        return tree.Left.findClosestValue(target, closest)
    } else if target > tree.Value {
        return tree.Right.findClosestValue(target, closest)
    } else {
        return closest
    }
}

func absdiff(a, b int) int {
    if a > b {
        return a - b
    }
    return b - a
}



func main() {
    bst := &BST{
        Value: 10,
        Left: &BST{
            Value: 5,
            Left: &BST{
                Value: 2,
                Left: &BST{
                    Value: 1,
                },
            },
            Right: &BST{
                Value: 5,
                Right: &BST{
                    Value: 5,
                },
            },
        },
        Right: &BST{
            Value: 15,
            Left: &BST{
                Value: 13,
                Right: &BST{
                    Value: 14,
                },
            },
            Right: &BST{
                Value: 22,
            },
        },
    }

    target := 12
    closest := bst.FindClosestValue(target)
    fmt.Println(closest)
}

