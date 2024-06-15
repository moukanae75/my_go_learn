package main

import (
	"fmt"
	"regexp"
)

func SplitLogLine(text string) []string {
	// the pattern mean star withe < and end with >
	// contain all thos caracter ~,*,=,-,
	// + mean can contain one or more than caracter of pattern
	re := regexp.MustCompile(`<[~*=\-]+>`)
	return re.Split(text, -1)

}

func main() {
	fmt.Println(SplitLogLine("section 1<*>section 2<~~~>section 3 <--->"))
}
