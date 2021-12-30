package main

import "fmt"

func main() {
	//
	//var v int8 = 127
	//fmt.Println(v)

	array()

}

func array() {

	var a [3]int
	for i, v := range a {
		fmt.Println(i, v)
	}

	//数组字面量初始化
	c := [...]int{1, 2, 3, 4, 5}
	fmt.Println(c)
	//c2 := append(c, 6) //xx

	//slice 写法 []T
	b := []int{1, 2, 3, 4, 5}
	fmt.Println(b)
	b2 := append(b, 6)
	fmt.Println(b2)

	//map
	e := map[string]int{"a": 1, "b": 2}
	fmt.Println(e)

	//struct 结构体字面量
	type Point struct{ x, y int }
	//结构体定义
	type element struct {
		ID   int
		name string
	}

}
