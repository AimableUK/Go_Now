package main

import "fmt"

func main() {
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
