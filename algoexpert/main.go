// Solution 1
// Average: O(log(n)) time | O(log(n)) space
// Worst: O(n) time | O(n) space
package main

type BST struct{
    Value int

    Left *BST
    Right *BST
}

func (tree *BST) FindClosestValue(target int) int{
    return tree.findClosestValue(target, tree.Value)
}

func (tree *BST) findClosestValue(target, closest int) int{
    if absdiff(target, closest) > absdiff(target, tree.Value){
        closest = tree.Value
    }
    if target < tree.Value && tree.Left != nil{
        return tree.Left.findClosestValue(target, closest)
    } else if target > tree.Value && tree.Right != nil{
        return tree.Right.findClosestValue(target, closest)
    }
    return closest
}

func absdiff(a, b int) int{
    if a > b {
        return a - b
    }
    return b - a
}


// Solution 2
// Average: O(log(n)) time | O(1) space
// Worst: O(n) time | O(1) space

package main

type BST struct{
    Value int

    Left *BST
    Right *BST
}

func(tree *BST) FindClosestValue(target int) int {
    return tree.findClosestValue(target, tree.Value)
}

func (tree *BST) findClosestValue(target, closest int) int{
    currentnode := tree
    for currentnode != nil{
        if absdiff(target, closest) > absdiff(target, currentnode.Value){
            closest = currentnode.Value
        }
        if target < currentnode.Value{
            currentnode = currentnode.Left
        } else if target > currentnode.Value{
            currentnode = currentnode.Right
        } else {
            break
        }
    }
    return closest
}

func absdiff(a, b int) int{
    if a > b {
        return a - b
    }
    return b - a
}

/*
{
  "tree": {
    "nodes": [
      {"id": "10", "left": "5", "right": "15", "value": 10},
      {"id": "15", "left": "13", "right": "22", "value": 15},
      {"id": "22", "left": null, "right": null, "value": 22},
      {"id": "13", "left": null, "right": "14", "value": 13},
      {"id": "14", "left": null, "right": null, "value": 14},
      {"id": "5", "left": "2", "right": "5-2", "value": 5},
      {"id": "5-2", "left": null, "right": null, "value": 5},
      {"id": "2", "left": "1", "right": null, "value": 2},
      {"id": "1", "left": null, "right": null, "value": 1}
    ],
    "root": "10"
  },
  "target": 12
}
*/