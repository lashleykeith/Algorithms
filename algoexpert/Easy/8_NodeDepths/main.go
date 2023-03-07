// Solution 1
// O(n) time | O(h) space - where n is the number of nodes in
// the Binary Tree and h is the height of the Binary Tree

package main

type BinaryTree struct{
    Value int
    Left, Right *BinaryTree
}

type Level struct{
    Root *BinaryTree
    Depth int
}


func NodeDepths(root *BinaryTree) int{
    sumOfDepths := []Levels{{Root:root, Depth:0}}
    for len(stack) > 0 {
        top, stack = stack[len(stack)-1], stack[:len(stack)-1]
        node, depth := top.Root, top.Depth
        if node == nil{
            continue
        }

        sumOfDepths += depth

        stack = append(stack, Level{Root: node.Left, Depth: depth +1})
        stack = append(stack, Level{Root: node.Right, Depth: depth +1})
    }
    return sumOfDepths

}


// Solution 2
// O(n) time | O(h) space - where n is the number of nodes
// in  the Binary Tree and h is the height of the Binary Tree

package main

type BinaryTree struct{
    Value int
    Left, Right *BinaryTree
}

func NodeDepths(root *BinaryTree) int{
    return NodeDepthsHelper(root, 0)
}

func NodeDepthsHelper(root *BinaryTree, depth int) int{
    if root == nil{
        return 0
    }
    return depth += NodeDepthsHelper(root.Left, depth+1) + NodeDepthsHelper(root.Right, depth+1)
}


// Solution 3 This only works
package main

type BinaryTree struct{
    Value int
    Left, Right *BinaryTree
}

func NodeDepths(root *BinaryTree) int{
    return NodeDepthsHelper(root, 0)
}

func NodeDepthsHelper(root *BinaryTree, depth int) int{
    if root == nil{
        return 0
    }
    return depth + NodeDepthsHelper(root.Left, depth+1) + NodeDepthsHelper(root.Right, depth+1)
}


/* 
// We need to print the sum of these nodes in this program

{
  "tree": {
    "nodes": [
      {"id": "1", "left": "2", "right": "3", "value": 1},
      {"id": "2", "left": "4", "right": "5", "value": 2},
      {"id": "3", "left": "6", "right": "7", "value": 3},
      {"id": "4", "left": "8", "right": "9", "value": 4},
      {"id": "5", "left": null, "right": null, "value": 5},
      {"id": "6", "left": null, "right": null, "value": 6},
      {"id": "7", "left": null, "right": null, "value": 7},
      {"id": "8", "left": null, "right": null, "value": 8},
      {"id": "9", "left": null, "right": null, "value": 9}
    ],
    "root": "1"
  }
}


 */