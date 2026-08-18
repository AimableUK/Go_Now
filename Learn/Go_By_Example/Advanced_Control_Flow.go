package main

import (
	"fmt"
)

func mainB() {
	fmt.Println("Advanced Control Flow...")

	// === Closures with Structs ===
	fmt.Println("\n--- Struct Closures ---")

	// Create a new tracker for a specific user
	userStats := statsTracker("Alice")

	fmt.Println(userStats(10)) // Added 10, total 10
	fmt.Println(userStats(5))  // Added 5, total 15

	// Another tracker for another user
	userStats2 := statsTracker("Bob")
	fmt.Println(userStats2(20)) // Added 20, total 20

	c := counter()
	fmt.Println(c())

	// === VARIADIC === It can take as many values at you want
	sum()
	sum(5)
	sum(5, 10)
	sum(5, 10, 20, 30)

	// === Multiple Return Values ===
	q, r := divide(40, 10)
	fmt.Println(q, r)
	fmt.Println(divide(10, 2))

	// === DEFER ===
	defer fmt.Println("First")
	defer fmt.Println("Second")
	defer fmt.Println("Third")
	exaDefer()

	// === Range over Iterators ===
	for n := range count {
		fmt.Println(n)
	}

}

type Stats struct {
	Name  string
	Total int
}

func statsTracker(name string) func(int) Stats {
	// The closure captures this struct
	s := Stats{Name: name, Total: 0}

	return func(score int) Stats {
		s.Total += score
		return s
	}
}

// ====== Closures ======
// Example Integer
func counter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

// Example String
func stringAppender() func(string) string {
	current := "Able"

	return func(add string) string {
		current := current + " " + add
		return current
	}
}

// Example Slices
func itemTracker() func(string) []string {
	// the closure captures slice
	items := []string{}

	return func(newItem string) []string {
		items = append(items, newItem)
		return items
	}
}

// ===== VARIADIC ===== It behaves as a slice.

func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	fmt.Println("variadic hehehe: ", numbers, total)

	return total
}

// ===== Multiple Return Values ======
func divide(a int, b int) (int, int) {
	quo := a / b
	rem := a % b

	return quo, rem
}

// ===== DEFER ======
func exaDefer() {
	defer fmt.Println("Gooddd")

	fmt.Println("Hello")
}

// ===== PANIC =====
func exaPanic() {
	panic("Sth went Wrong!!")
}

// recover best called inside the defer function and can help us to handle the program
// continuity if the panic happens so that it can be solved.
func protect() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	panic("Sth went wrong!..")
}

// ===== Range over Iterators =====
func count(yield func(int) bool) {
	for i := 1; i <= 5; i++ {
		if !yield(i) {
			return
		}
	}
}
