package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Setting for scan input string
	scanner := bufio.NewScanner(os.Stdin)

	var stack Stack
	var ch int
	runProgram := true

	for {
		fmt.Println("")
		fmt.Println("*--------------------------------------------------------------------*")
		fmt.Println("1.PUSH  2.POP  3.DISPLAY  4.ISFULL  5.ISEMPTY  6.TOP  7.SIZE  8.EXIT")
		fmt.Println("*--------------------------------------------------------------------*")
		fmt.Print("Input menu : ")
		
		fmt.Scanln(&ch)

		switch ch {
			case 1:
				var str string
				fmt.Print("Input string : ")
				scanner.Scan()
				str = scanner.Text()
				stack.Push(str)
				break
			
			case 2:
				stack.Pop()
				break
			
			case 3:
				stack.Display()
				break

			case 4:
				if stack.IsFull() {
					fmt.Println("Stack is full")
				} else {
					fmt.Println("Stack is not full")
					fmt.Println("You can Push")
				}
				break

			case 5:
				if stack.IsEmpty() {
					fmt.Println("Stack is empty")
				} else {
					fmt.Println("Stack is not empty")
					fmt.Println("You can Pop")
				}
				break

			case 6:
				stack.Top()
				break

			case 7:
				stack.Size()
				break

			case 8:
				fmt.Println("Exiting menu")
				runProgram = false
				break

			default:
				fmt.Println("Invalid input")
				runProgram = false
				break
		}

		if !runProgram {
			break
		}
	}
}