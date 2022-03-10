package main

import "fmt"

func main() {
	/*
		用二维数组打印如下效果：
		1 0 0 0 0 0
		0 2 0 0 0 0
		0 0 3 0 0 0
		0 0 0 4 0 0
		0 0 0 0 5 0
		0 0 0 0 0 6
	*/

	//定义一个二维数组
	var multiArray = [6][6]int{}
	// multiArray[0][0] = 1
	// multiArray[1][1] = 2
	for i := 0; i < len(multiArray); i++ {
		for k := 0; k < len(multiArray[i]); k++ {
			multiArray[i][i] = i + 1
			fmt.Print(multiArray[i][k], " ")
		}
		fmt.Println()
	}
}
