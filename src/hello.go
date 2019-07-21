package main

import "fmt"

var (
	p  int    = 4
	uu string = "hellllo"
)

func main() {
	fmt.Println("helllo")
	var i int

	i = 30
	fmt.Println(i)
	var j int = 60
	fmt.Println(j)
	k := 70
	fmt.Println(k)
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", k, k)
	fmt.Println(p)
	var p int = 5
	//p := 6 // not posiible as p already exists
	p = 7 // possible
	fmt.Println(p)

	var t int = 77
	fmt.Printf("%v, %T\n", t, t)

	var r float64
	r = float64(r)
	fmt.Printf("%v, %T\n", r, r)

}
