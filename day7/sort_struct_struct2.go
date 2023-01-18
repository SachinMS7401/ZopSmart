package main

import (
	"fmt"
	"sort"
)

type person1 struct {
	name string
	age  int
}

type people struct {
	sliceperson []person1
}

func (a *people) sorting() {
	sort.Slice(a.sliceperson, func(p, q int) bool {
		return a.sliceperson[p].age < a.sliceperson[q].age
	})
}

func (a *people) adding() {
	a.sliceperson = append(a.sliceperson, person1{"sac", 21}, person1{"tanu", 19}, person1{"ajay", 20}, person1{"karthik", 18}, person1{"mahadev", 20})
}
func main() {
	//p1 := person1{"sac", 21}
	//p2 := person1{"tanu", 19}
	//p3 := person1{"ajay", 20}
	//p4 := person1{"karthik", 18}
	//p5 := person1{"mahadev", 18}
	p := []person1{}
	g := people{p}
	g.adding()
	fmt.Println(g)
	g.sorting()
	fmt.Println(g)

}
