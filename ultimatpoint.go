package main

import f "fmt"

func UltimatePointOne(n ****int){

	var x int
	x =1
	****n = x


}
func main() {
	a := 0
	b := &a
	c := &b
	n := &c
	UltimatePointOne(&n)
	f.Println(a)
}