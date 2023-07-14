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
    values := []int{2, 7, 3, 5}
    head := &LinkedList{}

    current := head
    for _, value := range values {
        current.Next = &LinkedList{Value: value}
        current = current.Next
    }

    // Call the MiddleNode function
    result := MiddleNode(head.Next)

    // Print the result
    fmt.Println(result.Value)
}
