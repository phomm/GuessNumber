package main

import (
	"fmt"
	"math/rand"
	"time"
)

const maxTries = 3
const min = 1;
const max = 9;
var answers = map[int]string{
	-1: "Your number is less",
	0: "Your number is the same, you WON !",
	1: "Your number is more",
	maxTries: "You Lose !",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for again := true; again; again = askAgain() {
		playOnce()		
	}
}

func playOnce() {
	guess := min - 1
	tries := 0
	number := rand.Intn(max + 1 - min) + min
	fmt.Print("Guess number (", min, "..", max, "): ")		
	for tries < maxTries && guess != number {		
		if _, err := fmt.Scan(&guess); err != nil {				
			fmt.Print("No! Type a number (", min, "..", max, "): ")
			continue
		}
		tries++
		triesLeft := maxTries - tries;
		won := guess == number
		lost := !won && triesLeft == 0
		nexttry := !won && !lost
		switch {
			case won:
				printAnswer(0)
			case lost:
				fmt.Println("Number was ", number)
				printAnswer(maxTries)
			case nexttry:
				printAnswer(Sign(guess - number))
				fmt.Print("Try ", triesLeft, " more time(s): ")
		}
	}
}

func printAnswer(key int) {	
	if val, ok := answers[key]; ok {
		fmt.Println(val) 
	} else {	
		fmt.Println("bug! no key!")
	}
}

func Sign(a int) int {
	switch {
		case a < 0:
			return -1
		case a > 0:
			return +1
	}
	return 0
}

func askAgain() bool {
	fmt.Print("Again(y/n)? ")
	confirm := "";		
	fmt.Scan(&confirm)
	return confirm == "y"
}
	
