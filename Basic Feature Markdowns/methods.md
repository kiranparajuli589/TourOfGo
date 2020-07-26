## Methods
Go does not have `classes`. However, you can define methods on `types`.

A `method` is a function with a special `receiver` argument.

The receiver appears in its own argument list between the `func` keyword and the method name.

In this example, the `Abs` method has a receiver of type `Vertex` and named v

```go
package main

import (
    "fmt"
    "math"
)
type Vertex struct {
    X, Y float64
}
func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
func main()  {
	v := Vertex{3,4}
	fmt.Println(v.Abs())
}
```
### Methods are functions
Remember: a method is just a function with a `receiver` argument.

Here's `Abs` written as a regular function with no change in functionality.
```go
package main
import (
    "fmt"
    "math"
)
type Vertex struct{
	X, Y float64
}
func Abs(v Vertex) float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func main()  {
    v := Vertex{3,4}
    fmt.Println(Abs(v))
}
```

You can declare a method on non-struct types too.

In this example we see a numeric type `MyFloat` with an `Abs` method.

You can only declare a method with a receiver whose type is defined in the same package as the method. Your cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).

```go
package main
import (
    "fmt"
    "math"
)
type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}
func main() {
    f := MyFloat(-math.Sqrt2)
    fmt.Println(f.Abs())
}
```

### Pointer receivers
You can declare methods with pointer recievers.

This means the receiver type has the literal syntax `*T` for some type `T`. (Also, `T` cannot be a pointer such as `*int`.)

For example, the `Scale` method here is defined on `*Vertex`.

Methods with pointer receivers can modify the value to which the receiver points (as `Scale` does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

```go
package main

import (
    "fmt"
    "math"
)
type Vertex struct{
	X, Y float64
}
func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// f is a scale factor
func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func main()  {
    v:= Vertex{3,4}
    v.Scale(10)
    fmt.Println(v.Abs())
}
```

In the above example, try removing the `8` from the declaration of the `Scale` function on line 16 and observe how the program's behavior changes.

With a value receiver, the `Scale` method must have a pointer receiver to change the `Vertex` value declared in the `main` function.

### Pointers and Functions
Here we see the `Abs` and `Sclae` methods written as functions.

```go
package main
import (
    "fmt"
    "math"
)
type Vertex struct{
    X, Y float64
}
func Abs(v Vertex) float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func Scale(v *Vertex, f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}
func main() {
    v := Vertex{3,4}
    Scale(&v, 10)
    fmt.Println(Abs(v))
}
```
Again, try removing the `*` from line 16. Same behavior again!

### Methods and pointer indirection
Comparing the previous two programs, you might notice that functions with a ponter argument must take a pointer:

```go
var v Vertex
ScaleFunc(v, 5) // Compile error!
ScaleFunc(&v, 5) // OK
```
While methods with pointer receivers take either a value or a pointer as the receiver when they are called:
```go
var v Vertex
v.Scale(5) // OK
p := &v
p.Scale(10) // OK
```
For the statement `v.Scale(5)`, even though `v` is a value and not a pointer, the method with the pointer receiver is called automatically. That is, as a convenience, Go interprets the statement `v.Scale(5)` as (&v).Scale(5) since the `Scale` method has a pointer receiver.

```go
package main
import (
    "fmt"
    "math"
)
type Vertex struct{
	x, y float64
}
func (v *Vertex) Scale(f float64) {
    v.x = v.x * f
    v.y = v.y *f
}
func ScaleFunc(v *Vertex, f float64) {
    v.x = v.x * f
    v.y = v.y * f
}
func main() {
    v := Vertex{3,4}
    v.Scale(2)
    ScaleFunc(&v, 10)
    
    p := &Vertex{4,3}
    p.Scale(3)
    ScaleFunc(p, 8)
    
    fmt.Println(p, v)
}
```

The equivalent thing happens in the reverse direction.

Functions thath take a value argument must take a value of that specific type:

```go
type Vertex struct {}
var v Vertex
fmt.Println(AbsFunc(v)) // OK
fmt.Println(AbsFunc(&v)) // Compile Error!
```
```go
type Vertex struct {}
var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
```
In this case, the method call `p.Abs()` is interpreted as `(*p).Abs()`

### Choosing a value or pointer receiver
There are two reasons to use a pointer receiver.

The first is so that the method can modify the value that its receiver points to.

The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, For example:

In this example, both `Scale` and `Abs` are with receiver type `*Vertex`, even though the `Abs` method needn't modify the receiver.

In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both. (We'll see why over the next few pages.)

```go
package main
import (
    "fmt"
    "math"
)
type Vertex struct{
	x, y float64
}
func (v *Vertex) Scale(f float64) {
    v.x = v.x * f
    v.y = v.y * f
}
func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.x*v.x + v.y*v.y)
}
func main()  {
    v := &Vertex{3,4}
    fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
    v.Scale(5)
    fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
```
