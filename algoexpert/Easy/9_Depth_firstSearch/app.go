package main

import "fmt"

type Node struct {
    Name     string
    Children []*Node
}

func (n *Node) DepthFirstSearch(array []string) []string {
    array = append(array, n.Name)
    for _, child := range n.Children {
        array = child.DepthFirstSearch(array)
    }
    return array
}

func main() {
    // define the nodes of the graph
    A := &Node{Name: "A"}
    B := &Node{Name: "B"}
    C := &Node{Name: "C"}
    D := &Node{Name: "D"}
    E := &Node{Name: "E"}
    F := &Node{Name: "F"}
    G := &Node{Name: "G"}
    H := &Node{Name: "H"}
    I := &Node{Name: "I"}
    J := &Node{Name: "J"}
    K := &Node{Name: "K"}

    // define the connections between the nodes
    A.Children = []*Node{B, C, D}
    B.Children = []*Node{E, F}
    D.Children = []*Node{G, H}
    F.Children = []*Node{I, J}
    G.Children = []*Node{K}

    // call DepthFirstSearch on the startNode A
    result := A.DepthFirstSearch([]string{})

    // print the result in the desired order
    expected := []string{"A", "B", "E", "F", "I", "J", "C", "D", "G", "K", "H"}
    if fmt.Sprintf("%v", result) == fmt.Sprintf("%v", expected) {
        fmt.Println("Result matches expected output:", result)
    } else {
        fmt.Println("Result does not match expected output. Expected:", expected, "Actual:", result)
    }
}
