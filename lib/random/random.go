package random

import "time"

func NewRandomString(length int) string {
	setRunes := []rune(
		"QWERTYUIOPASDFGHJKLZXCVBNM" +
			"qwertyuiopasdfghjklzxcvbnm" +
			"1234567890",
	)

	runes := make([]rune, length)
	for index, _ := range runes {
		time.Sleep(time.Millisecond)
		randomNum := time.Now().Nanosecond() % len(setRunes)
		runes[index] = setRunes[randomNum]
	}
	return string(runes)
}
