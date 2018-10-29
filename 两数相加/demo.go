package main

import (
	"fmt"
	"strconv"
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
	var str1 string
	var str2 string

	currentL1 := l1
	str1 = strconv.Itoa(currentL1.Val)
	for currentL1.Next != nil {
		currentL1 = currentL1.Next
		str1 += strconv.Itoa(currentL1.Val)
	}

	currentL2 := l2
	str2 = strconv.Itoa(currentL2.Val)
	for currentL2.Next != nil {
		currentL2 = currentL2.Next
		str2 += strconv.Itoa(currentL2.Val)
	}

	num1, _ := strconv.Atoi(Strrev(str1))
	num2, _ := strconv.Atoi(Strrev(str2))
	sum := num1 + num2

	f := &Feed{}
	for i := len(strconv.Itoa(sum)[:]) - 1; i >= 0; i-- {
		res, _ := strconv.Atoi(string(strconv.Itoa(sum)[:][i]))
		f.insert(&ListNode{Val: res})
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
	f1 := &Feed{}
	f1.insert(&ListNode{Val: 2})
	f1.insert(&ListNode{Val: 4})
	f1.insert(&ListNode{Val: 3})

	f2 := &Feed{}
	f2.insert(&ListNode{Val: 5})
	f2.insert(&ListNode{Val: 6})
	f2.insert(&ListNode{Val: 4})

	result := addTwoNumbers(f1.list, f2.list)

	fmt.Println(result)

	f3 := &Feed{}
	f4 := &Feed{}

	f3.insert(&ListNode{Val: 1})
	f3.insert(&ListNode{Val: 8})
	f4.insert(&ListNode{Val: 0})
	fmt.Println(addTwoNumbers(f3.list, f4.list))
}
