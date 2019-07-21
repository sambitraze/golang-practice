package main

import (
	"fmt"
)

var jumbledbox [10]int = [10]int{9, 4, 3, 9, 3, 2, 1, 9, 6, 1}

func bubbleSort(forsort [10]int) {
	n := 10
	ans := true
	for ans {
		// set ans to false
		ans = false
		for i := 1; i < n; i++ {
			if forsort[i-1] > forsort[i] {
				fmt.Println("Swapping")
				forsort[i], forsort[i-1] = forsort[i-1], forsort[i]
				ans = true
			}
		}
	}
	fmt.Println(forsort)
}

func main() {
	fmt.Println("lets sort a phone number")
	bubbleSort(jumbledbox)
}
