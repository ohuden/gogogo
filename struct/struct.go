package main

//Структура — это тип, содержащий именованные поля
//type вводит новый тип
//За ним следует имя нового типа Circle и ключевое слово struct
//struct говорит, что мы определяем структуру и список полей внутри фигурных скобок
//Каждое поле имеет имя и тип
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
	fmt.Println(c.x, c.y, c.r) //Получить доступ к полям можно с помощью оператора .
	c.x = 10
	c.y = 5
}

//Методы
//метод — функцию особого типа
func (c *Circle) area() float64 {//Между ключевым словом func и именем функции мы добавили «получателя»
    return math.Pi * c.r*c.r
}
//Получатель похож на параметр — у него есть имя и тип, но объявление функции таким способом позволяет нам вызывать функцию с помощью оператора 
fmt.Println(c.area())

//Встраиваемые типы
type Person struct {
    Name string
}
func (p *Person) Talk() {
    fmt.Println("Hi, my name is", p.Name)
}
//если мы хотим создать новую структуру Android, то можем сделать так:
type Android struct {
    Person Person
    Model string
}
a := new(Android)
a.Person.Talk()
//Это отношение работает достаточно интуитивно: личности могут говорить, андроид это личность, значит андроид может говорить

//Интерфейсы

type Shape interface {//вместо того, чтобы определять поля(как в структурах), мы определяем «множество методов»
    area() float64
}
/*
В нашем случае у Rectangle и Circle есть метод area, который возвращает float64,
получается они оба реализуют интерфейс Shape. Само по себе это не очень полезно,
но мы можем использовать интерфейсы как аргументы в функциях:
*/
func totalArea(shapes ...Shape) float64 {
    var area float64
    for _, s := range shapes {
        area += s.area()
    }
    return area
}

//в мейн функции вызывать будем так
fmt.Println(totalArea(&c, &r))

//Интерфейсы также могут быть использованы в качестве полей:
type MultiShape struct {
    shapes []Shape
}
//Мы можем даже хранить в MultiShape данные Shape, определив в ней метод area:
func (m *MultiShape) area() float64 {
    var area float64
    for _, s := range m.shapes {
        area += s.area()
    }
    return area
}
//Теперь MultiShape может содержать Circle, Rectangle и даже другие MultiShape.