package main

import "fmt"

func hello(name string) string {
	msg := fmt.Sprintf("Hi, %v. Welcome!", name)
	return msg
}
