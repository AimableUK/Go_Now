package main

import "fmt"

func main() {
	// === Hello World ===
	fmt.Println("Hello, World!")

	// variables
	var stu1 string = "John"
	var stu2 = "Peter"

	// this can't be done outside a function as it can be done for thi: var stu1 string = "John"
	x := 2

	fmt.Print(stu1, "\n")
	fmt.Print(stu2)
	fmt.Println(x)

	// === value assignment ===
	var stud string
	stud = "John"
	fmt.Println(stud)

	// === Multivariable ===
	var a, b, c = 4, "hiii", 5
	var v, y, z int = 4, 5, 6
	d, e := 5, "world!"
	var (
		n int
		m int
		o string = "hellow"
	)

	fmt.Println("multiVariable")
	fmt.Println(a, b, c, v, y, z, d, e, n, m, o)

	// camel case:
	// Each word, except the first, starts with a capital letter:
	// myVariableName = "John"

	// Pascal Case
	// Each word starts with a capital letter:
	// MyVariableName = "John"

	// Snake Case
	// Each word is separated by an underscore character:
	// my_variable_name = "John"

	// CONSTANTS
	// Untyped constant
	const PI = 3.14
	// Typed constant
	const MODEL string = "TOYOTA"
	fmt.Println(PI)

	const (
		A int = 1
		B     = 3.14
		C     = "Hi!"
	)

	// === Printing ===
	var ij string = "Hello"
	var j int = 15

	fmt.Printf("i has the value: %v and type: %T\n", ij, ij)
	fmt.Printf("j has the value: %#v and type: %T\n", j, j)

	var i = 15.5
	var txt = "Hello World!"
	var abc = "hello"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)
	fmt.Printf("%T\n", i)

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)

	fmt.Printf("%q\n", abc)

	// === ARRAY ===
	// var array_name = [length]datatype{values} here length is defined
	// var array_name = [...]datatype{values} here length is inferred
	// ====
	// array_name := [length]datatype{values} here length is defined
	// array_name := [...]datatype{values} here length is inferred

	var arr1 = [3]int{1, 2, 3}
	arr2 := [...]int{4, 5, 6, 7}

	fmt.Print("Arrays: ", arr1, "\n", arr2)
	fmt.Printf("Best Array type: %T, and values: %v", arr1, arr1)
	fmt.Println("\n", arr2[3])
	arr2[3] = 35
	fmt.Print(arr2[3], "\n")

	// Initialization
	arr3 := [5]int{}              //not initialized
	arr4 := [5]int{1, 2}          //partially initialized
	arr5 := [5]int{1, 2, 3, 4, 5} //fully initialized
	arr6 := [5]int{1: 3, 4: 6}    //Initialize Only Specific Elements
	fmt.Println(arr3, arr4, arr5, arr6, len(arr6))

	// === SLICES ===
	// slice_name := []datatype{values}
	// myslice := []int{} Initializing empty slice

	myslice1 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println("\n", "Slices: ", len(myslice1), cap(myslice1))
	fmt.Println(myslice1)

	// == SLICE FROM ARRAY ==
	// var myarray = [length]datatype{values} // An array
	// myslice := myarray[start:end] // A slice made from the array

	arr10 := [6]int{10, 11, 12, 13, 14, 15}
	slc := arr10[2:4]
	slc1 := make([]int, 5, 10) // datatype, length, capacity, no capacity means that the capacity will be equal to the length.
	fmt.Println("\n", slc, cap(slc1))

	myslice2 := make([]int, 5)
	fmt.Printf("myslice2 = %v\n", myslice2)
	fmt.Printf("length = %d\n", len(myslice2))
	fmt.Printf("capacity = %d\n", cap(myslice2))

	myslice2 = append(myslice2, 20, 21)

	myslice3 := append(slc1, myslice2...)
	fmt.Printf("two slices: %v", myslice3)

	// == Memory Efficiency ==
	// copy(dest, src)
	// The copy() function takes in two slices dest and src, and copies data from src to dest. It returns the number of elements copied.
	fmt.Println("\n", "===== Memory Efficiency =====")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))

	// === If & If Else statement ===
	// if condition {
	// 	code to be executed if condition is true
	// }

	// 1
	if 20 > 18 {
		fmt.Printf("Better")
	}

	// 2
	if 20 > 18 {
		// code to be executed if condition is true
	} else {
		// code to be executed if condition is false
	}

	time := 20
	// 3
	if time < 10 {
		fmt.Println("Good morning.")
	} else if time < 20 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}

	// === NESTED IF ===

	// Syntax
	// if condition1 {
	// 	// code to be executed if condition1 is true
	// 	if condition2 {
	// 		// code to be executed if both condition1 and condition2 are true
	// 	}
	// }

	num := 20
	if num >= 10 {
		fmt.Println("Num is more than 10.")
		if num > 15 {
			fmt.Println("Num is more than 15.")
		}
	} else {
		fmt.Println("Num is less than 10.")
	}

	// === switch Statement ===

	// Syntax
	// switch expression {
	// case x:
	//   // code block
	// case y:
	//   // code block
	// case z:
	// ...
	// default:
	//   // code block
	// }

	day := 2

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Invalid")
	}

	// === MULTI CASE
	// Syntax
	// switch expression {
	// case x,y:
	//   // code block if expression is evaluated to x or y
	// case v,w:
	//   // code block if expression is evaluated to v or w
	// case z:
	// ...
	// default:
	//   // code block if expression is not found in any cases
	// }

	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}

	// === LOOPS ===

	// ==== Go for Loop ====
	// Syntax
	// for statement1; statement2; statement3 {
	//   // code to be executed for each iteration
	// }

	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	for i := 0; i <= 100; i += 10 {
		// if i == 80 {
		// 	break
		// }
		fmt.Println(i)
	}

	// ==== Nested Loops ====
	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange", "banana"}

	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(adj[i], fruits[j])
		}
	}

	// ==== The Range Keyword ====

	// Syntax
	// for index, value := range array|slice|map {
	//   // code to be executed for each iteration
	// }

	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}

	for _, val := range fruits {
		fmt.Printf("%v\n", val)
	}

	// === Go Functions ===
	// Syntax
	// func FunctionName() {
	//  // code to be executed
	// }

	myMessage()
	myMessage()

	// ==== Function Parameters and Arguments ====
	// Syntax
	// func FunctionName(param1 type, param2 type, param3 type) {
	//  // code to be executed
	// }

}

func myMessage() {
	fmt.Println("I just got Executed")
}
