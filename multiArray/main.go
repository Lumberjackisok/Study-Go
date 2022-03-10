package main

import "fmt"

//二维数组统计三个班的成绩，每个班五个人
//求出每个班的总成绩，平均成绩;所有班的总成绩，平均成绩

func Score() {
	var score = [3][5]float32{}
	for i := 0; i < len(score); i++ {
		for k := 0; k < len(score[i]); k++ {
			fmt.Printf("请输入第%d班第%d个学生的成绩:\n", i+1, k+1)
			fmt.Scanln(&score[i][k])
		}
	}
	fmt.Println(score)

	//遍历成绩
	var sum float32
	for i := 0; i < len(score); i++ {
		//声明一个循环体内的临时变量，目的是进入下一轮循环时初始值为零，用于储存每个班的总成绩
		var classSum float32
		for k := 0; k < len(score[i]); k++ {
			classSum += score[i][k]
			sum += score[i][k] //用循环体外部声明的临时变量将全部班级的成绩相加并存储
		}
		fmt.Printf("第%d班的总成绩为%v,平均成绩为%.1f\n", i+1, classSum, classSum/5)
		// fmt.Printf("第%d班的平均成绩为%.2f\n", i+1, classSum/5)
	}
	//fmt.Println("所有班级的总成绩为:", sum)
	fmt.Printf("所有班级的总成绩为%.2f,所有班级的平均成绩为%v\n:", sum, sum/(3*5))

}

func main() {
	Score()
}
