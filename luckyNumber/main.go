package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
// ---------------------------------------------------------

const (
	usage = `🍀                                                   🍀
	   I want to play a game with you.
 Pick up some number and let's see how lucky you are.
🍀                                                   🍀`
	positive   = `Provide a positive, non-zero number.`
	maxTurns   = 3
	regularWin = `🎉 You won! Random nubber was`
	superWin   = `🥇 Wow! You won from the first turn! Seems like you're super lucky!
Random nubber was indeed`
	loose = `Sorry, you're not a winner...`
)

func main() {

	//rand.Seed(time.Now().UnixNano())
	rand.NewSource(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println(usage)
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("It's is not a number.")
		return
	}

	if guess <= 0 {
		fmt.Println(positive)
		return
	}

	for turn := 1; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)

		if n == guess {
			if turn == 1 {
				fmt.Println(superWin, n)
			} else {
				fmt.Println(regularWin, n)
			}
			return
		}
		/*  probably better solution due not nesting if statemnt
		if n != guess {
			continue
		}

		if turn == 1 {
			fmt.Println("superWin, n")
		} else {
			fmt.Println("regularWin, n")
		}
		return
		*/
	}
	fmt.Println(loose)
}
