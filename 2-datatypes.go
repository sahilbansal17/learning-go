package main

import "fmt"

func main() {
	var x1 int8 = 127  // 2^7 - 1
	var x2 int8 = -128 // -2^7

	fmt.Println(x1, x2)

	var y1 int16 = 32767  // 2^15 - 1
	var y2 int16 = -32768 // -2^15

	var z1 int32 = 2147483647  // 2^31 - 1
	var z2 int32 = -2147483648 // -2^31

	var a1 int64 = 9223372036854775807  // 2^63 - 1
	var a2 int64 = -9223372036854775808 // -2^63

	fmt.Println(y1, y2)
	fmt.Println(z1, z2)
	fmt.Println(a1, a2)

	var b1 uint8 = 255 // 2^8 - 1
	var b2 uint8 = 0   // 0

	var c1 uint16 = 65535 // 2^16 - 1
	var c2 uint16 = 0     // 0

	var d1 uint32 = 4294967295 // 2^32 - 1
	var d2 uint32 = 0          // 0

	var e1 uint64 = 18446744073709551615 // 2^64 - 1
	var e2 uint64 = 0                    // 0

	fmt.Println(b1, b2)
	fmt.Println(c1, c2)
	fmt.Println(d1, d2)
	fmt.Println(e1, e2)

	// rune  	// alias for int32
	// byte 	// alias for int8 => single character

	var x rune // default int32, value 0
	var y byte // default int8, value 0

	fmt.Println(x, y)

	var z bool = true // default false
	fmt.Println(z)

	var a float32 = 3.14 // default float64
	var b float64 = 3.14

	fmt.Println(a, b)

	var c string = "Hello World!" // default ""
	fmt.Println(c)

	c = "Hello World! 2"
	fmt.Println(c)

	const d string = "Hello World! 3"
	fmt.Println(d)

	// d = "Hello World! 4" // error: cannot assign to d
}
