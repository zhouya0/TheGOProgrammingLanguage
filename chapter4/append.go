package main

import (
	"fmt"
)

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	// 所以这个else表示了不一定会发生内存分配
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		// This is really costy!!!
		// Try everything to avoid this action!!!
		z = make([]int, zlen, zcap)
		// delete是内置的用于map的函数
		// copy是内置的用于slice的函数
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

```
如何单看这个appendInt的实现，可以看出，它有时候是使用原slice的内存，有时候是开辟新的内存空间：
TheGoProgrammingLanguage 原文
“因此，我们并不知道append调用是否导致了内存的重新分配， 因此我们也不能确认新的slice和原始的slice
是否引用的是相同的底层数组空间。同样，我们不能确定在原先的slice上的操作是否会影响到新的slice。因此，
通常将append返回的结果直接赋值给输入的slice变量。

runes = append(runes, r)

```
func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}


