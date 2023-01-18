package main

import "fmt"

func main() {
	m := make([]int, 0)
	var n map[struct{}]int
	if n == nil {
		fmt.Println("it is a nil")
	}
	fmt.Println(m)
}
