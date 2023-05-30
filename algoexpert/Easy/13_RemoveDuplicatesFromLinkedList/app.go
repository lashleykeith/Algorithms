package main


// This is an input struct.  Do no edit.

type LinkedList struct{
    Value int
    Next *LinkedList
}


func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList{
    curr := linkedList
    for curr.Next != nil {
            if curr.Value == curr.Next.Value{
                curr.Next = curr.Next.Next
            } else {
                curr = curr.Next
            }
        }

        return linkedList
}