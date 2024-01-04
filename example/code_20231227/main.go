package main

import (
	"example/code_20231227/func1"
	"fmt"
)

func main() {
	// function
	msg := func1.Hello("test")
	println(msg)

	// map
	ages := map[string]int{
		"Tester": 30,
		"Johnny": 20,
	}

	ages["Kyle"] = 24
	fmt.Println(ages["Kyle"])

	// array
	data := []int{1, 2, 3, 4, 5}
	for index, value := range data {
		fmt.Printf("%d: %d\n", index, value)
	}

	// 문자열 자르기
	for index, runeValue := range "hello" {
		fmt.Printf("%d: %c\n", index, runeValue)
	}
}
