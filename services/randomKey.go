package services

import (
	"math/rand"
	"time"
)

func RandomAccount() (string, error) {
	randomCaracteres := "1234567890"
	randomCaracteresLength := 10

	rand.Seed(time.Now().UnixNano())
	usedChars := make(map[byte]bool)
	result := ""

	for i := 0; i < randomCaracteresLength; i++ {
		randIndex := rand.Intn(randomCaracteresLength)
		char := randomCaracteres[randIndex]

		for usedChars[char] {
			randIndex = rand.Intn(randomCaracteresLength)
			char = randomCaracteres[randIndex]
		}

		result += string(char)
		usedChars[char] = true
	}

	return result, nil
}
