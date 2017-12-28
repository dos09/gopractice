package basics

import (
	"fmt"
	"math/cmplx"
)

const Pi = 3.14

func RunDataTypesConst() { //with capital letter to be exported
	fmt.Println(" * Run_data_types_const")
	var v_bool bool = true
	var v_string string = "ASd"
	var v_int int = 1
	var v_int8 int8 = 2
	var v_int16 int16 = 3
	var v_int32 int32 = 4
	var v_int64 int64 = 5
	var v_uint uint = 6
	var f_float32 float32 = 1.1
	var f_float64 float64 = 2.2
	var c_complex128 complex128 = cmplx.Sqrt(-5 + 12i)
	fmt.Printf("Type: %T, Value: %v\n", v_bool, v_bool)
	fmt.Printf("Type: %T, Value: %v\n", v_string, v_string)
	fmt.Printf("Type: %T, Value: %v\n", v_int, v_int)
	fmt.Printf("Type: %T, Value: %v\n", v_int8, v_int8)
	fmt.Printf("Type: %T, Value: %v\n", v_int16, v_int16)
	fmt.Printf("Type: %T, Value: %v\n", v_int32, v_int32)
	fmt.Printf("Type: %T, Value: %v\n", v_int64, v_int64)
	fmt.Printf("Type: %T, Value: %v\n", v_uint, v_uint)
	fmt.Printf("Type: %T, Value: %v\n", f_float32, f_float32)
	fmt.Printf("Type: %T, Value: %v\n", f_float64, f_float64)
	fmt.Printf("Type: %T, Value: %v\n", c_complex128, c_complex128)
	fmt.Println("Pi:", Pi)
}
