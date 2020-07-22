package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

// tow or more consecutive named function parameters share a type
// you can omit the type from all but the last
func add(x, y int) int {
	return x + y
}
// a function can return any number of results
func swap(x, y string) (string, string)  {
	return y, x
}
// go's return values may be named. If so, they are treated as variable defines at the top of the function
// A return statement without arguments return the named return values. This is known as "naked" return
// Naked return statements should be used only in short functions, as they can harm readability in longer ones
func split(sum int) (x, y int)  {
	x = sum * 4 / 9
	y = sum - x
	return
}
// var statement declares a list of variables; as in function argument lists, the type is last
// a var statement can be at package or function level.
var c, python, java bool
// variable initializer
var j, k int = 2, 3

// the example shows variables of several types, and also that variable declarations
// may be "factored" into blocks, as with import statements
var (
	ToBe bool = false
	Hello string = "Hello"
	Inta int = 589
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5+12i)
)

func main()  {
	var i int

	fmt.Println("Hello World")
	fmt.Println("the time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(math.Pi)

	fmt.Println(add(5,6))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))

	fmt.Println(i, j, k, c, python, java)

	// inside a function, the := short assignment statement can be used in place of var declaration with implicit type
	// outside a function, every statement begins with a keyword (var, func and so on) and so the := construct is not available
	var e, f int = 1, 2
	g := 3
	sharp, django, laravel := true, false, "oho!"
	fmt.Println(e, f, g, sharp, django, laravel)
	// With the Go fmt package you can format numbers and strings padded with spaces or zeroes, in different bases
	// fmt.Printf formats and writes to standard output
	// fmt.Sprintf to format a string without printing
	// %T -> type of the value | %t -> boolean as true or false
	// %v -> default format | %+d -> always show sign
	// %s -> quoted string | %q -> quoted string
	// there are several other formatting options ref https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Printf("Type: %T Value: %v\n", Hello, Hello)
	fmt.Printf("Type: %T Value: %v\n", Inta, Inta)

	// if variables are declared without an explicit initial value, their zero value is set
	var ii int
	var ff float64
	var bb bool
	var ss string
	fmt.Printf("%v %v %v %q\n", ii, ff, bb, ss)

	//The expression T(v) converts the value v to the type T.
	//Some numeric conversions:
	iii := 42
	fff := float64(iii)
	uuu := uint(fff)
	fmt.Printf("Type: %T Value: %v\n", iii, iii)
	fmt.Printf("Type: %T Value: %v\n", fff, fff)
	fmt.Printf("Type: %T Value: %v\n", uuu, uuu)

	// When declaring a variable without specifying an explicit type(either by using the := syntax or var = expression syntax),
	// the variable's type is inferred from the value on the right hand side.
	//var i int
	//j := i // Now j is an int

	//Constants
	// Constants are declared like variables, but with the const keyword.
	// Constants can be character, string, boolean, or numeric values.
	// Constants cannot be declared using the := syntax
	const Pi = 3.14
	const World = "World"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
}
