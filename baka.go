package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%"
	numberSet      = "0123456789"
	allCharSet     = upperCharSet + specialCharSet + numberSet
	service        string
)

func main() {

	fmt.Println("What is this password for?")
	fmt.Scanf("%s", &service)

	rand.Seed(time.Now().Unix())
	minSpecialChar := 1
	minNum := 1
	minUpperCase := 3
	passwordLength := 7
	password := generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase)
	fmt.Println("Your password has been generated and saved.")
	fmt.Println("j0e" + password)
	filename := "baka.txt"

	f, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer f.Close()

	fmt.Fprintf(f, "\nService: "+(service)+" Password: "+("j0e"+password))

}

func generatePassword(passwordLength int, minSpecialChar, minNum, minUpperCase int) string {
	var password strings.Builder

	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})

	return string(inRune)

}

func joePass() {
	// this creates a file but places old output not good for saving :(

	file, err := os.Create("joepasswords.txt")
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString("Service: " + (service) + " Password:")

}
