package main

import "fmt"

func main() {
	fmt.Println("Hello world program")

	foo()

	for i := 1 ; i<100 ; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("End main function")
}

func foo() {
	fmt.Println("I am in foo")
}
