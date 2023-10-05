package main

import (
	"errors"
)

//implementing a stack

type Stack struct {
	data []interface{}
}

func (s *Stack) push(item interface{}) {
	s.data = append(s.data, item)
}

func (s *Stack) peek() (interface{}, error) {
	if s.isEmpty() {
		return nil, errors.New("Stack is empty")
	}
	lastElemIndex := len(s.data) - 1
	item := s.data[lastElemIndex]
	return item, nil
}

func (s *Stack) pop() (interface{}, error) {
	if s.isEmpty() {
		return nil, errors.New("Stack is empty")
	}
	lastElemIndex := len(s.data) - 1
	item := s.data[lastElemIndex]
	s.data = s.data[:lastElemIndex]
	return item, nil
}

func (s *Stack) isEmpty() bool {
	if len(s.data) == 0 {
		return true
	}
	return false
}

func (s *Stack) size() int {
	return len(s.data)
}

// func main() {
// 	stack := Stack{}

// 	for i := 1; i < 10; i++ {
// 		stack.push(i * 10)
// 	}

// 	fmt.Println(stack.data...) // prints 10 20 30 40 50 60 70 80 90
// 	fmt.Println(stack.peek())  //90
// 	stack.pop()
// 	fmt.Println(stack.data...) // 10 20 30 40 50 60 70 80
// }
