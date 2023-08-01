package main

import (
	"fmt"
	"reflect"
	"sort"
)

type Reason uint

const (
	SPRING Reason = iota + 1
	SUMMER
	AUTUMN
	WINTER
)

func main() {
	fmt.Println(SPRING)
	fmt.Println(SUMMER)
	fmt.Println(AUTUMN)
	fmt.Println(WINTER)
	fmt.Println(reflect.TypeOf(WINTER))
	students := []Student{
		{ID: 1, Name: "A", Score: 88},
		{ID: 2, Name: "B", Score: 90},
		{ID: 3, Name: "C", Score: 78},
		{4, "D", 99},
	}
	sort.Sort(Students(students))
	fmt.Println(students)
}

func (r Reason) String() string {
	return ReasonText[r]
}

var ReasonText = map[Reason]string{
	SPRING: "春天",
	SUMMER: "夏天",
	AUTUMN: "秋天",
	WINTER: "冬天",
}
