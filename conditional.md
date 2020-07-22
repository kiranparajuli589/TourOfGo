## Go Conditional
### If
Go's if statements are like its for loops; the expression need not be surrounded by parentheses `()` but the braces `{}` are required

```go
package main

import (
"fmt"
 "math"
)

func sqrt(x float64) string {
    if x < 0 {
        return sqrt (-x) + "i"
    }
    return fmt.Sprint(math.Sqrt(x))
}

func main() {
    fmt.Println(sqrt(2), sqrt(-4))
}
```

#### If with a short statement
Like for, the if statement can start with a short statement to execute before the condition.
Variables declared by the statement are only in scope until the end of the if.
```go
package main
import (
    "fmt"
    "math"
)
func pow(x, n, lim float64) float64  {
    if v:=math.Pow(x, n); v<lim {
        return v
    }
    return lim
}
func main()  {
    fmt.Println(
        pow(3,2,10),
        pow(3,3,10),
    )
}
```
#### If and else
Variables declared inside an if short statement are also available inside any of the else blocks.

(Both calls to pow return their results before the call to fmt.Println in main begins.)
```go
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
```
### Switch statements
A switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression.

Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.
```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
```

#### Switch evaluation order
Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
For example, 
```go
switch i {
    case 0:
    case f():
}
```
does not call `f` if `i == 0`
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
```
#### Switch with no condition
Switch without a condition is the same as `switch true`
This construct can be a clean way to write long if-then-else chains.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
```

### Defer
A defer statement defers the execution of a function until the surrounding function returns.
The deferred call's arguments are evaluated immediately, nut the function call is not executed until the surrounding function returns.
```go
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
```
#### Stacking defers
Derferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
[ReadMore Here](https://blog.golang.org/defer-panic-and-recover)
```go
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
```


