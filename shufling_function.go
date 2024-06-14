package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//shufling function 3 is the len of shufling
	x := []string{"a", "b", "c", "d", "e"}

	rand.Shuffle(3, func(i, j int) {
		x[i], x[j] = x[j], x[i]
	})
	fmt.Print(x)

}
