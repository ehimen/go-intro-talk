package main

import "fmt"

var n = 1000

func main() {
	products := []int{3, 5}

	fmt.Println(sumOfProducts(products))
}

func sumOfProducts(numbers []int) int {
	total := 0

	for i := 1; i < n; i++ {
		for _, product := range numbers {
			if i%product == 0 {
				total += i
				break
			}
		}
	}

	return total
}
