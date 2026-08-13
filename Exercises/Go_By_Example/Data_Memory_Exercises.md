# Data & Memory - Go

## Go Fundamentals → Intermediate Checkpoint

> **Instructions:** Answer without using Google or ChatGPT. For code questions, predict the output or explain what happens. If you don't know, say "I don't know" rather than guessing.

---

### 1. Variables & Constants

What's the difference between:

```go
var age int = 20
```

and:

```go
const age = 20
```

And what happens if you try:

```go
age = 25
```

when `age` is a constant?

---

### 2. Short Variable Declaration

What is the difference between:

```go
var name string = "Aimable"
```

and:

```go
name := "Aimable"
```

When can you use `:=`?

---

### 3. Printing & Formatting

What's the difference between:

```go
fmt.Print("Hello")
fmt.Println("Hello")
fmt.Printf("Hello %s", name)
```

What is `Printf` particularly useful for?

---

### 4. Arrays

What is the difference between:

```go
var numbers [3]int
```

and:

```go
var numbers []int
```

Which one has a fixed size?

---

### 5. Slices

Given:

```go
numbers := []int{10, 20, 30, 40, 50}

x := numbers[1:4]
```

What values are inside `x`?

And is `x` an array or a slice?

---

### 6. Slice Behavior

What will this print?

```go
numbers := []int{10, 20, 30}

x := numbers[1:]
x[0] = 99

fmt.Println(numbers)
```

Explain why.

---

### 7. Maps

Given:

```go
ages := map[string]int{
    "John": 20,
    "Mary": 25,
}
```

How do you:

1. Get John's age?
2. Add `"Alex": 30`?
3. Delete John?

---

### 8. Map Lookup

What is the difference between:

```go
age := ages["Bob"]
```

and:

```go
age, exists := ages["Bob"]
```

Why is the second form useful?

---

### 9. Control Flow

What will this print?

```go
age := 18

if age >= 18 {
    fmt.Println("Adult")
} else {
    fmt.Println("Minor")
}
```

Then explain the basic purpose of `if/else`.

---

### 10. Switch

What will this print?

```go
day := 2

switch day {
case 1:
    fmt.Println("Monday")
case 2:
    fmt.Println("Tuesday")
case 3:
    fmt.Println("Wednesday")
default:
    fmt.Println("Unknown")
}
```

Does Go automatically continue into the next `case` after finding a match?

---

### 11. Loops

What does this print?

```go
for i := 0; i < 3; i++ {
    fmt.Println(i)
}
```

And what are the three parts:

```go
for initialization; condition; update
```

doing?

---

### 12. Go's `for`

Does Go have a separate `while` keyword?

If not, how would you create a loop equivalent to:

```text
while x < 10
```

?

---

### 13. `range`

What does this do?

```go
names := []string{"John", "Mary", "Alex"}

for index, name := range names {
    fmt.Println(index, name)
}
```

What are `index` and `name`?

---

### 14. Functions

What is the difference between these two functions?

```go
func add(a int, b int) {
    fmt.Println(a + b)
}
```

and:

```go
func add(a int, b int) int {
    return a + b
}
```

---

### 15. Multiple Returns

What is special about this Go function?

```go
func divide(a, b float64) (float64, error) {
    // ...
}
```

Why might Go functions return multiple values?

---

### 16. Recursion

What is recursion?

Consider:

```go
func countDown(n int) {
    if n == 0 {
        return
    }

    fmt.Println(n)
    countDown(n - 1)
}
```

What happens when:

```go
countDown(3)
```

is called?

---

### 17. Recursion - Important

Why is this dangerous?

```go
func forever(n int) {
    fmt.Println(n)
    forever(n)
}
```

What is missing?

---

### 18. Structs

Given:

```go
type User struct {
    Name string
    Age  int
}
```

Create a `User` whose:

```text
Name = "Aimable"
Age = 20
```

Then explain what a struct is conceptually.

---

### 19. Struct Embedding

Given:

```go
type Person struct {
    Name string
}

type Employee struct {
    Person
    Salary int
}
```

If:

```go
e := Employee{
    Person: Person{Name: "Aimable"},
    Salary: 500000,
}
```

What does this print?

```go
fmt.Println(e.Name)
```

Why does `e.Name` work even though `Name` isn't directly declared inside `Employee`?

---

### 20. Pointers

What does `&` mean here?

```go
age := 20
p := &age
```

And what does `*p` mean?

---

### 21. Pointer Reasoning

What will this print?

```go
func change(x *int) {
    *x = 100
}

func main() {
    number := 10

    change(&number)

    fmt.Println(number)
}
```

Explain the entire process.

---

### 22. Methods

Given:

```go
type User struct {
    Name string
}
```

What's the difference between:

```go
func (u User) Greet() {
    fmt.Println(u.Name)
}
```

and:

```go
func (u *User) Greet() {
    fmt.Println(u.Name)
}
```

What is `u` called?

---

### 23. Pointer Receiver

What will happen here?

```go
type User struct {
    Name string
}

func (u *User) ChangeName() {
    u.Name = "John"
}

func main() {
    user := User{Name: "Aimable"}

    user.ChangeName()

    fmt.Println(user.Name)
}
```

Why?

---

### 24. Enums

Go doesn't have an `enum` keyword.

How can you create an enum-like type using `iota`?

For example, create:

```text
Pending
Approved
Rejected
```

Then tell me what values they will normally receive.

---

### 25. Interfaces + Generics - Final Challenge

Consider:

```go
type Speaker interface {
    Speak()
}

type Dog struct{}

func (d Dog) Speak() {
    fmt.Println("Woof")
}

type Cat struct{}

func (c Cat) Speak() {
    fmt.Println("Meow")
}
```

#### A.

Does `Dog` explicitly declare:

```go
implements Speaker
```

?

Why can it still be used as a `Speaker`?

#### B.

Would this work?

```go
func makeSpeak(s Speaker) {
    s.Speak()
}

makeSpeak(Dog{})
makeSpeak(Cat{})
```

Explain why.

#### C.

Now consider:

```go
func First[T any](items []T) T {
    return items[0]
}
```

What does `T` represent?

And why can this function work with both:

```go
First([]int{1, 2, 3})
```

and:

```go
First([]string{"Go", "Python"})
```

?
