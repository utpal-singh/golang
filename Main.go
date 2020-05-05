package main

import (
	"fmt"
	"strconv"
)

var (
	actorname string = "Utpal Singh"
	companion string = "Male"
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

	var i int = 42

	fmt.Printf("%v, %T\n", i, i)

	var j string

	j = strconv.Itoa(i)

	fmt.Printf("%v, %T", j, j)
}
