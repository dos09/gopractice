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

	fmt.Println(" * Finished")
}

func run() {
	m := make(map[string]int)
	fmt.Println(m["Asen"])
}

func runBasics() {
	fmt.Println(basics.Pi)
	basics.RunIfSwitch()
	basics.RunDataTypesConst()
	basics.RunVariablesDeclarationsFunctions()
	basics.RunStructArrayMap()
	basics.RunFuncAsParameterClosure()
	basics.RunFuncWithReceiver()
}
