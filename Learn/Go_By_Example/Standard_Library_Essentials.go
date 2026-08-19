package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	type TheUser struct {
		Name string `json:"name"`
		Age  int    `json:"age"` // this  `json:"age"` is to help show the age as the key. It can be called whatever, or hidden using "-" symbol
	}

	user := TheUser{
		Name: "Aimable",
		Age:  21,
	}
	// Go -> Marshal -> JSON
	data, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	// JSON -> Unmarshal -> Go Struct
	jsonData := []byte(`{"name": "Aimable", "age":21}`)

	var UnUser TheUser
	UnEr := json.Unmarshal(jsonData, &UnUser)
	if UnEr != nil {
		panic(UnEr)
	}

	fmt.Println(strings.Contains("Hello Aimable", "Aimable"))

	words := []string{"Go", "is", "Awesome"}
	result := strings.Join(words, " ")

	text := "I like Java"

	text = strings.Replace(text, "Java", "Go", 1)

	fmt.Println(text, result)

	message := fmt.Sprintf("Hello %s", "Aimable")

	fmt.Println(message)

	// === SORTING ===
	numbers := []int{5, 2, 8, 1, 3}
	sort.Ints(numbers)
	fmt.Println(numbers)

	type UserSor struct {
		Name string
		Age  int
	}

	usersSor := []UserSor{
		{Name: "John", Age: 30},
		{Name: "Aimable", Age: 22},
		{Name: "David", Age: 27},
	}

	sort.Slice(usersSor, func(i, j int) bool {
		return usersSor[i].Age < usersSor[j].Age
	})

	fmt.Println(usersSor)

	// === Epoch Time ===

	now := time.Now()
	fmt.Println(now.Unix(), now.UnixMilli(), now.UnixNano())

	dateString := "2026-08-18"

	date, timeErr := time.Parse("2006-01-02", dateString)
	if timeErr != nil {
		panic(err)
	}

	fmt.Println(date)

	// Number Parsing

	num, numErr := strconv.ParseInt("123", 10, 64)
	// "123" → value
	// 10    → base 10
	// 64    → int64
	fmt.Println(num, numErr)

	// URL Parsing

	u, err := url.Parse("https://example.com/users?page=2&limit=10")

	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme, " ", u.Host, " ", u.Path, " ", u.RawQuery, u.Query())

	query := u.Query()
	fmt.Println(query.Get("page"), query.Get("limit"))

	// ==== Creating your URL
	que := url.Values{}

	que.Set("page", "3")
	que.Set("limit", "8")
	fmt.Println(que.Encode())

	// === ENV Variables ===

	dbUrl := os.Getenv("DATABASE_URL")

	val, exists := os.LookupEnv("DATABSE_URL")
	if !exists {
		fmt.Println("DB_URL is missing")
	}

	fmt.Println(dbUrl, val)

	// READING & WRITTING FILES
	data, readErr := os.ReadFile("hello.txt")
	if readErr != nil {
		panic(readErr)
	}
	fmt.Println(string(data))

	newData := []byte("Hello from Go <<!!")

	newErr := os.WriteFile("hello.txt", newData, 0644)

	if newErr != nil {
		panic(newErr)
	}

	fmt.Println(string(data))

	// === File Path ===

	fipath := filepath.Join("users", "aimable", "hello.txt")
	fmt.Println(fipath)

	fmt.Println(filepath.Base("home/aimable/hello.txt"))
	fmt.Println(filepath.Dir("home/aimable/hello.txt"))
	fmt.Println(filepath.Ext("home/aimable/hello.txt"))

	// === DIRECTORIES ===
	dirErr := os.Mkdir("data", 0755)

	if dirErr != nil {
		panic(dirErr)
	}

	errfull := os.MkdirAll("data/users/profiles", 0755)

	if errfull != nil {
		panic(errfull)
	}

	entries, diErr := os.ReadDir("data")
	if diErr != nil {
		panic("diErr")
	}

	for _, entry := range entries {
		fmt.Println(entry.Name())
		if entry.IsDir() {
			fmt.Println("Directory:", entry.Name())
		} else {
			fmt.Println("File:", entry.Name())
		}
	}

}
