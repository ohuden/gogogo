package main

import "fmt"

/*
func zero(x int) {
	x = 0
}
func main() {
	x := 5
	zero(x)
	fmt.Println(x) // x всё еще равен 5
}
*/
// * Указатели указывают (прошу прощения за тавтологию) на участок в памяти, где хранится значение.
func zero(xPtr *int) { //Используя указатель (*int) в функции zero, мы можем изменить значение оригинальной переменной.
	*xPtr = 0
}
func main() {
	x := 5
	//& используется для получения адреса переменной
	zero(&x)       // &x вернет *int (указатель на int) потому что x имеет тип int.
	fmt.Println(x) // x is 0
}

// new принимает аргументом тип, выделяет для него память и возвращает указатель на эту память
func one(xPtr *int) {
	*xPtr = 1
}
func main2() {
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // x is 1
}

//Go — язык с автоматической сборкой мусора. Это означает, что область памяти очищается автоматически, когда на неё не остаётся ссылок.