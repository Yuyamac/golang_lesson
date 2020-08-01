package main

import "fmt"

func main()  {
	b := []byte{72, 73}
	fmt.Println(b)
	// ASCIIコードを確認
	// https://www.ascii-code.com/
	fmt.Println(string(b))

	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))
}