package main

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
