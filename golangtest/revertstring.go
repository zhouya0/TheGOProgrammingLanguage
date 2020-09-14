package main

import "fmt"

var testSting = "abcacdef"

func revert(s string) string{
	list := []byte(s)
	length := len(list)
	for i:=0; i<length/2; i++ {
		list[i], list[length-i-1] = list[length-i-1],list[i]
	}
	return string(list)
}

func isHuiWen(num int) bool {
	// 比较一半确实是可以的，其奥义就在于可以直接比较大小
	// 因为不管是不是回文数，其位数总是可以绝对的。
	if num < 10 {
		return true
	}
	var revertNum int
	for revertNum < num {
		revertNum = num%10 +  revertNum*10
		num = num/10
	}
	return revertNum == num || revertNum/10 == num
}

func mostNotDup(s string) int {
	list := []byte(s)
	var start int
	var maxLength int
	dupMap := make(map[byte]int)
	for k,v := range list {
		// 算法的秘诀是，最大的长度不可能建立在已有的重复上
		if lastV,ok :=  dupMap[v]; ok && lastV >= start{
			start = lastV +1
		}

		if (k - start + 1) > maxLength {
			maxLength = k - start + 1
		}
		dupMap[v] = k

	}
	return maxLength
}

func main() {
	test := 123454321
	fmt.Println(isHuiWen(test))
}
