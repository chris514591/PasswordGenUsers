package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	for {
		fmt.Print("Enter the desired length of the password: ")
		var length int
		_, err := fmt.Scanln(&length)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			continue
		}

		password := generatePassword(length)
		fmt.Println("Generated password:", password)

		fmt.Print("What is this password for? ")
		var purpose string
		_, err = fmt.Scanln(&purpose)
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		saveToFile(password, purpose)

		fmt.Print("Do you want to generate another password? (y/n) ")
		var answer string
		_, err = fmt.Scanln(&answer)
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			continue
		}
		if strings.ToLower(answer) != "y" {
			break
		}
	}
}

func generatePassword(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	rand.Seed(time.Now().UnixNano())

	password := make([]byte, length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	return string(password)
}

func saveToFile(password string, purpose string) {
	file, err := os.OpenFile("test123.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	line := fmt.Sprintf("%s\t%s\n", password, purpose)
	if _, err := file.WriteString(line); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Password and what for saved to test123.txt")
}
