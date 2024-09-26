package main

import "fmt"

func main() {
	// fmt.Println("Hello World!")
	Greeting()
	advanceGreeting("Dipansu")
	fmt.Println(add(5, 9))
	q, r := divide(45, 9)
	fmt.Printf("Quotient : %d, Remainder : %d", q, r)

}

// func Greeting() {
// 	fmt.Println("HEY THERE!")
// }

// func advanceGreeting(name string) {
// 	fmt.Printf(" Hello %s\n\n", name)
// }
