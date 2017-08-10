//Declaration

func имяФункции ([список параметров]) [возвращаемые значения] {
	тело функции
}

//несколько значений одного типа
func sumLight(i, j int) int {
	return i + j
}

//Будет выполнена во время начала исполнения программы
var stuff = "not ready" 

func init() {
	stuff = "ready" //подготовка переменных 
}

func main() {
	fmt.Println(stuff) //ready
}

//Recursion - expensive in terms of resources

func factorial(x uint) uint{
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	fmt.Println(factorial(4))
}

//Variadic functions

func main() {
	sumMore(4, 8, 7, 4, 6, 7) //we can pass as many arguments as we want
}

func sumMore(stuff ...int) (res int)  {
	for i := range stuff {
		res += stuff[i]
	}
	return
}

//Error handling

type error interface {
	Error() string
}

type MyError struct {
	Msg string
	File string
	Line int
}

switch err:=err.(type) {
	case nil:
		//call succeeded, nothing to do
	case *MyError:
		 fmt.Println("error occured on line:", err.Line)
	default:
		//unknown error
}

if err != nil {
	fmt.Println("handles an error here")
}

//Panic and recover

//go doesnt have exeptions

panic("Panic") //current function stops

//check if panic occured
if r := recover(); r != nil {
	if err, ok := r.(error); ok {
		fmt.Println()
	}
	fmt.Println("Panic deferred")
}

r := recover()

defer //can be used to handle errors when panic is used

//Named result parameters

//возврат нескольких именованых значений
func sumOnlyNatural(stuff ...int) (res int, err error)  {
	for i := range stuff {
		if stuff[i] < 0 {
			err = fmt.Errorf("Only natural numbers expected - given %d", stuff[i])
			return
		}
		res += stuff[i]
	}
	return res, err
}

//Multiple retrun values

//возврат нескольких значений
func sumOnlyNatural(stuff ...int) (int, error)  {
	res := 0
	for i := range stuff {
		if stuff[i] < 0 {
			return 0, fmt.Errorf("Only natural numbers expected - given %d", stuff[i])
		}
		res += stuff[i]
	}
	return res, nil
}
