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

var result = make(map[int][]int)


func cengci(b *BinaryTree, level int){
	if _,ok := result[level]; !ok {
		result[level] = []int{}
	}
	if b != nil {
		if level % 2 == 0{
			result[level] = append(result[level], b.this)
		} else {
			result[level] = append([]int{b.this}, result[level]...)
		}

	}
	if b.left != nil {
		cengci(b.left, level + 1)
	}
	if b.right !=nil {
		cengci(b.right, level + 1)
	}
	return
}

func findSumNumber(b *BinaryTree, target int) []int {
	sum := []int{}
	return sum
}

var sums []int
func findOneSum(b *BinaryTree, sum []int) []int {
	if b != nil {
		sum = append(sum, b.this)
	}
	if b.left == nil && b.right == nil {
		fmt.Println(sum)
	}
	if b.left != nil {
		findOneSum(b.left, append([]int{}, sum...))

	}
	if b.right != nil {
		findOneSum(b.right, append([]int{}, sum...))
	}
	return sum
}

func findMaxLength(b *BinaryTree) int {
	if b == nil {
		return 0
	}
	left := findMaxLength(b.left) + 1
	right := findMaxLength(b.right) + 1
	return max(left, right)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
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


	findOneSum(&t0, []int{})
}
