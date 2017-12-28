package basics

import (
	"fmt"
)

func RunFuncAsParameterClosure() {
	funcAsParameter()
	funcAsClosure()
}

func funcAsParameter() {
	fmt.Println("funcAsParameter")
	fmt.Println(calc(add2, 1, 4))
}

func add2(x, y int) int { // already exists in another .go file
	return x + y
}

func calc(f func(int, int) int, op1 int, op2 int) int {
	return f(op1, op2)
}

func accumulate() func(int) int {
	acc := 0
	return func(val int) int {
		acc += val
		return acc
	}
}

func funcAsClosure() {
	fmt.Println("funcAsClosure")
	fAcc := accumulate()
	fmt.Println(fAcc(4))
	fmt.Println(fAcc(10))
}