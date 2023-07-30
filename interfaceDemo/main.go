package main

import (
	"fmt"
	"reflect"
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
