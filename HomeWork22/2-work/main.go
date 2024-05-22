package main

import "fmt"

func main() {
	n := []int{1, 2, 3, 4, 5, 7, 6, 9}
	Printer(n)
}

func Printer(n []int) {
	for i := 0; i < len(n); i++ {
		fmt.Println(n[i])
	}
}
