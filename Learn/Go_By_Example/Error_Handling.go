package main

import (
	"errors"
	"fmt"
)

func mainD() {
	errr := findUser(-1)

	if errr != nil {
		fmt.Println(errr)
	}

	// === Custom Error ===

	ageErr := ValidateAge(15)

	if ageErr != nil {
		fmt.Println(ageErr)
	}

	// === WRAPPING Errors ===
	getProfile()

	// === Error.Is ===
	UserErr := findTheUser(0)
	// THis can be problematic, "Is this error, or anything underneath this wrapped error, this particular error?"
	if UserErr == ErrUserNotFound {
		fmt.Println("User doesn't Exist")
	}

	// thats why we use error.is()
	if errors.Is(UserErr, ErrUserNotFound) {
		fmt.Println("This User doesnt Esist")
	}

	// === Sentinel Errors ===
	// These are called sentinel errors.
	// var ErrUserNotFound = errors.New("user not found")
	// var ErrPermissionDenied = errors.New("permission denied")
	// var ErrInvalidInput = errors.New("invalid input")
	balError := withdraw(100, 200)
	if errors.Is(balError, ErrInsufficientBalance) {
		fmt.Println("Your Broke Broo!!!")
	}

	// === Error.As ===
	var validationErr ValidationError

	if errors.As(ageErr, &validationErr) {
		fmt.Println("Invalid field:", validationErr.Field)
	}
}

func findUser(id int) error {
	if id <= 0 {
		return errors.New("invalid user ID")
	}

	return nil
}

// Custom Error Types
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return e.Field + ": " + e.Message
}

func ValidateAge(age int) error {
	if age < 18 {
		return ValidationError{
			Field:   "age",
			Message: "must be atleast 18",
		}
	}
	return nil
}

// Wrap this error while preserving the original error underneath.
func readUser() error {
	return errors.New("User not found")
}

func getProfile() error {
	profError := readUser()

	if profError != nil {
		return fmt.Errorf("failed to get profile: %w", profError)
	}
	return nil
}

var ErrUserNotFound = errors.New("User not found")

func findTheUser(id int) error {
	if id == 0 {
		return ErrUserNotFound
	}
	return nil
}

// ===== Sentinel Erros =====
var ErrInsufficientBalance = errors.New("insufficient balance")

func withdraw(balance, amount int) error {
	if amount > balance {
		return ErrInsufficientBalance
	}
	return nil
}
