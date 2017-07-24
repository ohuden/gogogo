package main

import "fmt"

//'for' cycle

func evener() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
}

//'Switch' construction
func switcher() {
	for i := 0; i <= 6; i++ {
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		default:
			fmt.Println("Unknown Number")
		}
	}
}

func main() {
	/*
		i := 1
		for i <= 10 {
			fmt.Println(i)          //first version
			i = i + 1
		}
	*/
	/*
		for i := 1; i <= 10; i++ {
			fmt.Println(i)          //looks better right?
		}
	*/
	switcher()
	//evener()
}
