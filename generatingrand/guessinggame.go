package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Game: Guess a number between 0 and 10")
	fmt.Println("You have three(3) tries ")

	source := rand.NewSource(time.Now().UnixNano())

	randomizer := rand.New(source)

	secretNumber := randomizer.Intn(11)

	var guess int

	for try := 1; try <=3; try++ {
		fmt.Println("Round:", try)

		fmt.Println("Please enter your number: ")
		fmt.Scan(&guess)

		if guess < secretNumber {
			fmt.Printf("Sorry, wrong guess ; number is too small\n ")
		} else if guess > secretNumber {
			fmt.Printf("Sorry, wrong guess ; number is too large\n")
		} else {
			fmt.Print("You win!\n")
			break
		}
		
		if try == 3 {
			fmt.Printf("Game over!!\n ")
			fmt.Printf("The correct number is %d\n", secretNumber)
		}
	}
	fmt.Println("Thanks for playing, see you next time!")
}