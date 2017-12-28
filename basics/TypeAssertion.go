package basics

//https://tour.golang.org/methods/15

import (
	"fmt"
)

func RunAssertion() {
	fmt.Println(" * RunAssertion")
	practiceAssertion1()
	practiceAssertion2()
}

func practiceAssertion1() {
	fmt.Println("practiceAssertion1")
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

//	f = i.(float64) // panic
//	fmt.Println(f)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func practiceAssertion2() {
	fmt.Println("practiceAssertion2")
	do(21)
	do("hello")
	do(true)
}