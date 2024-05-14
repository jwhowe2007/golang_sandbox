package main

import (
	"fmt"
	"math"
	"math/rand"
	"math/cmplx"
)

// Package level variables
var c, python, golang, java, php bool

// Package level variables with initializers
var i, j int = 1, 2

const sqMeridian = 1.41425

// Standard function with a return value
func cube(x float64) float64 {
	return x * x * x
}

// Multiple args of the same type can be combined
func mult(a, b int) int {
	return a * b
}

// Returning the inputs in opposite order
func swap(a, b string) (string, string) {
	return b, a
}

// Short return func
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func types() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y)) + 0.5
	var z uint = uint(f)
	foobar := math.Pi

	fmt.Println(x,y,f,z,foobar)
	fmt.Printf("foobar is of type %T\n", foobar)
}

func needInt(x int) int { return x*10 + 1}
func needFloat(x float64) float64 { return x * 0.1}

func constants() {
	const World = "世界"
	const Flag = true

	const (
		Bignum = 1 << 100
		Smallnum = Bignum >> 99
	)

	fmt.Println("Hello", World, "!")
	fmt.Println("Happy", sqMeridian, "Day!")
	fmt.Println("Go rules?", Flag)
	fmt.Println(needInt(Smallnum))
	fmt.Println(needFloat(Smallnum))
	fmt.Println(needFloat(Bignum))
}

func main() {
	var a float64 = 2
	var b float64 = 2
	var s1 string = "foo"
	var s2 string = "bar"
	swapped1, swapped2 := swap(s1, s2)
	implicit := "foobar"

	var (
		logical 		 bool       = true
		foo				 string     = "foo"
		genericSignedInt int        = -17
		signedByte		 int8  	    = -120
		signedShort		 int16 	    = -32700
		signedLong		 int32 	    = -2000000000
		signedWord		 int64 	    = -920000000000000000
		genericUnsigned	 uint  	    = 25
		unsignedByte	 uint8 	    = 250
		unsignedShort	 uint16     = 65500
		unsignedLong	 uint32     = 4100000000
		unsignedWord	 uint64     = 18000000000000000000
		unsignedPointer	 uintptr    = 32
		byteVar			 byte	    = 1<<8  - 1
		inscription		 rune	    = 0x16CF
		shortFloat		 float32    = math.Pi / 7
		longFloat		 float64    = math.Pi / 256
        complexNum       complex64  = complex(10,20)
		z  	             complex128 = cmplx.Sqrt(complex(1,-1))
	)

	fmt.Println("Your number of the day is:", rand.Intn(100))
	fmt.Println("Pythagorean theorem example (c^2=a^2+b^2):", math.Sqrt(a*a+b*b))

	fmt.Println("pi/4 radians:", math.Pi / 4.0)

	fmt.Println("2 cubed is:", cube(2))
	fmt.Println("3 multiplied by 123 is:", mult(3,123))
	fmt.Println("Swapping 'foo' and 'bar':", swapped1, swapped2)

	fmt.Println("Split number 27:")
	fmt.Println(split(27))

	fmt.Println(c, python, java, php, golang)
	fmt.Println(i, j)

	fmt.Println("Implicit variable:", implicit)

	fmt.Println("Boolean value:", logical)
	fmt.Println("String value:", foo)
	fmt.Println("SignedInt value:", genericSignedInt)
	fmt.Println("Signed byte value:", signedByte)
	fmt.Println("Signed short int value:", signedShort)
	fmt.Println("Signed long int value:", signedLong)
	fmt.Println("Signed word int value:", signedWord)
	fmt.Println("UnsignedInt value:", genericUnsigned)
	fmt.Println("Unsigned byte value:", unsignedByte)
	fmt.Println("Unsigned short value:", unsignedShort)
	fmt.Println("Unsigned long value:", unsignedLong)
	fmt.Println("Unsigned word value:", unsignedWord)
	fmt.Println("Unsigned pointer value:", unsignedPointer)
	fmt.Println("Byte value:", byteVar)
	fmt.Printf("Runic inscription value: %+#U\n", inscription)
	fmt.Println("Short float value:", shortFloat)
	fmt.Println("Long float value:", longFloat)
	fmt.Println("Complex value:", complexNum)
    fmt.Println("Complex z-value:", z)

	types()
	constants()
}