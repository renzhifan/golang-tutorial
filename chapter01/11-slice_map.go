package main

import (
	"fmt"
	"sort"
)

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
)

func main() {
	items := SliceMap()
	fmt.Println(items)
	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}
	//将 map 的键值对调
	invMap := make(map[int]string, len(barVal))
	for k, v := range barVal {
		invMap[v] = k
	}
	fmt.Println("inverted:")
	intKeys := make([]int, len(invMap))
	j := 0
	for k, _ := range invMap {
		intKeys[j] = k
		j++
	}
	fmt.Println(intKeys)
	sort.Ints(intKeys)
	fmt.Println(intKeys)
	vintMap := make(map[string]int, len(barVal))

	for _, v := range intKeys {
		vintMap[invMap[v]] = v
		//fmt.Printf("Key: %v, Value: %v / ", v, invMap[v])
	}
	fmt.Println(vintMap)
}

func SliceMap() []map[string]string {
	items := make([]map[string]string, 1)
	for i := range items {
		items[i] = make(map[string]string)
		items[i]["name"] = "test"
		items[i]["age"] = "11"
	}
	return items
}
