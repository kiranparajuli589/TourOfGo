### Arrays
The type `[n]T` is an array of `n` values of type `T`.
The expression:
```go
package main
var a [10] int
```
declares a variable `a` as an array of ten integers.
An array's length is a part of its type, so array cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays.
```go
package main
import "fmt"
func main() {
    var a [2]string
    a[0] = "Hello"
    a[1] = "World"
    fmt.Println(a[0], a[1])
    fmt.Println(a)
    
    primes := [6]int{2,3,5,7,11,13}
    fmt.Println(primes)
}
```

#### Slices
An array has a fixes size. A slice, on the other hand, is a dynamically-sized,, flexible wiew into the elements of an array, In practice, slices are much more common than arrays.
The type `[]T` is a slice with elements of type `T`.
A slice is formed by specifying two indices, a low and high bound, separated by a colon:
```go
a[low: high]
```
This selects a half-open range which includes the first element, but excludes the last one.
The following expression creates a slice which includes elements 1 through of `a`:
```go
a[1:4]
```

```go
package main
import "fmt"
func main()  {
    primes := [6]int{2,3,5,7,11,12}
    var s []int = primes[1:4]
    fmt.Println(s)
}
```

#### Slices are like references to arrays
A slice does not store any data, it just describes a section of an underlying array. 
Changing the elements of a slice modifies the corresponding elements of its underlying array.
Other slices that share the same underlying array will see those changes.
```go
package main
import "fmt"
func main()  {
    names := [4]string{
        "John",
        "Paul",
        "George",
        "Kiran",
    }
    fmt.Println(names)
    
    a := names[0:2]
    b := names[1:3]
    fmt.Println(a, b)
    
    b[0] = "XXX"
    fmt.Println(a, b)
    fmt.Println(names)
}
```
#### Creating slice with make
Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
Then make function allocates a zeroed array and returns a slice that refers to that array:
```go
a := make([]int, 5)    // len(a) = 5
```
To specify a capacity, pass third argument to make:
```go
b := make([]int, 0, 5)    // len(b) = 0, cap(b) = 5

b = b[:cap(b)]   // len(b) = 5, cap(b) = 5
b = b[1:]   // len(b) = 4, cap(b) = 4
```

```go
package main
import "fmt"
func printSlice(s string, x []int) {
    fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
func main()  {
    a := make([]int, 5)
    printSlice("a", a) // len 5 cap 5 [0 0 0 0 0]
    b := make([]int, 0, 5)
    printSlice("b", b) // len 0 cap 5 []
    c := b[:2]
    printSlice("c", c) // len 2, cap 5 [0 0]
    d := c[2:5]
    printSlice("d", d) // len 3, cap 3 [0 0 0]
}
```

#### Slices of slices
Slices can contain any type, including other slices

```go
package main
import (
    "fmt"
    "strings"
)
func main()  {
    // Create a tic-tac-toe board
    board := [][]string{
        []string{"_","_","_"},
        []string{"_","_","_"},
        []string{"_","_","_"},
    }
    // The players take turns
    board[0][0] = "X"
    board[2][2] = "0"
    board[1][2] = "X"
    board[1][0] = "0"
    board[0][2] = "X"
    
    for i := 0; i < len(board); i++ {
        fmt.Printf("%s\n", strings.Join(board[i], " "))
    }
}
```

#### Appending to a slice
It is common to append new elements to a slice, and so Go provides a built-in `append` function. The [documentation](https://golang.org/pkg/builtin/#append) of the built-in package describes append.

```go
func append(s []T, vs ...T) []T
```
The first parameter `s` of `append` is a slice of type `T`, and the rest are `T` values to append to the slice.

The resulting value of `append` is a slice containing all the elements of the original slice plus the provided values.

If the backing array of `s` is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.

([To learn more about slices, read the Slices: usage and internals article.](https://blog.golang.org/slices-intro))
```go
package main
import "fmt"
func main()  {
    var s []int
    printSlice(s)
    
    //append works on nil slices.
    s = append(s, 0)
    printSlice(s)

    //The slice grows as needed
    s = append(s, 1)
    printSlice(s)
    
    // we can add more than one element at a time
    s = append(s, 2,3,4)
    printSlice(s)
}
func printSlice(s []int)  {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```
