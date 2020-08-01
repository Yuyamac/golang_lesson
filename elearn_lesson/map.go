package main

import "fmt"
// map
func main()  {
	// Pythonの辞書型に近い
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["banana"] = 300
	fmt.Println(m)
	m["new"] = 500
	fmt.Println(m)

	// 中身がない場合: 0が返ってくる
	fmt.Println(m["nothing"])

	// 中身がある場合: valueとtrueが返ってくる
	v, ok := m["apple"]
	fmt.Println(v, ok)

	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)

	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)

	/*
	var m3 map[string]int
	m3["pc"] = 5000
	fmt.Println(m3)
	 */

	var s []int
	if s == nil{
		fmt.Println("Nil")
	}
}