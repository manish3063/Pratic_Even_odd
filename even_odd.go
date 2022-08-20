package main

import "fmt"

func main() {
	final := Evenodd(1)
	fmt.Println(final)

}
func Evenodd(num int) bool {
	result := false
	if num%2 == 0 {
		fmt.Println(num, "number is even")
		result = true
		return result

	} else {
		fmt.Println(num, "number is odd")
		result = false
		return result

	}
	return result
}
