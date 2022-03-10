package bubble_sort

import(
	_"fmt"
)


//定义一个传入切片的函数，切片本身为引用类型，遵循引用传递规则，长度可以动态改变，使用起来更灵活方便
func Bubble_sort(arr []int) {
	a := 0 //临时变量a用于数组元素的交换
	for j := 0; j < len(arr)-1; j++ {
		//根据规律得出轮数等于总长度减1
		for i := 0; i < len(arr)-1-j; i++ {
			//根据规律得出第一轮交换的次数=数组的总长度减1再减去当前轮数
			if arr[i] > arr[i+1] { //如果前一个大于下一个
				a = arr[i]        //先将a赋值为数组的前一个元素
				arr[i] = arr[i+1] //前一个元素赋值为后一个元素
				arr[i+1] = a      //后一个元素赋值为a,a=前一个元素
			}
		}
	}
}

// func main() {
// 	arr01 := [...]int{89, 65, 12, 45, 33, 78, 99, 9, 67, 28} //声明一个无序数组
// 	arr01_slice := arr01[:]                                  //将无序数组转为切片
// 	Bubble_sort(arr01_slice)                                 //传入切片
// 	fmt.Println(arr01_slice)                                 //[9 12 28 33 45 65 67 78 89 99]
// }
