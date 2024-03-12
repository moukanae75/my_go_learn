package main

import "github.com/01-edu/z01"

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	z01.PrintRune('\n')
}
func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	var res int
	for n > 0 {
		digit := n % 10
		res = res*10 + digit
		n = n / 10

	}
	for res > 0 {
		digit := res % 10
		res /= 10
		z01.PrintRune(rune(digit + '0'))

	}

}
