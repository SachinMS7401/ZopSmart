package main

import (
	"fmt"
	"sort"
)

func main() {
	Person := []struct {
		name string
		age  int
	}{
		{"sac", 21},
		{"tanu", 19},
		{"ajay", 20},
		{"mahadev", 20},
		{"karthik", 18},
	}
	sort.Slice(Person, func(p, q int) bool {
		return Person[p].age < Person[q].age
	})
	fmt.Println(Person)

	sort.Slice(Person, func(p, q int) bool {
		return Person[p].name < Person[q].name
	})
	fmt.Println(Person)
}
