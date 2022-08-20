package main

import "fmt"

func main() {
	num := 25

	var count, i int

	for i = 1; i <= num; i++ {

		if num%i == 0 {
			count++

		}

	}

	if count == 2 {
		fmt.Println(num, "prime number")
	} else {
		fmt.Println(num, "not a prime number")
	}
}
