package main

import "fmt"

/*
定义一个4行4列的二维数组，逐个从键盘输入值，
然后将第1行和第4行的数据进行交换，将第2行和第3行的数据进行交换
*/

func main() {

	var multiArray = [4][4]int{}
	//给二维数组挨个赋值
	for i := 0; i < len(multiArray); i++ {
		for k := 0; k < len(multiArray[i]); k++ {
			fmt.Printf("请输入第%d行第%d个元素:\n", i+1, k+1)
			fmt.Scanln(&multiArray[i][k])
		}
	}

	//把尚未进行交换的二维数组打印出来看看
	for i := 0; i < len(multiArray); i++ {
		for k := 0; k < len(multiArray[i]); k++ {
			fmt.Print(multiArray[i][k], " ")
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
	//-----------------------------分割线,以下部分为对二维数组内的子数组进行交换---------------------------------------
	//两个临时的用于交换的空数组
	swap1 := [len(multiArray[0])]int{}
	swap2 := [len(multiArray[1])]int{}

	//对二维数组内的子数组进行交换
	swap1 = multiArray[0]
	multiArray[0] = multiArray[3]
	multiArray[3] = swap1

	swap2 = multiArray[1]
	multiArray[1] = multiArray[2]
	multiArray[2] = swap2

	//将交换后的整个二维数组打印出来
	for i := 0; i < len(multiArray); i++ {
		for k := 0; k < len(multiArray[i]); k++ {
			fmt.Print(multiArray[i][k], " ")
		}
		fmt.Println()
	}

}
