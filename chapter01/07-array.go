package main

import "fmt"

func main() {
	//数组的定义格式
	//[capacity]data_type{element_values} [容量]数据类型{元素值}
	var a [8]byte           // 长度为8的数组，每个元素为一个字节
	var b [3][3]int         // 二维数组（9宫格）
	var c [3][3][3]float64  // 三维数组（立体的9宫格）
	var d = [3]int{1, 2, 3} // 声明时初始化
	var e = new([3]string)  // 通过 new 初始化
	f := [5]int{1, 2, 3, 4, 5}
	g := [5]int{1: 3, 2: 4} //初始化指定下标位置的元素值，未设置的位置也会以对应元素类型的零值填充 下标为1的值等于3 下标为2的值等于4
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
	fmt.Println("f:", f)
	fmt.Println("g:", g)
	// 数组的初始化
	a1 := [5]uint8{1, 2}
	fmt.Println("a1:", a1)

	a2 := [3]string{"hello", "world"}
	fmt.Println("a2:", a2)

	// 遍历数组
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println("Element", i, "of arr is", arr[i])
	}
	for key, val := range arr {
		fmt.Println("Element", key, "of arr is", val)
	}

	// 通过二维数组打印九九乘法表
	var multi [9][9]string
	for j := 0; j < 9; j++ {
		for i := 0; i < 9; i++ {
			n1 := i + 1
			n2 := j + 1
			if n1 < n2 { // 摒除重复的记录
				continue
			}
			multi[i][j] = fmt.Sprintf("%dx%d=%d", n2, n1, n1*n2)
		}
	}
	fmt.Printf("multi", multi)
	// 打印九九乘法表
	for _, v1 := range multi {
		for _, v2 := range v1 {
			fmt.Printf("%-8s", v2) // 位宽为8，左对齐
		}
		fmt.Println()
	}
}
