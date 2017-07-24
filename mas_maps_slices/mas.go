package main

import (
	"fmt"
)

//array

var x [5]int //definition

//slice
var y []float64 //definition

//map
var m map[string]int

func main() {

	//array
	x[4] = 100
	fmt.Println(x) //[0 0 0 0 100]

	//slice
	z := make([]float64, 5, 10)      //declaration of slice (10 - len of array), (5 - len of slice)
	fmt.Println(z)                   //[1 2 3 4 5]
	arr := [5]float64{1, 2, 3, 4, 5} //declaration of array
	s := arr[0:5]                    //declaration of slice [low: high]
	fmt.Println(s)                   //[1 2 3 4 5]

	//slice functions
	//append
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5, 6, 7) //appends the first element with all next
	fmt.Println(slice1, slice2)          //[1 2 3] [1 2 3 4 5 6 7]
	//copy
	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)        // sclice3 content will be copired to slice4(currently 2 elements will be copied as there is no more space in slice4)
	fmt.Println(slice3, slice4) // [1 2 3] [1 2]

	//map
	m := make(map[string]int)
	m["key"] = 10
	fmt.Println(m["key"]) //10

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Li"])

	delete(elements, "H") //deletes form elements map "H" key pair
	fmt.Println(elements)

}
