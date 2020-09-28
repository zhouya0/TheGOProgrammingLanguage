package main

import (
	"fmt"
)

var testNumber = []int{3,3,3,3,3,3,5,54,6453,65,65,3654,6435,76,4,4,4,4,4,4,4,4,4,4,4}

var twodimension = [][]int{{4,5,6,7,8},{20,21,23,25}, {78,90,91,100} }

func findMostNumber(list []int) int {
	numberMap := make(map[int]int)
	for _,v := range list {
		numberMap[v] = numberMap[v] + 1
	}
	MaxNumber := 0
	var result int
	for k,v := range numberMap {
		if v >= MaxNumber {
			MaxNumber = v
			result = k
		}
	}
	return result
}

func findThree(list []int, target int) [][]int {
	var result [][]int
	for k,v := range list {
		if k == len(list) - 2{
			break
		}
		twoResult := findPaires(list[k+1:], target-v)

		for _,j := range twoResult {
			result = append(result, append(j, v))
		}
	}
	return result
}

func findPaires(list []int, target int) [][]int {
	numMap := make(map[int]interface{})
	var result [][]int
	for _,v := range list {
		if _,ok := numMap[v]; ok {
			result = append(result, []int{v, target-v})
		} else {
			numMap[target - v] = nil
		}
	}
	return result
}

func findKNumber(list []int, k int) []int{
	return nil
}

func babooSort(list []int) []int{
	length := len(list)
	if length == 1 {
		return list
	}
	for i:=0;i<length;i++{
		for j:=0;j<(length-i-1);j++{
			if list[j] > list[j+1] {
				list[j+1],list[j] = list[j],list[j+1]
			}
		}
	}
	return list
}

func selectSort(list []int) []int{
	length := len(list)
	for i:=0; i<length;i++{
		for j:=i; j<length;j++{
			if list[j] > list[i]{
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	return list
}

var listtest1  = []int{3,67,1234,54666,65437573}
var listtest2 = []int{1,5,7,8,11,15}

func mergeSort(list1 []int, list2 []int) []int {
	var result []int
	length1 := len(list1)
	length2 := len(list2)
	var i int
	var j int
	for  {
		if list1[i] < list2[j] {
			result = append(result, list1[i])
			i ++
		} else {
			result = append(result, list2[i])
			j ++
		}
		if i == length1 -1 {
			result = append(result, list2[j:]...)
			break
		}
		if j == length2 -1 {
			result = append(result, list1[i:]...)
			break
		}
	}
	return result
}


func quickSort(list []int) {
	if len(list) <= 1{
		return
	}
	fmt.Println(list)
	newindex := partition(list)
	quickSort(list[:newindex])
	quickSort(list[newindex+1:])
}

func partition(list []int) int{
	start := 0
	end := len(list)-1
	temp := list[0]
	for start < end {
		for start < end && temp>=list[start] {
			start ++
		}
		for start < end && temp<=list[end] {
			end --
		}
		if start < end {
			list[start], list[end] = list[end], list[start]
		}
	}

	if start > 0 {
		list[0], list[start-1] = list[start-1], list[0]
		return start -1
	}
	return 0
}

func findTwoDimensionNum(matrix [][]int, target int)bool {
	var firstcolumn []int
	for _,v := range matrix {
		firstcolumn = append(firstcolumn, v[0])
	}
	var colume int
	var found bool
	for k,v := range firstcolumn {
		if k<(len(firstcolumn)-1) && target >= v && target <firstcolumn[k+1] {
			colume = k
			found = true
			break
		}
		if k==(len(firstcolumn)-1) && target >= v {
			colume = k
			found = true
		}
	}
	if !found {
		return false
	}
	for _,v := range matrix[colume] {
		if v == target {
			return true
		}
	}
	return false
}

func ThreeNums(list []int, target int) [][]int {
	list = sortFunc(list)

	var result [][]int
	for k,v := range list {
		if k == len(list) - 3{
			if (v + list[k+1] + list[k+2]) ==target {
				result = append(result,[]int{v, list[k+1], list[k+2]})
			}
			break
		}
		start := k+1
		end := len(list) -1
		for start < end {
			cur := v + list[start] + list[end]
			if cur > target{
				end --
			} else if cur < target {
				start ++
			} else {
				result = append(result, []int{v, list[start], list[end]})
				end --
				start ++
			}

		}
	}
	return result
}


func sortFunc(list []int) []int {
	length := len(list)
	for i:=0; i<length-1;i++ {
		for j:=0;j<length-i-2;j++{
			if list[j] > list[j+1]{
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list
}


func main() {
	fmt.Println(mergeSort(listtest1, listtest2))
}