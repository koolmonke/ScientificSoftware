package main

import "fmt"

func main() {
	n := 5
	m := 9

	array := make([][]int, n)
	for i := range array {
		array[i] = make([]int, m)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			array[i-1][j-1] = i * j
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%2d ", array[i][j])
		}
		fmt.Println()
	}
}
