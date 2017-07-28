package main

import "fmt"

//Functions
//
//func main( /*some input parameters(имя тип, имя тип)*/ ) /*type of the returned value "float64"*/ { //arguments+ returned value are signature of function
//some body
//}

func f() {
	fmt.Println(x) //undefined
}
func main() {
	x := 5
	f()
}

func f(x int) {
	fmt.Println(x)
}
func main() {
	x := 5
	f(x)
}

//Стек вызовов
//Каждая вызываемая функция помещается в стек вызовов, каждый возврат из функции возвращает нас к предыдущей приостановленной подпрограмме;

func main() {
	fmt.Println(f1())
}
func f1() int {
	return f2()
}
func f2() (r int) {
	r = 1
	return
}

//retun multiple values

func f() (int, int) {
	return 5, 6
}

func main() {
	x, y := f()
}

func add(args ...int) int { //Использование ... перед типом последнего аргумента означает, что функция может содержать ноль и более таких параметров. ей можно передать любое количество аргументов типа int.
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
func main() {
	fmt.Println(add(1, 2, 3))
}

//Closures

func main() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 1))
}

//Recursion

func factorial(x uint) uint { //factorial вызывает саму себя, что делает эту функцию рекурсивной.
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}
