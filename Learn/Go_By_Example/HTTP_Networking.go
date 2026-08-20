package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type TheUser struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {

	// === Go -> API

	// ======================
	// 1. Make the request
	response, fetchErr := http.Get("https://jsonplaceholder.typicode.com/users/1")
	if fetchErr != nil {
		fmt.Println(fetchErr)
		return
	}
	defer response.Body.Close()

	// 2. Read the stream ONCE into a memory variable called 'body'
	body, responseErr := io.ReadAll(response.Body)
	if responseErr != nil {
		fmt.Println(responseErr)
		return
	}

	// 3. You can print the string version as much as you want now
	// because it's stored in your RAM, not in the network stream
	fmt.Println("Raw response:", string(body))

	// 4. Parse the bytes from the 'body' variable
	var user TheUser
	decErr := json.Unmarshal(body, &user)

	if decErr != nil {
		fmt.Println("Error decoding JSON:", decErr)
		return
	}

	// Now this will work!
	fmt.Println("User Name:", user.Name)

	// === HTTP.NEWREQUEST() ===

	request, reqErr := http.NewRequest(
		http.MethodGet,
		"https://example.com/users",
		nil,
	)

	request.Header.Set(
		"Authorization",
		"Bearer abc123",
	)
	request.Header.Set(
		"Content-Type",
		"application/json",
	)

	client := &http.Client{}
	response, resErr := client.Do(request)

	if resErr != nil || reqErr != nil {
		fmt.Println("error", resErr, reqErr)
	}

	// Supposed you want to send the name and the email:
	// {
	//    "name": "Aimable",
	//    "email": "aimable@example.com"
	// }

	type Me struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	me := Me{
		Name:  "Aimble",
		Email: "aimable@example.com",
	}

	data, _ := json.Marshal(me)
	bodyy := bytes.NewBuffer(data)

	req, _ := http.NewRequest(
		http.MethodPost,
		"https://example.com/users",
		bodyy,
	)
	req.Header.Set(
		"Content-Type", "application/json",
	)

	clientt := &http.Client{
		Timeout: 10 * time.Second,
	}

	res, _ := client.Do(request)

	fmt.Println(clientt, res, '\n', '\n')
	// =================

	// === Client  -> GO ===

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/weather", getWeather)

	// Instead of this:
	http.ListenAndServe(":8080", nil)

	// You can use this:
	server := &http.Server{
		Addr:         ":8080",
		Handler:      &http.ServeMux{},
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	errr := server.ListenAndServe()
	if errr != nil {
		fmt.Println(errr)
	}
}

// === Client  -> GO ===
func helloHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	fmt.Println(w, "Hello from Go!")
}

func userHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(
			w,
			"Method not allowed",
			http.StatusMethodNotAllowed,
		)
	}

	user := map[string]string{
		"name":  "Aimable",
		"email": "aimable@example.com",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

	ctx := r.Context()
	// db.QueryContext(ctx, query)
}

func getWeather(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"https://example.com/weather",
		nil,
	)

	if err != nil {
		http.Error(w, "request error", 500)
		return
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		http.Error(w, "API error", 500)
		return
	}

	defer response.Body.Close()

	// process response...

	w.WriteHeader(http.StatusOK)
}
