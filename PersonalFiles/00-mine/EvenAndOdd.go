package main

import "fmt"

func main2() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range nums {
		if v%2 == 0 {
			fmt.Println(v, "is odd")
		} else {
			fmt.Println(v, "is even")
		}

	}
}
