package main

import "fmt"

func main() {
	// fmt.Println("Hello World!")
	Greeting()
	advanceGreeting("Dipansu")
	fmt.Println(add(5, 9))
	q, r := divide(45, 9)
	fmt.Printf("Quotient : %d, Remainder : %d\n", q, r)
	q1, r1 := divide(45, add(2, 4))
	fmt.Printf("Quotient : %d, Remainder : %d", q1, r1)

}

// func Greeting() {
// 	fmt.Println("HEY THERE!")
// }

// func advanceGreeting(name string) {
// 	fmt.Printf(" Hello %s\n\n", name)
// }
