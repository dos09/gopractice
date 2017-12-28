package basics

import (
	"fmt"
)

func RunIfSwitch() { //with capital letter to be exported
	fmt.Println(" * Run_if_switch")
	fmt.Println(sum(5))
	fmt.Println(sum2(5))
	practiceIfs()
	practiceSwitch()
	fmt.Println("finished execution")
}

func sum(limit int) int {
	sum := 0
	for i := 0; i <= limit; i++ {
		sum += i
	}
	return sum
	
	//infinite loop
//	for {
//	}
}

func sum2(limit int) int {
	sum := 0
	i := 0
	//for loop without init and end steps, only condition step
	for i <= limit {
		sum += i
		i++
	}
	return sum
}

func practiceIfs() {
	fmt.Println("practice_ifs")
	a := 1
	if a == 1 {
		fmt.Println("a == 1")
	}
	if b := 4; b > a { //allows short statement before the condition
		fmt.Println("b is not visible after the if")
	} else {
		//b is visible here
		fmt.Println("empty")
	}

	if 1 == 2 {
		fmt.Println(1)
	} else if 1 == 1 {
		fmt.Println(2)
	}
}

func practiceSwitch() {
	fmt.Println("practice_switch")
	a := 1
	res := ""
	switch a {
	case 1:
		res = "a 1"
	default:
		res = "a default"

	}
	fmt.Println(res)
	switch b := 4; b {
	case 1:
		res = "b 1"
	default:
		res = "b default"

	}
	fmt.Println(res)
	a = 100
	switch {
	case a < 0:
		res = "is negative"
	case a == 0:
		res = "zero"
	case isPositive(a):
		res = "is positive"
	}
	fmt.Println(res)
}

func isPositive(a int) bool {
	return a > 0
}
