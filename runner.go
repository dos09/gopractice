package main

import (
	"fmt"
	"gopractice/basics"
)

//go install gozhulien
//the exe is in bin folder

func main() {
	fmt.Println(" * Started")

	run()

	fmt.Println("\n * Finished")
}

func run() {
	fmt.Println(rot13(65))
	fmt.Println(rot13(66))
	fmt.Println(rot13(78)) 
}

func rot13(b byte) byte {
	var min, max, offset byte
	if b >= 65 && b <= 90 {
		min = 64
		max = 90
	} else if b >= 97 && b <= 122 {
		min = 96
		max = 122
	} else {
		return b
	}
	
	b = b + 13
	if b > max {
		offset = b - max
		b = min + offset
	}
	return b
}

func runBasics() {
	fmt.Println(basics.Pi)
	basics.RunIfSwitch()
	basics.RunDataTypesConst()
	basics.RunVariablesDeclarationsFunctions()
	basics.RunStructArrayMap()
	basics.RunFuncAsParameterClosure()
	basics.RunFuncWithReceiver()
	basics.RunArraySlice()
	basics.RunInterface()
	basics.RunAssertion()
	basics.RunRot13ReaderExcercise()
}
