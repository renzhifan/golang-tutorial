package main

import "fmt"

func main() {
	// 基本使用
	months := [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	q2 := months[3:6]     // 第二季度
	summer := months[5:8] // 夏季
	all := months[:]
	firsthalf := months[:6]
	secondhalf := months[6:]
	q1 := firsthalf[:3] // 基于 firsthalf 的前 3 个元素构建新切片
	fmt.Println("all", all)
	fmt.Println("firsthalf", firsthalf)
	fmt.Println("secondhalf", secondhalf)
	fmt.Println("q2", q2) //q2 [April May June]

	fmt.Println(len(q2))          // 3
	fmt.Println(cap(q2))          // 9
	fmt.Println("summer", summer) //summer [June July August]

	fmt.Println("q1", q1) //q1 [January February March]

	fmt.Println(len(q1)) //3
	fmt.Println(cap(q1)) //12

	mySlice1 := make([]int ,5,10)
	fmt.Println("mySlice1",mySlice1) //12
	fmt.Println(len(mySlice1)) //5
	fmt.Println(cap(mySlice1)) //10
	mySlice2 := []int{1,2,3,4,5}
	printSlice(mySlice2,"mySlice2")
	fmt.Println("mySlice2",mySlice1) //12
	fmt.Println(len(mySlice1)) //5
	fmt.Println(cap(mySlice1)) //10
	// 切片复制
	//slice1 := []int{1, 2, 3, 4, 5}
	//slice2 := []int{5, 4, 3}

	// 复制 slice1 到 slice 2
	//copy(slice2, slice1) // 只会复制 slice1 的前 3 个元素到 slice2 中
	//fmt.Println(slice2)
	// 复制 slice2 到 slice 1
	//copy(slice1, slice2) // 只会复制 slice2 的 3 个元素到 slice1 的前 3 个位置
	//fmt.Println(slice1)

	// 数据共享
	slice1 := make([]int, 4, 5)
	slice2 := slice1[1:3]
	slice1 = append(slice1, 0)
	slice1[1] = 2
	slice2[1] = 6

	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)
}

func printSlice(s []int,name string) {
	fmt.Printf("name=%s  len=%d cap=%d %p %v\n",name, len(s), cap(s), s, s)
}