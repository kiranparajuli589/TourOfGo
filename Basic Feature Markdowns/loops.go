package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("The sum of digits upto 10 is", sum)
	fmt.Println("+ --------------------------------------------------------- +")
	fmt.Println("| Multiplication Table                                      |")
	fmt.Println("+ --------------------------------------------------------- +")
	for i := 1; i <= 10; i++ {
		fmt.Println("+ --------------------------------------------------------- +")
		fmt.Println("| Table of", i, "                                              |")
		fmt.Println("+ --------------------------------------------------------- +")
		for j:=1; j<=10; j++ {
			fmt.Println("|", i, "*", j, "=", i*j, "                                              |")
		}
	}
	fmt.Println("+ --------------------------------------------------------- +")

	// For is Go's "while"
	sumation := 1
	for sumation < 1000 {
		sumation +=sumation
	}
	fmt.Println(sumation)
	// if you omit the loop condition it loops forever, so an infinite loop is compactly expressed
	//for {
	//	fmt.Println("Hello Forever")
	//}

	// Range
	// The range form of the for loop iterates over a slice or map.
	// When ranging over a slice, two values are returned for each iteration. The first is the index,
	// and the second is a copy of the element at that index
	pow := []int{1,2,4,8,16,32,64,128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	// you can skip the index or value by assigning to _
	power := make([]int, 10)
	for i := range power {
		power[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range power {
		fmt.Printf("%d\n", value)
	}
}
