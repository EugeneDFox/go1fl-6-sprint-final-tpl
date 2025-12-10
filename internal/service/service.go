package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConverterToTextToMorse(s string) (string, error) {
	if s == "" {
		return "", errors.New("empty line")
	}
	isInvalidMorse := func(r rune) bool {
		switch r {
		case '.', '·', '-', '−', ' ', '/':
			return false
		default:
			return true
		}
	}
	if strings.IndexFunc(s, isInvalidMorse) != -1 {
		result := morse.ToMorse(s)
		if result == "" {
			return "", errors.New("failed to encode text to Morse")
		}
		return result, nil
	} else {
		result := morse.ToText(s)
		if result == "" {
			return "", errors.New("failed to decode Morse code")
		}
		return result, nil
	}
}
