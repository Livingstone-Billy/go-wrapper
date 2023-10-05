package main

import "fmt"

func main() {

	queue := Queue{}

	for i := 1; i < 10; i++ {
		queue.push(i * 10)
	}

	fmt.Println(queue.data...) // prints 10 20 30 40 50 60 70 80 90
	fmt.Println(queue.top())   // 10
	queue.pop()
	fmt.Println(queue.data...) // 20 30 40 50 60 70 80 90

	//stack example
	stack := Stack{}

	for i := 1; i < 10; i++ {
		stack.push(i * 10)
	}

	fmt.Println(stack.data...) // prints 10 20 30 40 50 60 70 80 90
	fmt.Println(stack.peek())  //90
	stack.pop()
	fmt.Println(stack.data...) // 10 20 30 40 50 60 70 80
}
