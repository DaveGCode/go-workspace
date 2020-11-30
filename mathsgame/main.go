/*
Programming (C) - Assignment 1 - https://github.com/DaveGCode/DIT-College-work/blob/master/Maths%20Game%20-%20C%20Programming

--- Go version ---

Mathematics Quiz Game

- Menu Driven
- Enter the number Questions to be asked
- Start quiz
- Display the number of Questions answered correctly/incorrectly
- Exit Program gracefully

*/

package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	// Var declaration
	const numQs, ansCorrect, menuOption, correctAns, totalCorrectAns = 0, 0, 0, 0, 0
	clrscr := exec.Command("cmd", "/c", "cls")
	clrscr.Stdout = os.Stdout

	for true {
		for true {
			// Loops program, returms to menu at end
			fmt.Println("\n============================================================")
			fmt.Println("\t\t   |Mathematics Quiz!|")
			fmt.Println("\t\t..::|| Main Menu ||::..")
			fmt.Println("\n============================================================")

			fmt.Println("\t\tPlease chose an option below:")

			fmt.Println("============================================================")
			fmt.Println("1. How many questions to be asked this round?(Max 5)")
			fmt.Println("2. STart Quiz")
			fmt.Println("3. Question Answered correctly")
			fmt.Println("4. Exit Game")
			fmt.Println("============================================================")

			fmt.Scanf("%d", menuOption)
			if menuOption == 1 {

				clrscr.Run()

				fmt.Println("============================================================")
				fmt.Println("Enter the number of questions to be asked this round of the Quiz (Max 5)")
				fmt.Println("============================================================")
				fmt.Scanf("%d", numQs)

				clrscr.Run()
				fmt.Println("============================================================")
				fmt.Println("You have chosen %d questions for this round! Good Luck!", numQs)
				fmt.Println("============================================================")
				time.Sleep(2 * time.Second)
				clrscr.Run()
				break

			} else if menuOption == 2 {
				if numQs < 1 {
					fmt.Println("============================================================")
					fmt.Println("Number of questions for this round has not been chosen. Returning to main menu")
					fmt.Println("============================================================")
					break
				} else if numQs > 0 && numQs < 6 {
					clrscr.Run()
					fmt.Println("============================================================")
					fmt.Println("The Quiz will now begin, Good luck!")
					fmt.Println("============================================================")
					time.Sleep(2 * time.Second)
					clrscr.Run()

					if numQs == 1 {
						fmt.Println("test1")
						// totalCorrectAns = Question1(correctAns, numQs)
					}
				}
			}

		}
	}

}
