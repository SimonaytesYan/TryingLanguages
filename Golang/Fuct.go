//==============================================================================
// 									FACTORIAL
//==============================================================================

package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}

	return fact(n - 1) * n
}

func main() {
	var n int
	fmt.Scanln(&n)
	fmt.Println("Factorial of", n, "=", fact(n))	
}