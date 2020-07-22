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
}
