package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr : 0 plus
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128 : really large number

	// Using var

	// *** assignment ***//
	// var name string = "Ivy"    should omit type, it's inferred
	var name = "Ivy"
	// var age int = 25
	var age = 25
	const isCool = true
	// *** END assignment ***//

	// shorthand
	//name := "liang"

	// *** test ***//
	fmt.Println(name, age, isCool)
	// get the type of the value
	fmt.Printf("%T\n", isCool)
	// *** END of test ***//
}
