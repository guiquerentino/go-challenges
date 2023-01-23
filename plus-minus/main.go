package main

import "fmt"

/*

Given an array of integers, calculate the ratios of its elements that are positive, negative, and zero. Print the decimal value of each fraction on a new line with  places after the decimal.

*/

var (
	positive = 0
	negative = 0
	zero     = 0
)

func main() {

	var arr = []int32{-4, 3, -9, 0, 4, 1}

	for _, j := range arr {

		if j < 0 {
			negative++
		}

		if j > 0 {
			positive++
		}

		if j == 0 {
			zero++
		}

	}

	fmt.Println(float64(positive) / float64(len(arr)))
	fmt.Println(float64(negative) / float64(len(arr)))
	fmt.Println(float64(zero) / float64(len(arr)))

}
