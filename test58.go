package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	input1 := make(chan string, 15)
	output1 := make(chan string, 5)
	output2 := make(chan string, 5)

	for i := 0; i < 15; i++ {
		input1 <- "bello " + strconv.Itoa(i)
	}
	close(input1)
	wg.Add(3)

	go func(op chan<- string, ip <-chan string) {
		defer wg.Done()
		for msg := range ip {
			op <- "output1 " + msg
		}
		close(op)
	}(output1, input1)

	go func(op chan<- string, ip <-chan string) {
		defer wg.Done()
		for msg := range ip {
			op <- "output2 " + msg
		}
		close(op)
	}(output2, input1)

	go func(op1 <-chan string, op2 <-chan string) {
		defer wg.Done()
		var msg string
		var ok1 bool
		var ok2 bool
		for {
			select {
			case msg, ok1 = <-op1:
				if ok1 {
					fmt.Println(msg)
				}
			case msg, ok2 = <-op2:
				if ok2 {
					fmt.Println(msg)
				}
			}
			if !ok1 && !ok2 {
				break
			}
		}
	}(output1, output2)
	wg.Wait()
}
