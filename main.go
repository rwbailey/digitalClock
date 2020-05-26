package main

import "fmt"

type digit [5]string

var line = "█"
var sep = "░"

func main() {
	digits := [10]digit{
		digit{
			"███",
			"█ █",
			"█ █",
			"█ █",
			"███",
		},
		digit{
			"██ ",
			" █ ",
			" █ ",
			" █ ",
			"███",
		},
		digit{
			"███",
			"  █",
			"███",
			"█  ",
			"███",
		},
		digit{
			"███",
			"  █",
			"███",
			"  █",
			"███",
		},
		digit{
			"█ █",
			"█ █",
			"███",
			"  █",
			"  █",
		},
		digit{
			"███",
			"█  ",
			"███",
			"  █",
			"███",
		},
		digit{
			"███",
			"█  ",
			"███",
			"█ █",
			"███",
		},
		digit{
			"███",
			"  █",
			"  █",
			"  █",
			"  █",
		},
		digit{
			"███",
			"█ █",
			"███",
			"█ █",
			"███",
		},
		digit{
			"███",
			"█ █",
			"███",
			"  █",
			"███",
		},
	}

	sep := [5]string{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}
	_ = sep

	printDigits(digits)
}

func clearScreen() {
	fmt.Print("\033[2J")
}

func moveTopLeft() {
	fmt.Print("\033[H")
}

func printDigits(d [10]digit) {
	for i := 0; i < 5; i++ {
		for _, n := range d {
			fmt.Print(n[i], " ")
		}
		fmt.Println()
	}
}
