package main

import (
	"fmt"
	"os"
	"strings"
)
func main(){
	//arg := os.Args[1:]
	read,_ := os.ReadFile("standard.txt")
	file := strings.Split(string(read),"\n\n")
	in := os.Args[1]
	for _,c :=range in {
		i := int(c - 32)
		if i>=0 && i<=94{
			fmt.Print(file[i])
		}
	}
	fmt.Println()


}
