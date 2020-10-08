package main

import "fmt"

func getName() (string, string) {
	return "王", "引起"
}

func main() {
	surname, _ := getName()
	_, personalName := getName()
	fmt.Printf("My surname is %v and my personal name is %v\n", surname, personalName)
}
