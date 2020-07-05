package main

import "fmt"

func main() {
	c := generate()
	receive(c)

	fmt.Println("Exiting")
}

func generate() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
