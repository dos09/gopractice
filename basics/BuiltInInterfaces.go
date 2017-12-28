package basics

import(
	"fmt"
)

func RunBuiltInInterfaces() {
	fmt.Println(" * RunBuiltInInterfaces")
	practiceError()
	practiceStringer()
}

func practiceError() {
	fmt.Println("practiceError")
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	//without casting 'e', Sprintf will go into infinite loop
	return fmt.Sprintf("Can't Sqrt negative value %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return 0, nil
}

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v)", p.Name, p.Age)
}

func practiceStringer() {
	fmt.Println("practiceStringer")
	p := Person{"Asen", 30}
	fmt.Println(p)
}
