package main

import (
	"fmt"
)

//编写一个函数，用切片的形式呈现斐波那契数列
//函数传参为一个整数n int，返回一个切片 slice []uint64
func fbn(n int) []uint64 {
	var slice []uint64 = make([]uint64, n) //定义一个切片存放接下来循环的斐波那契数
	slice[0] = 1                           //根据斐波那契特性得出第一个数和第二个数都为1
	slice[1] = 1
	for i := 2; i < n; i++ {
		slice[i] = slice[i-1] + slice[i-2] //当前数=上一个数+上两个数
	}
	return slice
}

func main() {

	fmt.Println(fbn(8))//[1 1 2 3 5 8 13 21]
}
