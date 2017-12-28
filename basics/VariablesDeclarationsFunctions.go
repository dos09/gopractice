package basics

import (
	"fmt"
)

var gg string = "this one is not in function"

func RunVariablesDeclarationsFunctions() { //with capital letter to be exported
	fmt.Println(" * Run_variable_declarations_functions")
	fmt.Println(gg)

	var a, b int = 1, 2
	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, sub(a, b))

	var c, d string           //declare variable
	c, d = swap("def", "abc") //init variable
	fmt.Println("swap", c, d)

	e, f := swap("345", "12") //declare and init variable, one way
	fmt.Println("swap", e, f)

	var g, h string = swap("1", "4") //declare and init variable, another way
	fmt.Println("swap", g, h)

	x, y := glupost()
	fmt.Println("glupost", x, y)

	var aa, bb = true, "text" //can omit the data type if declare+init
	fmt.Println("aa, bb", aa, bb)
	
	praticeDefer()
}

func add(a int, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func swap(a, b string) (string, string) {
	return b, a
}

func glupost() (x, y int) { //return named parameters
	x = 100
	y = 200
	return
}

func praticeDefer() {
	defer fmt.Println("world") //arguments are evaluated immediately, but the
	//function call is delayed
	//A defer statement defers the execution of a function until the surrounding function returns.
	fmt.Println("defer hello")
}
