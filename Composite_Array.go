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
	fmt.Println("test")

	// 4.2.1 append function
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap = %d\t%v\n", i, cap(y), y)
		x = y
	}

	// 4.2.2 in-place slice techniques
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2))
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// 4.2.1 The append function
func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// there is room to grow, extend the slice
		z = x[:zlen]
	} else {
		// there is in sufficient space, allocate a new array
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
