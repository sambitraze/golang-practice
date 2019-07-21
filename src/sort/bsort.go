package main

import (
	"fmt"
)

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
	var jumbledbox [10]int
	fmt.Println("Enter 10 nos:\n")
	for a := 0; a < 10; a++ {
		fmt.Scanln(&jumbledbox[a])
	}
	bubbleSort(jumbledbox)
}
