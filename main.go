package main

import (
	"fmt"
	"sort"
)
func main() {
	items := []int{ 31, 45, 63, 70, 100,1,2, 9, 20, 31, 45, 63, 70, 100}

	sort.Ints(items)
	fmt.Println(items)
}