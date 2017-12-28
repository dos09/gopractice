package basics

import (
	"fmt"
)

func RunArraySlice() {
	fmt.Println(" * RunArraySlice")
	practiceArraySlice1()
	practiceArraySlice2()
}

func practiceArraySlice1() {
	fmt.Println("practiceArraySlice1")
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

func practiceArraySlice2() {
	fmt.Println("practiceArraySlice2")
	s1 := make([]int, 5)
	s2 := s1[1:3]
	s2[0] = 1
	s3 := make([]int, len(s2))
	copy(s3, s2)
	s3[0] = 2
	fmt.Printf("s1: %v, len: %d, cap: %d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))
	fmt.Printf("s3: %v, len: %d, cap: %d\n", s3, len(s3), cap(s3))
	//if s3 is returned the s1's backing array can be garbage collected
}