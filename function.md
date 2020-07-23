## Function values
Functions are values too. they can be passed around just like other values.
Function values may be used as function arguments and return values.
```go
package main
import (
    "fmt"
    "math"
)
func compute(fn func(float64, float64) float64) float64 {
    return fn(3, 4)
}
func main()  {
    hypot := func(x, y float64) float64 {
        return math.Sqrt(x*x + y*y)
    }
    fmt.Println(hypot(2, 3))
    fmt.Println(compute(hypot))
    fmt.Println(compute(math.Pow))
}
```
#### Function closures
Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variable
For example, the `adder` function returns a closure. Each closure is bound to its own `sum` variable.
```go
package main
import "fmt"
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}
func main() {
    pos, neg := adder(), adder()
    for i := 0; i< 10; i++ {
        fmt.Println(pos(i), neg(-2*i))
    }
}
```
#### Fibonacci closure
Let's have some fun
```go
package main
import "fmt"
func fibonacci() func () int {
    first, second := 0, 1
    return func() int {
        ret := first
        first, second = second, first + second
        return ret
    }
}
func main()  {
    f := fibonacci()
    for i:=0; i<10; i++ {
        fmt.Println(f())
    }
}
```
