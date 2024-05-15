package main

import "fmt"

func myPanic() {
	panic("Something went wrong")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main:", r)
		}
	}()

	myPanic()
}
