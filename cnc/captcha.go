package main

import (
	"math/rand"
	"strings"
)

var (
	lowerCharSet = "abcdedfghijklmnopqrst"
	upperCharSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ?"
	allCharSet   = upperCharSet
)

func generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
	var passwordd strings.Builder

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		passwordd.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		passwordd.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(passwordd.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}
