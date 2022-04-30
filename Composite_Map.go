package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("hello composite maps")
	// ages := make(map[string]int)
	// ages["bob"] = 39
	// ages := map[string]int{
	// 	"bob": 20,
	// }
	ages := map[string]int{}
	ages["bob"] = 24
	ages["kim"] = 30
	ages["ame"] = 32
	// a map lookup using a key that isn't present returns zero value for its type
	// delete(ages, "bob")
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}
	// in order to numerate the ages in order, we can use sort.Strings function to get sorted keys in advance
	// var names []string
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}
