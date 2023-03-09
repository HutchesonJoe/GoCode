package main

import (
    "fmt"
    "sort"
)

// record a "Person" Name, Age
type Person struct {
    Name   string
    Age    int
		Height int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type ByAge []Person

type ByHeight[]Person

func (a ByAge) Len() int {
	return len(a)
}
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func (a ByHeight) Len() int {
	return len(a)
}
func (a ByHeight) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByHeight) Less(i, j int) bool {
	return a[i].Height < a[j].Age
}

func main() {

	// Name, Age
	people := []Person{
			{"Bob", 31, 6},
			{"John", 42, 6},
			{"Michael", 17, 5},
			{"Jenny", 26, 6},
	}

	fmt.Println(people)

	sort.Sort(ByAge(people))
  fmt.Println(people)

	sort.Sort(ByHeight(people))
	fmt.Println(people)
}

