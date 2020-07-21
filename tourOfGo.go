package main

import (
	"fmt"
	"math"
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

func main()  {
	fmt.Println("Hello World")
	fmt.Println("the time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(math.Pi)

	fmt.Println(add(5,6))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))

}
