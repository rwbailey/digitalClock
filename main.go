package main

import (
	"fmt"
	"time"
)

type digit [5]string

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

	dots := digit{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	blank := digit{
		"   ",
		"   ",
		"   ",
		"   ",
		"   ",
	}

	var sep digit

	for {
		clearScreen()
		moveTopLeft()

		now := time.Now()
		hour := now.Hour()
		minute := now.Minute()
		second := now.Second()

		if second%2 == 0 {
			sep = blank
		} else {
			sep = dots
		}

		clock := [8]digit{
			digits[hour/10],
			digits[hour%10],
			sep,
			digits[minute/10],
			digits[minute%10],
			sep,
			digits[second/10],
			digits[second%10],
		}

		printClock(clock)

		time.Sleep(time.Second)
	}
}

func clearScreen() {
	fmt.Print("\033[2J")
}

func moveTopLeft() {
	fmt.Print("\033[H")
}

func printClock(c [8]digit) {
	for i := 0; i < 5; i++ {
		for _, n := range c {
			fmt.Print(n[i], " ")
		}
		fmt.Println()
	}
}
