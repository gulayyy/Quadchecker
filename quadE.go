package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func QuadE() {
	if len(os.Args) != 3 {
		z01.PrintRune(' ')
		z01.PrintRune('\n')
		return
	}

	x, errX := strconv.Atoi(os.Args[1])
	y, errY := strconv.Atoi(os.Args[2])

	if errX != nil || errY != nil || x <= 0 || y <= 0 {
		z01.PrintRune(' ')
		z01.PrintRune('\n')
		return
	}

	for i := 1; i <= y; i++ {
		if i == 1 {
			for j := 1; j <= x; j++ {
				if j == 1 {
					z01.PrintRune('A')
				} else if j == x {
					z01.PrintRune('C')
				} else {
					z01.PrintRune('B')
				}
			}
		} else if i == y {
			for j := 1; j <= x; j++ {
				if j == 1 {
					z01.PrintRune('C')
				} else if j == x {
					z01.PrintRune('A')
				} else {
					z01.PrintRune('B')
				}
			}
		} else {
			for k := 1; k <= x; k++ {
				if k == 1 || k == x {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}

func main() {
	QuadE()
}
