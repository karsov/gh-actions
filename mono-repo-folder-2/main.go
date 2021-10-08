package main

import "fmt"

func main() {
	fmt.Println("The sum is", Sum("1", "1"))
}

// Sum concatenates two strings
func Sum(a, b string) string {
	return a + b
}
