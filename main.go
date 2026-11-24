package main

import "fmt"

// IsHappy reports whether repeatedly replacing n by the sum of the squares
// of its digits reaches 1. Unhappy numbers fall into a cycle — detected
// with Floyd's tortoise and hare, no visited set.
func IsHappy(n int) bool {
	slow, fast := n, step(n)
	for fast != 1 && slow != fast {
		slow = step(slow)
		fast = step(step(fast))
	}
	return fast == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		d := n % 10
		sum += d * d
		n /= 10
	}
	return sum
}

func main() {
	fmt.Println(IsHappy(19), IsHappy(2))
}
