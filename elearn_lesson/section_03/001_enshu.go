package main

import "fmt"

// 1.1
//func main()  {
//	f := 1.11
//	i := int(f)
//	fmt.Printf("%T %v\n", i, i)
//}

// 1.2
//func main()  {
//	s := []int{1,2,5,6,2,3,1}
//	fmt.Println(s[2:4])
//}

// 1.3
func main()  {
	m := map[string]int{"Mike": 20,"Nancy": 24, "Messi": 30}
	fmt.Printf("%T %v", m, m)
}
