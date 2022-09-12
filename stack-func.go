package main

import "fmt"

type Stack []string

const maxSize int = 10
var top int = 0

func (s *Stack) IsFull() bool {
	return top == maxSize
}

func (s *Stack) IsEmpty() bool {
	return top == 0
}

func (s *Stack) Push(str string) {
	if s.IsFull() {
		fmt.Println("Stack is full")
		return
	}

	*s = append(*s, str)
	top++
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		fmt.Println("Nothing in list")
		return
	}

	top--
	fmt.Println("Pops", (*s)[top])
}

func (s *Stack) Display() {
	if top == 0 {
		fmt.Println("Nothing to display")
		return
	}

	fmt.Println("Displaying...")
	for i := 0; i < top; i++ {
		fmt.Println((*s)[i])
	}
}

func (s *Stack) Top() {
	if s.IsEmpty() {
		fmt.Println("Nothing in list")
		return
	}
	
	fmt.Println("Top is", (*s)[top-1])
}

func (s *Stack) Size() {
	fmt.Println("size of Stack is", top)
}