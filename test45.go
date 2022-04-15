package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Println("Error occured")
		panic(err)
	} else {
		fmt.Println("No error")
	}
}

func main() {
	var line string
	f, err := os.Open("test2.txt")
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line = scanner.Text()
		fmt.Println(line)
	}
}
