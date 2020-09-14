package main

import "fmt"

type BinaryTree struct {
	this int
	left *BinaryTree
	right *BinaryTree
}

func qianxu(b *BinaryTree){
	recurseqianxu(b)

}
// append会发生值拷贝！！！
// append会发生值拷贝！！！
func recurseqianxu(b *BinaryTree) {
	fmt.Print(b.this)
	if b.left != nil {
		recurseqianxu(b.left)
	}
	if b.right != nil {
		recurseqianxu(b.right)
	}
}

func recursezhongxu(b *BinaryTree) {
	if b.left != nil {
		recursezhongxu(b.left)
	}
	fmt.Print(b.this)
	if b.right != nil {
		recursezhongxu(b.right)
	}
}

func recurserhouxu(b *BinaryTree) {
	if b.left != nil {
		recurserhouxu(b.left)
	}

	if b.right != nil {
		recurserhouxu(b.right)
	}

	fmt.Print(b.this)
}

func revertTree(b *BinaryTree) {
	b.left, b.right = b.right, b.left
	if b.left != nil {
		revertTree(b.left)
	}
	if b.right != nil {
		revertTree(b.right)
	}

}

func findSumNumber(b *BinaryTree, number int) bool {
	number = number - b.this
	if b.left != nil {
		boolLeft := findSumNumber(b.left, number)
		if boolLeft {
			return true
		}
	}

	if b.right != nil {
		boolRight := findSumNumber(b.right, number)
		if boolRight {
			return true
		}
	}
	if number == 0 {
		return true
	}
	return false
}

var result = make(map[int][]int)

func normalOrder(b *BinaryTree, level int ) {
	if result[level] == nil {
		result[level] = []int{}
	}

	result[level] = append(result[level], b.this)

	if b.left != nil {
		normalOrder(b.left, level+1)
	}
	if b.right != nil {
		normalOrder(b.right, level+1)
	}
}

func main() {
	t9 := BinaryTree{this: 9}
	t8 := BinaryTree{this: 8}
	t7 := BinaryTree{this: 7}
	t6 := BinaryTree{this: 6}
	t5 := BinaryTree{this: 5}
	t4 := BinaryTree{this: 4, left: &t9}
	t3 := BinaryTree{this:3, left: &t7, right: &t8}
	t2 := BinaryTree{this:2, left: &t5, right: &t6}
	t1 := BinaryTree{this: 1, left: &t3, right: &t4}
	t0 := BinaryTree{this:0, left: &t1, right: &t2}

	fmt.Println(findSumNumber(&t0, 9))
	normalOrder(&t0, 0)
	fmt.Println(result)
}
