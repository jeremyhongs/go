package main

import "fmt"

func move(n, from, to, via int) {
	if n <= 0 {
		return
	}

	move(n-1, from, via, to)
	fmt.Println(from, "->", to)

	move(n-1, via, to, from)
}

func hanoi(n int) {
	fmt.Println("Number of disk : ", n)
	move(n, 1, 2, 3)
}

func main() {
	hanoi(3)
	fmt.Println("JEREMY")
}
