// O(v + e) time | O(v) space

package main

type Node struct{
    Name        string
    Children    []*Node
}

func (n *Node) DepthFirstSearch(array []string) []string{
    array = append(array, n.Name)
    for _, child := range n.Children{
        array = child.DepthFirstSearch(array)
    }
    return array
}


/* 
// We need build a struct for the name and children

// Write the function DepthFirstSearch that points to the Node struct and creates a string array

// append the Name in the array to the array we are searching the depth in

// create a for loop that will look through the child of each map in the Children and will run through each child adding it.

 */