package main

//горутины — это функция, которая может работать параллельно с другими функциями.

import (
	"fmt"
	"math/rand"
	"time"
)

/*
func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}
*/
func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250)) //добавим задержку
		time.Sleep(time.Millisecond * amt)
	}
}

/*
func main() { //горутина
	go f(0) //горутина №2
	var input string
	fmt.Scanln(&input) //без него программа завершится еще перед тем, как ей удастся вывести числа.
}
*/
/*
func main() {
	for i := 0; i < 10; i++ { // сделаем много горутин
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}
*/
//Каналы - обеспечивают возможность общения нескольких горутин друг с другом, чтобы синхронизировать их выполнение

func pinger(c chan string) { //Тип канала представлен ключевым словом chan
	for i := 0; ; i++ {
		//Оператор <- используется для отправки и получения сообщений по каналу
		c <- "ping" //Конструкция c <- "ping" означает отправку "ping
	}
}
func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}
func printer(c chan string) {
	for {
		msg := <-c //Конструкция означает получение "ping" и сохранение в переменную msg
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

/*
func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
*/
//Направление каналов
/*
func pinger(c chan<- string)  // канал c будет только отправлять сообщение
func printer(c <-chan string) // канал будет только принимать сообщение
*/

//Селект - оператор select который работает как switch, но для каналов

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}

/*
Буферизированный канал
При инициализации канала можно использовать второй параметр:

c := make(chan int, 1)
и мы получим буферизированный канал с ёмкостью 1.
Обычно каналы работают синхронно - каждая из сторон ждёт, когда другая сможет получить или
передать сообщение. Но буферизованный канал работает асинхронно — получение или отправка
сообщения не заставляют стороны останавливаться. Но канал теряет пропускную способность,
когда он занят, в данном случае, если мы отправим в канал 1 сообщение, то мы не сможем отправить
туда ещё одно до тех пор, пока первое не будет получено.
*/
