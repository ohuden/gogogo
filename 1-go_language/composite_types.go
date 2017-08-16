//Arrays (len are part of its type) low level element
//====================================================
var a1 [3]int //[0,0,0]
a1[1] := 1 //[0,1,0]
var aa[3][3]int//array of arrays

//Slices
//====================================================
var sl []int //slice []
sl = append(sl, 100) //adding element to slice

/*

struct
	len 1					len 3
	cap 2			->		cap 4
	array [0,0]				array [0,0,0,0]

*/

len(sl) // 1
cap(sl) // 1

sl2 := []int{10, 20, 30}
cap(sl2) // 3
len(sl2) // 2

sl = append(sl ,sl2...)	//adding slice to slice types mistmatch

var slm [][]int
slm = append(slm, sl2) 	//types match

//creating slice with required length
slice3:=make([]int, 10)
slice3 = append(slice3, 1)

//creating slice with length and capacity
slice4:= make([], 10, 15)
len(slice4) // 10
cap(slice4) // 15

//copying slices correctrly
slice5 : =make([]int, len(slice4), len(slice4))
copy(sslice5, slice4)

slice5[1:4] // icluding
slice5[:4]
slice5[1:]

//creating slice from array

a := [...]int{5,6,7}
sl8 := a[:]
a[1] = 8
fmt.Println(sl8) // [5, 8, 7]

//Maps/hash/accociative array
//====================================================

var m map[string]string
var m2 map[string] map[string]string //hash 2 lvl or map of maps

//full initializations

var m map[string]string = map[string]string{}
m2 := map[string]string{}
m2["test"] = "ok" // map[test:ok]

var m3 = make(map[string]string)
m3["FirstName"] = "Vasia"

//if we call for unexisting key it'll give the default value
lastNamr := m3["lastName"] // empty string

//checking if theres a value
lastName, ok := m3["lastName"]
fmt.Println("lastName is", lastName, "exist:", ok) // lastName is  exist; false

//checking presence
_, exist := m3["firstName"]
fmt.Println("firstName exist:", exist)	// firstName exist: true

//deleting value
delete(m3, "firstName")



//Chan
//====================================================

//Каналы - обеспечивают возможность общения нескольких горутин друг 
//с другом, чтобы синхронизировать их выполнение

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


func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}

//Направление каналов

func pinger(c chan<- string)  // канал c будет только отправлять сообщение
func printer(c <-chan string) // канал будет только принимать сообщение


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

//Structs https://www.youtube.com/watch?v=8Rg4bOn9Qds
//====================================================

type Circle struct {
	x, y, r float64
}

//Inicialization
var c Circle
//or
c := new(Circle)

//adding some values to fields
c := Circle{x: 0, y: 0, r: 5}
//or
c := Circle{0, 0, 5}

func main()  {
	fmt.Println(c.x, c.y, c.r) //to get acces to value use .
	c.x = 10
	c.y = 5
}


//iota usage
//====================================================

const (
	_       = iota            	//iota = 0
	a uint64 = 1 << (10*iota) 	//iota = 1
	b							//iota = 2
	_							//iota = 3
	d 							//iota = 4
 )