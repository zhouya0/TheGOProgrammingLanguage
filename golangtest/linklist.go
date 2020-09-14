package main

import (
	"fmt"
)

type Link struct {
	Val int
	Next *Link
}


func main(){
	testLink1 := Link{Val: 3}
	testLink2 := Link{Val: 4}
	testLink3 := Link{Val: 5}
	testLink1.Next = &testLink2
	testLink2.Next = &testLink3
	fmt.Println(findCrossPoint(&testLink1, &testLink2))

	revertLink := myRevertLink(&testLink1)
	fmt.Println(revertLink.Val)
	fmt.Println(revertLink.Next.Val)
	fmt.Println(revertLink.Next.Next.Val)
}


func myRevertLink(link *Link) *Link {
	var result *Link
	for link != nil {
		temp := link.Next
		link.Next = result
		result = link
		link = temp
	}
	return result
}







func revertlinklist(link *Link) *Link {
	var result *Link
	for link != nil {
		// 关键在于，不要对link.Next做任何操作
		// 反而是,link做任何操作之后，在把link.Next返回给它
		test := link.Next
		link.Next = result
		result = link
		link = test
	}
	return result
}

func findCrossPoint(l1 *Link, l2 *Link) bool {
	listl1 := []*Link{}
	listl2 := []*Link{}
	for l1.Next != nil {
		listl1 = append(listl1, l1)
		l1 = l1.Next
	}
	for l2.Next != nil {
		listl2 = append(listl2, l2)
		l2 = l2.Next
	}
	length1 := len(listl1)
	length2 := len(listl2)
	if length1 > length2 {
		listl1 = listl1[(length1 - length2):]
	} else {
		listl2 = listl2[(length2 - length1):]
	}

	for i := 0 ; i < len(listl1); i++ {
		if listl1[i] == listl2[i] {
			return true
		}
	}
	return false
}
