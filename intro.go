package main

import "fmt"

var n = 1000

func main() {
	multiples := []int{3, 5}

	fmt.Println(sumOfMultiplesOf(multiples))
}

func sumOfMultiplesOf(numbers []int) int {
	//var total int = 0
	total := 0

	for potentialProduct := 1; potentialProduct < n; potentialProduct++ {
		for _, multiple := range numbers {
			if potentialProduct%multiple == 0 {
				total += potentialProduct
				break
			}
		}
	}

	return total
}
