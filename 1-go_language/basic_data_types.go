//Basic data types in Go

1. Integers:

	1) signed integer:
		int8, int16, int32, int64

	2) unsigned integers:
		uint8, uint16, uint32, uint64
	3) implementation-specific sizes:
		uint  either 32 or 64 bits
		int   same as above
		uintptr unsigned integer large enough to store uninterpreted bits of pointer value

	byte alias for uint8
	rune alias for int32

2. Floating point:
	float32, float64 // info size and representation

	1) Complex:
		complex64, complex128

3. Strings // array of bytes
	string
4. Booleans
	bool // true false
5. Constants
	const
	
	const Hello = "hello" // can be exported out of package
	const hello = "hello" // can't be exported out of package

6. Logical types
		&& И
		|| ИЛИ
		! НЕ

func main() {
	fmt.Println(true && true) 	// true
	fmt.Println(true && false) 	// false
	fmt.Println(true || true)	// true
	fmt.Println(true || false)	// true
	fmt.Println(!true)			// false
}