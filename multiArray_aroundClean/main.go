package main

import "fmt"

/*
定义一个三行四列的二维数组，逐个从键盘输入值，编写程序将四周的值清零
思路:将二维数组遍历出来，观察规律
1.先将每个子数组的第一个元素和最后一个元素变为0
1 1 1 1       0 1 1 0
1 1 1 1  -->> 0 1 1 0
1 1 1 1       0 1 1 0   arr[i][0]=0  arr[i][len(arr[i])-1]=0

2.再将整个二维数组的第一个子数组的全部元素和最后一个子数组的全部元素变为0
1 1 1 1       0 1 1 0        0 0 0 0
1 1 1 1  -->> 0 1 1 0   -->> 0 1 1 0
1 1 1 1       0 1 1 0        0 0 0 0
如果i=0（代表第一个子数组）或者i=len(arr)-1(代表最后一个子数组)，那么arr[i][k]=0
*/

func main() {
	var multiArr = [3][4]int{}

	for i := 0; i < len(multiArr); i++ {
		for k := 0; k < len(multiArr[i]); k++ {
			fmt.Printf("请输入第%d行第%d个数:\n", i+1, k+1)
			fmt.Scanln(&multiArr[i][k])
		}
	}
	fmt.Println()
	for i := 0; i < len(multiArr); i++ {
		for k := 0; k < len(multiArr[i]); k++ {
			multiArr[i][0] = 0
			multiArr[i][len(multiArr[i])-1] = 0
			if i == 0 || i == len(multiArr)-1 {
				multiArr[i][k] = 0
			}
			fmt.Printf("%v", multiArr[i][k])
		}
		fmt.Println()
	}
}
