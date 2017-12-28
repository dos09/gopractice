package basics

import (
	"fmt"
)

func RunFuncWithReceiver() {
	fmt.Println("RunFuncWithReceiver")
	runFuncWithPoint2AsReceiver()
	runFuncWithMyIntAsReceiver()
}

func runFuncWithPoint2AsReceiver() {
	p := Point2{x:1, y:2}
	p = p.IncXY()
	fmt.Printf("p = %v\n", p)
}

func runFuncWithMyIntAsReceiver() {
	var m MyInt = 10
	m = m.IncMyInt()
	fmt.Printf("m = %v\n", m)	
}

type Point2 struct {
	x, y float64
}

func (p Point2) IncXY() Point2 {
	p.x += 1
	p.y += 1
	return p
}

type MyInt int

func (m MyInt) IncMyInt() MyInt {
	return m + 1
}