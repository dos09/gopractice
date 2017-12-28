package basics

import (
	"fmt"
)

func RunInterface() {
	fmt.Println(" * RunInterface")
	practiceInterface1()
	practiceInterface2()
	
}

type IPoint interface {
	Print()
	GetInt() int
}

type Point3 struct {
	x, y float64
}

func (p *Point3) Print() {
	fmt.Println("Print: Point")
}

func (P *Point3) GetInt() int {
	return 1
}

func practiceInterface1() {
	fmt.Println("practiceInterface1")
	p := Point3{x: 1, y: 2}
	var ip IPoint = &p
	ip.Print()
	fmt.Println(ip.GetInt())
}

func practiceInterface2() {
	fmt.Println("practiceInterface2")
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}