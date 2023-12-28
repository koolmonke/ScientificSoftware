package main

import "fmt"

func permutations(n, k int) int {
	if k == 0 {
		return 1
	}

	return n * permutations(n-1, k-1)
}

func main() {
	n := 5
	r := 3
	result := permutations(n, r)
	fmt.Printf("Число размещений %d из %d элементов: %d\n", n, r, result)
}
