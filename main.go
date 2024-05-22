package main

import "fmt"

func main() {
	n := "Salom"
	Printer(n)
}

func Printer(n string) {
	for i := 0; i < len(n); i++ {
		fmt.Println(i)
	}
}
