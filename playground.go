package main
import (
	"fmt"
	"math"
)
type MyFloat float64

func (f MyFloat) Abs() int64 {
	if f < 0 {
		return int64(-f)
	}
	return int64(f)
}
func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f)
	fmt.Println(f.Abs())
}
