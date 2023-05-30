package main

import "fmt"

type LinkedList struct {
    Value int
    Next  *LinkedList
}

func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
    curr := linkedList
    for curr != nil && curr.Next != nil {
        if curr.Value == curr.Next.Value {
            curr.Next = curr.Next.Next
        } else {
            curr = curr.Next
        }
    }

    return linkedList
}

func createLinkedList(values []int) *LinkedList {
    if len(values) == 0 {
        return nil
    }

    head := &LinkedList{Value: values[0]}
    curr := head
    for i := 1; i < len(values); i++ {
        node := &LinkedList{Value: values[i]}
        curr.Next = node
        curr = node
    }

    return head
}

func printLinkedList(linkedList *LinkedList) {
    curr := linkedList
    for curr != nil {
        fmt.Printf("%d ", curr.Value)
        curr = curr.Next
    }
    fmt.Println()
}

func main() {
    values := []int{1, 5, 3, 4, 1, 6, 6, 6, 8, 7, 13, 5}
    linkedList := createLinkedList(values)

    fmt.Println("Original Linked List:")
    printLinkedList(linkedList)

    uniqueItems := RemoveDuplicatesFromLinkedList(linkedList)

    fmt.Println("Linked List after removing duplicates:")
    printLinkedList(uniqueItems)
}
