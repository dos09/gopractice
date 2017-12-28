package basics

import (
	"fmt"
)

func RunStructArrayMap() {
	practicePointer()
	practiceStruct()
	practiceArrays()
	practiceMaps()
}

func practicePointer() {
	fmt.Println("practice_pointer")
	num := 4
	p := &num //p contains the address of num
	fmt.Println(p)
	fmt.Println(*p, num)
	*p = 100 //change the value of num
	fmt.Println(*p, num)
}

type Point struct {
	x int
	y int
}

func practiceStruct() {
	fmt.Println("practice_struct")
	point := Point{1, 2}
	fmt.Println(point.x, point.y)

	ptr := &point
	fmt.Println(ptr.x, ptr.y)
	fmt.Println((*ptr).x, (*ptr).y)

	point2 := Point{x: 4} //y is implicitly 0
	fmt.Println(point2.x, point2.y)
}

func practiceArrays() {
	fmt.Println("practice_arrays")
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[1:4]
	fmt.Println(s)
	s[0] = 123
	fmt.Println(primes)

	arr := make([]int, 5)
	fmt.Printf("arr: %d, len: %d, capacity: %d\n", arr, len(arr), cap(arr))

	var myArr [3]int
	myArr[0] = 1
	fmt.Printf("myArr: %d, len: %d, capacity: %d\n", myArr, len(myArr), cap(myArr))

	var b []int
	b = append(b, 1, 2, 3) //used as dynamic array
	fmt.Println("b:", b)

	fmt.Println("range")
	c := []int{1, 2, 4, 8, 16}
	for index, value := range c {
		fmt.Println(index, value)
	}
	for _, value := range c {
		fmt.Println(value)
	}
	
	ta := make([]int, 1, 5) // ta: [0], len(ta): 1, cap(ta): 5
	fmt.Printf("ta: %v, len(ta): %d, cap(ta): %d\n", ta, len(ta), cap(ta))
	
	m := [][]int{
		[]int{0,0,0},
		[]int{1,1,1},
		[]int{2,2,2}, //LAST COMMA IS REQUIRED ...
	}
	fmt.Println(m)
}

func practiceMaps() {
	fmt.Println("practice_maps")
	orcKills := make(map[string]int)
	orcKills["Ra"] = 100
	fmt.Println(orcKills) 
	fmt.Printf("Ra has %d kills\n", orcKills["Ra"])
	kills, exists := orcKills["Mogka"]
	if exists {
		fmt.Printf("Mogka has %d kills\n", kills)
	} else {
		fmt.Printf("Mogka is not in orcKills\n")
	}
	
	fmt.Println(orcKills["Mogka"]) //when the key doesn't exist the zero value
	//for the value's data type is returned
	
	m := map[string]int {
		"Asen":30,
		"Ivan":12, //LAST COMMA IS REQUIRED ...
	}
	fmt.Println(m)
}