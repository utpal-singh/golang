package main

import "fmt"

var (
	actorname string = "NSobbs knkasw"
	companion string = "kacbbk"
)

var (
	age  int     = 21
	year float64 = 3.75
)

var (
	I int = 55
)

var j int = 32

func main() {
	//i := 42.

	//var j float32 = 42

	// var i int
	// i = 32.0

	//j := 23

	//var theHTTPRequest string = "https://google.com"

	//var i int = 42

	//fmt.Printf("%v, %T\n", i, i)

	//var j string

	//j = strconv.Itoa(i)

	//var n bool = true
	//var k bool = false

	//a := 1 == 1
	//b := 2 == 3

	//var f bool

	//var a1 int8 = 100
	//fmt.Printf("%v, %T\n", a1, a1)

	//var a2 uint16 = 20
	//fmt.Printf("%v, %T\n", a2, a2)

	//var a3 float32 = 5.0

	//x1 := 2
	//x2 := 3

	//fmt.Printf("%v\n", a3+a1)
	//fmt.Println(x1 - x2)
	//fmt.Println(x1 % x2)
	//fmt.Println(x1 / x2)

	//fmt.Printf("%v, %T\n", n, n)
	//fmt.Printf("%v, %T\n", k, k)
	//fmt.Printf("%v, %T\n", a, a)
	//fmt.Printf("%v, %T\n", b, b)
	//fmt.Printf("%v, %T\n", f, f)

	//Bitwise Operators
	//a := 10 //0010
	//b := 3  //1011

	//fmt.Println(a & b)  //0010 and
	//fmt.Println(a | b)  //1011 or
	//fmt.Println(a ^ b)  //1001 Any of the bitset but not both
	//fmt.Println(a &^ b) //0100 None of the bitset

	//BitShifting

	//a := 8              // 2^3 1000
	//fmt.Println(a << 3) // 2^3*2^3 = 2^6 1000000
	//fmt.Println(a >> 3) // 2^3/2^3 = 2^0 1

	//b := 10             // 2^3 + 2^1 1010
	//fmt.Println(b << 3) // 2^3*2^3 = 2^6 1000000
	//fmt.Println(b >> 3) // 2^3/2^3 = 2^0 1

	//Floating Point

	//n := 3.14
	//p := 13.6e72
	//q := 2.1E14

	//var r float32
	//r = 2e32

	//fmt.Printf("%v, %T\n", n, n)
	//fmt.Printf("%v, %T\n", p, p)
	//fmt.Printf("%v, %T\n", q, q)
	//fmt.Printf("%v, %T\n", r, r)

	//Complex Primitive type

	//var n complex64 = 1 + 2i

	//fmt.Printf("%v, %T\n", n, n)
	//fmt.Printf("%v, %T\n", real(n), real(n))
	//fmt.Printf("%v, %T\n", imag(n), imag(n))

	// var s complex128 = 1 + 2i

	// fmt.Printf("%v, %T\n", s, s)
	// fmt.Printf("%v, %T\n", real(s), real(s))
	// fmt.Printf("%v, %T\n", imag(s), imag(s))

	//Arrays

	// s := "This is an array"

	// fmt.Printf("%v, %v, %T\n", string(s[8]), s[8], s[8])

	//Constants

	//const myConst float64 = math.Sin(1.57)
	// fmt.Printf("%v, %T\n", math.Sin(1.57), math.Sin(1.57))
	//Won't work, Constants can only be made from Primitive dAtatypes

	const a int = 23
	const b string = "Hi, this is Go"
	const c bool = true
	const d float64 = 23.43

	// const d float64 = 2324.23

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)

	// const a = 15
	// var b int16 = 40

	// fmt.Printf("%v, %T\n", a+b.a+b)
	// //Works in previopus Go, not in current bersions

}
