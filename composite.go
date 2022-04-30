package main

import "fmt"

func main() {
	fmt.Println("Hello, goLang")
	// 4.1 array
	// var arr [3]int = [3]int{1, 3, 4}
	// arr := [3]int{1, 3, 4}
	arr := [...]int{1, 3, 4}
	for i, ele := range arr {
		fmt.Printf("index: %d, value: %d\n", i, ele)
	}

	// to learn how to use iota in Golang, please visit
	// https://yourbasic.org/golang/iota/
	// The iota keyword represents successive integer constants 0, 1, 2,
	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "!", GBP: "@", RMB: "&"}
	fmt.Println(RMB, symbol[RMB])

	// 4.2 Slice
	a := [...]int{0, 1, 2, 3, 4, 5}
	// reverse(a[:])
	// fmt.Println(a)
	reverse(a[2:])
	fmt.Println(a)

}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
