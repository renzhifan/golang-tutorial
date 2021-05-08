package main

import (
	"fmt"
	"sort"
)

func main() {
	var testMap map[string]int
	testMap = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	k := "two"
	v, ok := testMap[k]
	if ok {
		fmt.Printf("The element of key %q: %d\n", k, v)
	} else {
		fmt.Println("Not found!")
	}

	myMap := make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 3
	myMap["four"] = 4

	key := "four"
	value, findBool := myMap[key]
	if findBool {
		fmt.Printf("The element of key %q: %d\n", key, value)
	}
	delete(myMap, "four")

	fmt.Println(testMap)
	fmt.Println(myMap)
	//遍历字典

	for key, val := range myMap {
		fmt.Println(key, val)
	}

	//键值对调
	invMap := make(map[int]string, 3)
	for k, v := range myMap {
		invMap[v] = k
	}
	fmt.Println(invMap)

	//字典排序

	//按照键进行排序

	keys := make([]string,0)

	for k,_:=range myMap{
		keys=append(keys,k)
	}
	sort.Strings(keys)
	fmt.Println(keys)
	for _,v:=range keys{
		fmt.Println(v,myMap[v])
	}
}
