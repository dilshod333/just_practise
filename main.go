
package main

import (
	"fmt"
	"strings"
)

type User struct {
	Name  string
	Age   int
	Email string
}

func (user *User) ValidateUser() []error {
	var errors []error

	if len(user.Name) < 6 {
		errors = append(errors, fmt.Errorf("Name: length cannot be less than 6 characters"))
	}

	if user.Age < 0 || user.Age > 120 {
		errors = append(errors, fmt.Errorf("Age: should be between 0 and 120"))
	}

	if user.Email == "" {
		errors = append(errors, fmt.Errorf("Email: cannot be empty"))
	} else {
		atIndex := strings.Index(user.Email, "@")
		index := strings.LastIndex(user.Email, ".")
		if atIndex == -1 || index == -1 || index < atIndex {
			errors = append(errors, fmt.Errorf("Email: Wrong format"))
		}
	}

	return errors
}

func main() {
	var name string
	var age int
	var email string

	fmt.Print("Enter Name:")
	fmt.Scanln(&name)

	fmt.Print("Enter Age:")
	fmt.Scanln(&age)

	fmt.Print("Enter Email:")
	fmt.Scanln(&email)

	user := User{Name: name, Age: age, Email: email}
	errors := user.ValidateUser()

	if len(errors) > 0 {
		fmt.Println("\nErrors:")
		for _, err := range errors {
			fmt.Println(err)
		}
	} else {
		fmt.Println("successful.")
	}
}
