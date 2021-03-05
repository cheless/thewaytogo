package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		t := func(){fmt.Printf("The number is: %d\n", i)}
		t()
		fmt.Printf("t is of type %T and has value %v", t, t)
	}
}