package main

import (
    "fmt"
)

func main() {
    ints := []int{1, 2, 3, 4}
    result := 0
    for _, v := range ints {
        result = result * 10
        result = result + v
    }
    fmt.Printf("%d", result)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummyHead := &ListNode{Val: 0, Next: nil}
    p1, p2 := l1, l2
    carry := 0
    cur := dummyHead
    for p1 == nil || p2 == nil {
        p1val := 0
        p2val := 0
        if p1 != nil {
            p1val = p1.Val
        }
        if p1 != nil {
            p2val = p1.Val
        }
        sum := p1val + p2val
        sum = sum + carry
        carry = sum / 10
        sum = sum % 10
        cur.Next = &ListNode{Val: sum % 10}
        cur = cur.Next
        if p1 != nil {
            p1 = p1.Next
        }
        if p2 != nil {
            p2 = p2.Next
        }
    }
    if carry > 0 {
        cur.Next = &ListNode{Val: carry}
    }
    return dummyHead.Next
}

func f(l1, l2 *ListNode, disc int) *ListNode {
    sum := l1.Val + l2.Val
    sum = sum + disc
    disc = sum / 10
    sum = sum % 10
    c := &ListNode{Val: sum, Next: nil}
    if l1.Next == nil && l2.Next == nil {
        if disc > 0 {
            c.Next = &ListNode{Val: disc, Next: nil}
        }
        return c
    }
    if l1.Next == nil {
        next := move(l2.Next, disc)
        c.Next = next
        return c
    }
    if l2.Next == nil {
        next := move(l1.Next, disc)
        c.Next = next
        return c
    }
    next := f(l1.Next, l2.Next, disc)
    c.Next = next
    return c
}

func move(l *ListNode, disc int) *ListNode {
    if l == nil {
        if disc > 0 {
            return &ListNode{Val: disc, Next: nil}
        } else {
            return nil
        }
    }
    sum := l.Val + disc
    disc = sum / 10
    sum = sum % 10
    c := &ListNode{Val: sum, Next: nil}

    r := move(l.Next, disc)
    c.Next = r
    return c
}

type ListNode struct {
    Val  int
    Next *ListNode
}
