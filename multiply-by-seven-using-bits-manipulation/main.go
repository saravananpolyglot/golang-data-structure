package main

import "fmt"

func multiplyBySeven(n int) int {
	return ((n << 3) - n)
}

func main() {
	fmt.Println(multiplyBySeven(10))
}
