package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

    var head,tail *ListNode
    head,tail = nil,nil
    jinwei := 0
    p1,p2 := l1,l2
    for ;p1 != nil && p2 != nil;p1,p2 = p1.Next,p2.Next {
        p := &ListNode{}
        p.Val = p1.Val + p2.Val + jinwei

        if p.Val >= 10 {
            p.Val -= 10
            jinwei = 1
        }else{
            jinwei = 0
        }

        if head== nil {
           head = p
           tail = head
        }else{
           tail.Next = p 
           tail = p
        }
    }
    var more *ListNode
    more = nil
    if p1 != nil {
        more = p1
    }else if p2 != nil {
        more = p2
    }

    if more != nil {
        for ;more != nil;more = more.Next {
            p := &ListNode{more.Val, nil}
            p.Val += jinwei
            if p.Val >= 10 {
                p.Val -= 10
                jinwei = 1
            }else{
                jinwei = 0
            }
            tail.Next = p
            tail = p
        }
    }

    if jinwei == 1 {
        p := &ListNode{}
        p.Val = jinwei
        p.Next = nil
        tail.Next = p
        tail = p
    }
    
    return head
}

func main(){
    l1 := &ListNode{1,nil}
    l2 := &ListNode{9,nil}
    p := &ListNode{9,nil}
    l2.Next = p

    h := addTwoNumbers(l1, l2)
    for ;h != nil;h = h.Next {
        fmt.Println(h.Val)
    }
}
