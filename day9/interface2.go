package main

import "fmt"

type salarycal interface {
	calsalary() int
}

type Permanent struct {
	emp_id   int
	base_sal int
	pf       int
}

type Contract struct {
	emp_id   int
	base_sal int
}

func (p Permanent) calsalary() int {
	return p.base_sal + p.pf
}

func (c Contract) calsalary() int {
	return c.base_sal
}

func totalexpense(s []salarycal) {
	ex := 0
	for _, v := range s {
		ex += v.calsalary()
	}
	fmt.Printf("Total expenditure of a company is %d", ex)
	fmt.Println()
}

func main() {
	p1 := Permanent{1, 70000, 5000}
	p2 := Permanent{2, 65000, 4000}
	p3 := Permanent{3, 75000, 6000}
	c1 := Contract{4, 50000}
	c2 := Contract{5, 40000}
	Employees := []salarycal{p1, p2, p3, c1, c2}
	totalexpense(Employees)
}
