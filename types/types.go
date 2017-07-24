package main

import (
	"fmt"
)

//Types in Go
/*
1. Integers:
	signed integer:
		int8, int16, int32, int64
	unsigned integers:
		uint8, uint16, uint32, uint64
2. Float:
	float32, float64
	Complex:
		complex64, complex128

3. String
4. Logical types

	&& И
	|| ИЛИ
	! НЕ


*/
func main() {
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
