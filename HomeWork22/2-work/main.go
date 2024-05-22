package main

import "fmt"

func main() {
	n := 10
	Printer(n)
}

func Printer(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}
}
