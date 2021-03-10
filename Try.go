package main

import "fmt"

func main() {
	var i [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var sum int = 0
	for _, j := range i {
		sum = sum + j
	}
	fmt.Println(sum)
}
