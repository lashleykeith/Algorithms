package main

import "fmt"

type LinkedList struct {
    Value int
    Next  *LinkedList
}

func MiddleNode(linkedList *LinkedList) *LinkedList {
    later := linkedList
    earlier := linkedList

    for later != nil && later.Next != nil {
        later = later.Next.Next
        earlier = earlier.Next
    }

    return earlier
}

func main() {
    // Create the linked list: 2 -> 7 -> 3 -> 5
    node1 := &LinkedList{Value: 2}
    node2 := &LinkedList{Value: 7}
    node3 := &LinkedList{Value: 3}
    node4 := &LinkedList{Value: 5}

    node1.Next = node2
    node2.Next = node3
    node3.Next = node4

    // Call the MiddleNode function
    result := MiddleNode(node1)

    // Print the result
    fmt.Println(result.Value)
}
