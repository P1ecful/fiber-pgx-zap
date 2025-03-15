package common

import (
	"errors"
	"strings"
)

func GetVerse(text string, verse int) (string, error) {
	verses := strings.Split(text, "\n\n")
	if verse < 1 || verse > len(verses) {
		return "", errors.New("verse out of range")
	}

	return verses[verse-1], nil
}
