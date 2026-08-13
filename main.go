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

	familyName("Liam", 3)
	familyName("Jenny", 14)
	familyName("Anja", 30)

	// ==== Function Returns ====
	// Syntax
	// func FunctionName(param1 type, param2 type) type {
	//  // code to be executed
	//  return output
	// }

	fmt.Println(myFunction(1, 2))

	tot := myFunction(3, 4)
	fmt.Println(tot)

	fmt.Println(myFunc(5, "Hello"))

	va, vb := myFunc(10, "Greetings!")
	fmt.Println(va, vb)

	vc, _ := myFunc(10, "Greetings!")
	fmt.Println(vc)

	// ==== Recursion ====
	testcount(1)

	// Factorial function
	fmt.Println("factorial: ", factorial_recursion(4))

	// === STRUCT ===

	// Syntax
	// type struct_name struct {
	//  member1 datatype;
	//  member2 datatype;
	//  member3 datatype;
	//  ...
	// }

	var pers1 Person
	var pers2 Person

	// Pers1 specification
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	// Pers2 specification
	pers2.name = "Cecilie"
	pers2.age = 24
	pers2.job = "Marketing"
	pers2.salary = 4500

	// Access and print Pers1 info
	fmt.Println("Name: ", pers1.name)
	fmt.Println("Age: ", pers1.age)
	fmt.Println("Job: ", pers1.job)
	fmt.Println("Salary: ", pers1.salary)

	// Access and print Pers2 info
	fmt.Println("Name: ", pers2.name)
	fmt.Println("Age: ", pers2.age)
	fmt.Println("Job: ", pers2.job)
	fmt.Println("Salary: ", pers2.salary)

	printPerson(pers1)

	// === Go Maps ===

	// Syntax
	// var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
	// b := map[KeyType]ValueType{key1:value1, key2:value2,...}

	var ableMap = map[string]string{"brand": "ford", "model": "Mustang", "year": "1964"}
	cityMap := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", ableMap)
	fmt.Printf("b\t%v\n", cityMap)

	// ==== Usin Make to create a Map ====

	// Syntax
	// var a = make(map[KeyType]ValueType)
	// b := make(map[KeyType]ValueType)

	var ma = make(map[string]string)
	ma["brand"] = "Ford"
	ma["model"] = "Mustang"
	ma["year"] = "1964"

	// create an empty map

	var aMap map[string]float32
	fmt.Println(aMap == nil)
	fmt.Println(ableMap["brand"])

	ableMap["year"] = "1970" // Updating an element
	ableMap["color"] = "red" // Adding an element

	delete(ableMap, "year")
	fmt.Print(ableMap, "\n\n")

	// ====== Check For Specific Elements in a Map =======

	// Syntax
	// val, ok :=map_name[key]

	val1, ok1 := ableMap["brand"] // Checking for existing key and its value
	val2, ok2 := ableMap["color"] // Checking for non-existing key and its value
	val3, ok3 := ableMap["day"]   // Checking for existing key and its value
	_, ok4 := ableMap["model"]    // Only checking for existing key and not its value

	fmt.Println("Maps:")
	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)

	// maps are references to hash tables.

	fmt.Println(ableMap)
	avanaMap := ableMap
	avanaMap["color"] = "blue"
	fmt.Println(avanaMap, ableMap)

	// ===== Iterating over maps ======
	amazi := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	for k, v := range amazi {
		fmt.Printf("%v : %v, ", k, v)
	}

	var babo []string
	babo = append(babo, "one", "two", "three", "four")

	for k, v := range amazi { // no order
		fmt.Printf("%v %v", k, v)
	}

	// === POINTERS ===

	agee := 20
	var point *int = &agee

	fmt.Println(agee)   // 20
	fmt.Println(point)  // something like 0xc000...
	fmt.Println(*point) // 20

	changeAgee(&agee)
	fmt.Println(agee)

	artist := "Bex"
	changeBex(&artist)
	fmt.Println(artist)

	// === Struct Embedding ===

	type Persona struct {
		Name string
	}

	type Status struct {
		marital string
	}

	type Employee struct {
		Persona
		Status
		Salary int
	}

	employee := Employee{
		Persona: Persona{
			Name: "Aimable",
		},
		Salary: 500000,
		Status: Status{
			marital: "Single",
		},
	}
	fmt.Println(employee)
	fmt.Printf("Name: %v\t Salary: %v\t Status: %v\n", employee.Name, employee.Salary, employee.marital)

	// === ENUMS ===

	fmt.Println(Pending)  // 0
	fmt.Println(Approved) // 1
	fmt.Println(Rejected) // 2

	stat := Approved

	switch stat {
	case Pending:
		fmt.Println("Waiting")
	case Approved:
		fmt.Println("Accepted")
	case Rejected:
		fmt.Println("Denied")
	}

	// === METHODS ===

	user := User{Name: "Aimable", age: 30}
	user.Greet()

	usere := User{Name: "fik", age: 23}
	usere.Greet()

	abe := Able{class: "S4c"}
	abe.Goat()

	change := User{Name: "John"}
	change.ChangeName()
	fmt.Println(change.Name)

	dog := Dog{Name: "alain"}
	cat := Cat{}

	makeSpeak(dog)
	makeSpeak(cat)

	Max(20,30)
	Max(30.43, 22.45)
}

// ===== FUNCTIONS ======
func myMessage() {
	fmt.Println("I just got Executed")
}

func familyName(fname string, age int) {
	fmt.Println("Hello", age, "year old", fname, "Refsnes")
}

func myFunction(x int, y int) (result int) {
	result = x + y
	return result
}

func myFunc(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func testcount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return testcount(x + 1)
}

func factorial_recursion(x float64) (y float64) {
	if x > 0 {
		y = x * factorial_recursion(x-1)
	} else {
		y = 1
	}
	return
}

// ====== STRUCTS =======
type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func printPerson(pers Person) {
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)
}

// ===== PONTER =====
func changeAgee(agee *int) {
	*agee = 30
}

func changeBex(artist *string) {
	*artist = "Able"
}

// ==== ENUMS =====

type Stat int

const (
	Pending Stat = iota // iota automatically increments.
	Approved
	Rejected
)

func (s Stat) String() string {
	switch s {
	case Pending:
		return "Pending"
	case Approved:
		return "Approved"
	case Rejected:
		return "Rejected"
	}
	return "Unknown"
}

// ==== Methods ====

type User struct {
	Name string
	age  int
}

// Normal Function
func greet(user User) {
	fmt.Println("Hello", user.Name)
}

// Method
func (u User) Greet() {
	fmt.Println("Hello", u.Name, " Age: ", u.age)
}

type Able struct {
	class string
}

func (a Able) Goat() {
	fmt.Println("wtf", a.class)
}

// ==== Pointer Receiver ====

func (u *User) ChangeName() {
	u.Name = "Abel"
}

func (u *User) SetName(name string) {
	u.Name = name
}

// ==== INTERFACES ====

type Speaker interface {
	speak()
}

type Dog struct{ Name string }
type Cat struct{}

func (d Dog) speak() {
	fmt.Println("Woof", d.Name)
}

func (c Cat) speak() {
	fmt.Println("Meow")
}

func makeSpeak(s Speaker) {
	s.speak()
}

// ==== GENERICS ====

func Max[T int | float64](a T, b T) T {
	if a > b {
		return a
	}
	return b
}
