package main

import (
    "fmt"
    "testing"
)

func TestAddTwoNumbers(t *testing.T) {

    node1 := fillLinkedList([]int{5})
    node2 := fillLinkedList([]int{5})
    a := addTwoNumbers(node1, node2)
    fmt.Println(a)
}

func fillLinkedList(from []int) *ListNode {
    var prev *ListNode
    var first *ListNode;
    for _, v := range from {
        if prev == nil {
            prev = &ListNode{Val: v, Next: nil}
            first = prev
        } else {
            cur := &ListNode{Val: v, Next: nil }
            prev.Next = cur
            prev = cur
        }
    }
    return first
}
