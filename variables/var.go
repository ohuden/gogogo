package main

import "fmt"

//Variables
const z string = "U" //область видимости только внутри пакета
var (
	a = 5
	b = 10
	c = 15
)

func main() {
	/*
		var x string = "Hello World"
		fmt.Println(x)
	*/
	var x string //область видимости только внутри функции
	x = "Hello World"
	fmt.Println(x)

	const y string = "Hello 2"
	fmt.Println(y, z)

	// square programm
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * input

	fmt.Println(output)
}
