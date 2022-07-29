/*
* @Author: Ivy
* @Date: 2022/4/26 9:22 PM
**/
package main

import (
	"fmt"
)

//ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := ListNode{Val: 1}
	b := ListNode{Val: 2}
	c := ListNode{Val: 2}
	d := ListNode{Val: 1}
	a.Next = &b
	b.Next = &c
	c.Next = &d

	answer := isPalindrome(&a)
	fmt.Println(answer)

	a2 := ListNode{Val: 1}

	answer2 := isPalindrome(&a2)
	fmt.Println(answer2)

}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	var array []int
	for head.Next != nil {
		array = append(array, head.Val)
		head = head.Next
	}

	array = append(array, head.Val)

	for i := 0; i < len(array)/2; i++ {
		if array[i]-array[len(array)-i-1] == 0 {
			continue
		} else {
			return false
		}
	}
	return true
}
