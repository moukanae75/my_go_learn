package main

import f "fmt"

func main(){
	a:=0
	PointOne(&a)
	f.Println(a)
}
func PointOne(a *int){
	var x int
	x = 1
	*a = x
}