package main

import (
	"fmt"
)

/*
已知有个排序好的数组（升序），要求插入一个元素，插入后同样是升序排序的，最后打印该数组
*/

func work(sliceArr []int, num int) []int {
	sliceArr = append(sliceArr, num) //将要插入的值插入原切片

	//冒泡排序
	temp := 0
	for j := 0; j < len(sliceArr)-1; j++ {
		for i := 0; i < len(sliceArr)-1-j; i++ {
			if sliceArr[i+1] < sliceArr[i] {
				temp = sliceArr[i]
				sliceArr[i] = sliceArr[i+1]
				sliceArr[i+1] = temp
			}
		}
	}
	return sliceArr
}

func main() {
	//定义一个无序数组并将其转为切片
	var arr2 = [...]int{12, 89, 54, 6, 21, 97}
	sliceArr2 := arr2[:]
	fmt.Println(sliceArr2)

	var appendNum int
	fmt.Println("请输入要插入的数:")
	fmt.Scanln(&appendNum)
	fmt.Println(work(sliceArr2, appendNum)) //传入切片和要插入的数
}
