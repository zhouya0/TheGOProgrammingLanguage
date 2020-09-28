package main

import "fmt"

func fib(n int) int {
	start := 0
	next := 1
	var result int
	if n == 1 {
		return start
	}
	if n == 2 {
		return next
	}
	for i:=2; i<n;i++{
		result = start + next
		start ,next = next, start + next
	}
	return result
}

//给定不同面额的硬币 coins 和一个总金额 amount。
//编写一个函数来计算可以凑成总金额所需的最少的硬币个数。
//如果没有任何一种硬币组合能组成总金额，返回 -1。

// 1到amount，维护一个map，能达到就是其count数，不能达到就是-1。
func coinChange(coins []int, amount int) int {
	count := make([]int, amount+1)
	for i:=1; i<=amount; i++ {
		count[i] = -1
		for _,c := range coins {
			if i == c {
				count[i] = 1
			}
			if (i-c) >=0 && count[i-c] != -1 {
				if (count[i-c] + 1) < count[i] || count[i] == -1{
					count[i] = count[i-c] + 1
				}
			}
		}
	}
	return count[amount]
}







func MyCoin(coins []int, target int) int {
	targetRecord := make([]int, target + 1)
	for i :=1; i<=target; i++ {
		targetRecord[i] = -1
		for _,c := range coins {
			if i - c < 0  || targetRecord[i - c] == -1{
				continue
			}
			if i - c == 0 {
				targetRecord[i] = 1
				continue
			}
			current := targetRecord[i-c] + 1
			if targetRecord[i] == -1  || current < targetRecord[i]{
				targetRecord[i] = current
			}
		}
	}
	return targetRecord[target]
}



func main() {
	coins := []int{2,3,6}
	change := MyCoin(coins, 15)
	fmt.Println(change)
}