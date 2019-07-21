package main

import "fmt"

func check(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			println("number")
			return false
		}
	}
	println("alphabet")
	return true
}
func main() {
	fmt.Println("enter a string to check weeather its a word or number (true for alphabet false for number) ")
	var input string
	fmt.Scanln(&input)
	fmt.Println(check(input))

}
