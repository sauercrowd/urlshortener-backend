package persistence

import (
	"bytes"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

//generateID generates a random string containing A-Za-z
func generateID(size int) (string, error) {
	var buffer bytes.Buffer
	for i := 0; i < size; i++ {
		c := getLetter(r.Intn(52))
		_, err := buffer.WriteRune(c)
		if err != nil {
			return "", err
		}
	}
	return buffer.String(), nil
}

func getLetter(r int) rune {
	if r > 25 {
		return rune(r - 26 + 97)
	}
	return rune(r + 65)
}
