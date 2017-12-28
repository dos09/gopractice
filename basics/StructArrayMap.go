package basics

import (
	"fmt"
)

func RunStructArrayMap() {
	fmt.Println(" * RunStructArrayMap")
	practicePointer()
	practiceStruct()
	practiceMaps()
}

func practicePointer() {
	fmt.Println("practicePointer")
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
	fmt.Println("practiceStruct")
	point := Point{1, 2}
	fmt.Println(point.x, point.y)

	ptr := &point
	fmt.Println(ptr.x, ptr.y)
	fmt.Println((*ptr).x, (*ptr).y)

	point2 := Point{x: 4} //y is implicitly 0
	fmt.Println(point2.x, point2.y)
}

func practiceMaps() {
	fmt.Println("practiceMaps")
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

	m := map[string]int{
		"Asen": 30,
		"Ivan": 12, //LAST COMMA IS REQUIRED ...
	}
	fmt.Println(m)
}
