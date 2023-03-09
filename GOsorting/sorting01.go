package main

import (
	"fmt"
	"sort"
)

func main () {
	strs := []string{"r", "z", "f", "e", "e", "s", "e", "r"}

	sort.Strings(strs)

	fmt.Println("Strings:", strs)

	ints := []int{2, 4, 6, 0, 1}

	sort.Ints(ints)

	fmt.Println("Ints:   ", ints)
	s := sort.IntsAreSorted(ints)
  fmt.Println("Sorted: ", s)

}