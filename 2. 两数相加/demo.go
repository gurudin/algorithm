package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Feed struct {
	length int
	list   *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sum []int
	var cnt = 1

	currentL1 := l1
	currentL2 := l2

	sum = append(sum, currentL1.Val)
	for currentL1.Next != nil {
		currentL1 = currentL1.Next
		sum = append(sum, currentL1.Val)
	}

	sum[0] += currentL2.Val
	for currentL2.Next != nil {
		currentL2 = currentL2.Next

		if len(sum) < (cnt + 1) {
			sum = append(sum, currentL2.Val)
		} else {
			sum[cnt] += currentL2.Val
		}
		cnt++
	}

	for i := 0; i < len(sum); i++ {
		if sum[i] >= 10 {
			if len(sum) > i+1 {
				sum[i+1]++
			} else {
				sum = append(sum, 1)
			}

			sum[i] -= 10
		}
	}

	f := &Feed{}
	for _, val := range sum {
		f.insert(&ListNode{Val: val})
	}

	return f.list
}

// Strrev rev
func Strrev(text string) string {
	var tmp string
	for i := len(text) - 1; i >= 0; i-- {
		tmp += string(text[i])
	}

	return tmp
}

func (f *Feed) insert(newNode *ListNode) {
	if f.length == 0 {
		f.list = newNode
	} else {
		currentLn := f.list
		for currentLn.Next != nil {
			currentLn = currentLn.Next
		}
		currentLn.Next = newNode
	}
	f.length++
}

func main() {
	// -----------------
	f1 := &Feed{}
	f1.insert(&ListNode{Val: 2})
	f1.insert(&ListNode{Val: 4})
	f1.insert(&ListNode{Val: 3})

	f2 := &Feed{}
	f2.insert(&ListNode{Val: 5})
	f2.insert(&ListNode{Val: 6})
	f2.insert(&ListNode{Val: 4})
	fmt.Println(addTwoNumbers(f1.list, f2.list))
	// -------end-------

	// -----------------
	f3 := &Feed{}
	f4 := &Feed{}

	f3.insert(&ListNode{Val: 1})
	f3.insert(&ListNode{Val: 8})
	f4.insert(&ListNode{Val: 0})
	fmt.Println(addTwoNumbers(f3.list, f4.list))
	// -------end-------

	// -----------------
	arr1 := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	arr2 := []int{5, 6, 4}
	f5 := &Feed{}
	f6 := &Feed{}
	for _, val := range arr1 {
		f5.insert(&ListNode{Val: val})
	}
	for _, val := range arr2 {
		f6.insert(&ListNode{Val: val})
	}
	fmt.Println(addTwoNumbers(f5.list, f6.list))
	// -------end-------
}
