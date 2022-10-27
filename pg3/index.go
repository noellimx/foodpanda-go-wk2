package pg3

import "fmt"

// Factorial. i should be a non-negative integer.
func Factorial() {

	fmt.Printf("[Factorial]\n")
	fmt.Printf("Enter a number\n")

	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Printf("NaN %s\n", err.Error())
		return
	}

	result := factorial(i)
	fmt.Printf(" f(%d) = %d\n", i, result)

}

func factorial(i int) int {
	fmt.Printf("[Factorial]")

	if i == 0 {
		return 1
	}
	if i == 1 {
		return 1
	}

	return i * factorial(i-1)
}
