package main

import "fmt"

func main() {

	//数组
	var p1 = [5]int{5, 4, 12, 3, 2}
	fmt.Println(p1)
	//切片slice
	var s1 = []int{1, 2, 4, 5}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(s1), cap(s1), s1)

	//截取切片
	fmt.Println("s1[1]=", s1[1])
	fmt.Println("s1[:]=", s1[:])
	fmt.Println("s1[1:]=", s1[1:])
	fmt.Println("s1[:3]=", s1[:3])
}
