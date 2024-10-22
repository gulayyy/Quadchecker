package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	x, err1 := strconv.Atoi(args[0])
	y, err2 := strconv.Atoi(args[1])
	if err1 != nil {
		return
	}
	if err2 != nil {
		return
	}

	if x > 0 && y > 0 {
		if y == 1 && x == 1 {
			fmt.Println("o")
		}
		if x == 1 && y > 1 {
			fmt.Println("o")
			for i := 0; i < y-2; i++ {
				fmt.Println("|")
			}
			fmt.Println("o")
		}
		if y == 1 && x > 1 {
			fmt.Print("o")
			fmt.Print(strings.Repeat("-", x-2))
			fmt.Println("o")
		}
		if x > 1 && y > 1 {
			fmt.Print("o")
			fmt.Print(strings.Repeat("-", x-2))
			fmt.Println("o")
			for i := 0; i < y-2; i++ {
				fmt.Print("|")
				fmt.Print(strings.Repeat(" ", x-2))
				fmt.Println("|")
			}
			fmt.Print("o")
			fmt.Print(strings.Repeat("-", x-2))
			fmt.Println("o")
		}
	} else {
		fmt.Println(" ")
	}
}
