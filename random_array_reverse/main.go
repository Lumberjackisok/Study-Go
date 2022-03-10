package main

//随机生成十个整数保存到数组，并倒序打印，以及求出平均值，最大值，最大值的下标

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var arr = [10]int{}              //声明一个空数组用于存放随机数
	rand.Seed(time.Now().UnixNano()) //给随机数做种

	for i := 0; i < len(arr); i++ { //遍历数组
		ranNum := rand.Intn(100) + 1
		arr[i] = ranNum //给每一个下标赋值
	}
	fmt.Println(arr)

	//数组倒序，求平均值，最大值及其下标

	var newArr [10]int   //声明一个新的空数组，用于倒序存放原数组的数
	var sum int          //用于求和后求平均值
	var max int = arr[0] //假定最大值的下标为第一个元素
	index := 0           //用于交换下标
	k := len(arr)        //用于在遍历中递减

	//遍历数组
	for i := 0; i < len(arr); i++ {
		k--
		sum += arr[i]
		if max < arr[i] { //如果下一个元素大于上一个元素则替换
			max = arr[i]
			index = i
		}
		newArr[k] = arr[i]
	}
	arr = newArr
	fmt.Println("倒序后为:", arr)
	fmt.Println("平均值为:", sum/len(arr))
	fmt.Println("最大值及其下标为:", max, index)
}
