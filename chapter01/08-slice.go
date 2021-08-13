package main

import "fmt"

func main() {
	// 基本使用
	months := [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	numbers := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	firstNumHalf := numbers[6:]
	fmt.Println(len(firstNumHalf)) //6
	fmt.Println(cap(firstNumHalf)) //12
	fmt.Println("firstNumHalf", firstNumHalf)
	q2 := months[3:6]     // 第二季度
	summer := months[5:8] // 夏季
	all := months[:]
	firsthalf := months[:6]
	secondhalf := months[6:]
	q1 := firsthalf[:3]     // 基于 firsthalf 的前 3 个元素构建新切片
	fmt.Println("all", all) // [January February March April May June July August September October November December]

	fmt.Println("firsthalf", firsthalf) //[January February March April May June]

	fmt.Println("secondhalf", secondhalf) //[July August September October November December]

	fmt.Println("q2", q2) //q2 [April May June]

	fmt.Println(len(q2))          // 3
	fmt.Println(cap(q2))          // 9
	fmt.Println("summer", summer) //summer [June July August]

	fmt.Println("q1", q1) //q1 [January February March]

	fmt.Println(len(firsthalf)) //6
	fmt.Println(cap(firsthalf)) //12

	fmt.Println(len(q1)) //3
	fmt.Println(cap(q1)) //12

	//直接创建
	mySlice1 := make([]int, 5, 10)
	fmt.Println("mySlice1", mySlice1) //12
	fmt.Println(len(mySlice1))        //5
	fmt.Println(cap(mySlice1))        //10
	mySlice2 := []int{1, 2, 3, 4, 5}
	printSlice(mySlice2, "mySlice2")
	fmt.Println("mySlice2", mySlice1) //12
	fmt.Println(len(mySlice1))        //5
	fmt.Println(cap(mySlice1))        //10
	//遍历切片
	for i := 0; i < len(firsthalf); i++ {
		fmt.Println("firsthalf[", i, "] =", firsthalf[i])
	}
	for key, val := range firsthalf {
		fmt.Println("firsthalf[", key, "] =", val)
	}
	//动态增加元素
	oldSlice := make([]int, 5, 10)
	fmt.Println("len(oldSlice):", len(oldSlice))
	fmt.Println("cap(oldSlice):", cap(oldSlice))
	appendSlice := []int{1, 2, 3, 4, 5}
	newSlice := append(oldSlice, appendSlice...)
	printSlice(newSlice, "newSlice") //name=newSlice  len=10 cap=10 0xc000018140 [0 0 0 0 0 1 2 3 4 5]

	//自动扩容
	appendSlice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	newSlice1 := append(oldSlice, appendSlice1...)
	printSlice(newSlice1, "newSlice1") //name=newSlice1  len=14 cap=20 0xc00010e000 [0 0 0 0 0 1 2 3 4 5 6 7 8 9]

	// 内容复制
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}

	// 复制 slice1 到 slice 2
	copy(slice2, slice1)         // 只会复制 slice1 的前 3 个元素到 slice2 中
	printSlice(slice2, "slice2") //name=slice2  len=3 cap=3 0xc000016140 [1 2 3]

	// 复制 slice2 到 slice 1
	copy(slice1, slice2)         // 只会复制 slice2 的 3 个元素到 slice1 的前 3 个位置
	printSlice(slice1, "slice1") //name=slice1  len=5 cap=5 0xc00001c0c0 [1 2 3 4 5]
	//动态删除元素

	// 数据共享
	//slice1 := make([]int, 4, 5)
	//slice2 := slice1[1:3]
	//slice1 = append(slice1, 0)
	//slice1[1] = 2
	//slice2[1] = 6
	//
	//fmt.Println("slice1:", slice1)
	//fmt.Println("slice2:", slice2)
}

func printSlice(s []int, name string) {
	fmt.Printf("name=%s  len=%d cap=%d %p %v\n", name, len(s), cap(s), s, s)
}
