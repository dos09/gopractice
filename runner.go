package main

import (
	"fmt"
	"gopractice/basics"
	"gopractice/concurrency"
	"gopractice/excercises"
)

//from parent dir of /src dir:
//go install gopractice && bin\gopractice.go

func main() {
	fmt.Println(" * Started")

	excercises.RunExerciseWebCrawler()

	fmt.Println("\n * Finished")
}

func run() {
	fmt.Println(3 / 2.)
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
	basics.RunDeferPanicRecover()
}

func runConcurrency() {
	concurrency.RunChannel()
	concurrency.RunSelect()
	concurrency.RunDefaultSelect()
	concurrency.RunMutex()
}

func runExcercises() {
	excercises.RunExcerciseRot13Reader()
	excercises.RunExerciseWebCrawler()
}
