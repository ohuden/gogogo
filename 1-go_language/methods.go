package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	//fname string
	licenceToKill bool
}

//Methods: Declaration
//method 1
func (ohuden person) speak() {
	fmt.Println("My name is", ohuden.fname)
}

//method 2
func (sa secretAgent) speak() {
	if sa.licenceToKill == true {
		fmt.Println("i have a licence to kill")
	} else {
		fmt.Println("My name is", sa.fname, sa.lname, "and i dont have a licence to kill(((")
	}
}

//interface defines behaviour allows to have polymorphysm
type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		"Oleh",
		"Hudenko",
	}
	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	saySomething(p1)  //polymorphysm
	saySomething(sa1) //polymorphysm
	//p1.speak() // My name is Oleh

	//sa1.speak() //My name is James Bond and i dont have a licence to kill(((
}

//Values and expressions
