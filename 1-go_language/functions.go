//Declaration

func имяФункции ([список параметров]) [возвращаемые значения] {
	тело функции
}

//несколько значений одного типа
func sumLight(i, j int) int {
	return i + j
}

//для получения произвольного списка однотипный значений
func sumMore(stuff ...int) (res int)  {
	for i := range stuff {
		res += stuff[i]
	}
	return
}
//Будет выполнена во время начала исполнения программы
var stuff = "not ready" 

func init() {
	stuff = "ready" //подготовка переменных 
}

func main() {
	fmt.Println(stuff) //ready
}

//Recursion
//Variadic functions
//Error handling
//Panic and recover
//Named result parameters


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