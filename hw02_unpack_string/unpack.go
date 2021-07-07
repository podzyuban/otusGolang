package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var (
	ErrInvalidString = errors.New("invalid string")
	escapeSymbol     = `\`
)

func Unpack(value string) (string, error) {
	var errorResult error = nil
	var resultBuilder strings.Builder
	var repeatSymbol string

	for _, valueItem := range value {
		currentSymbol := string(valueItem)

		if unicode.IsDigit(valueItem) {
			if len(repeatSymbol) == 0 {
				errorResult = ErrInvalidString
				break
			}

			if repeatSymbol == escapeSymbol {
				repeatSymbol = currentSymbol
			} else {
				repeat, errorResult := strconv.Atoi(currentSymbol)
				if errorResult != nil {
					break
				}
				tryAppend(repeatSymbol, repeat, &resultBuilder)
				repeatSymbol = ""
			}
			continue
		}

		if len(repeatSymbol) == 0 {
			repeatSymbol = currentSymbol
			continue
		}

		if repeatSymbol == escapeSymbol {
			repeatSymbol += currentSymbol
		} else {
			tryAppend(repeatSymbol, 1, &resultBuilder)
			repeatSymbol = currentSymbol
		}
	}

	if errorResult == nil && len(repeatSymbol) != 0 {
		tryAppend(repeatSymbol, 1, &resultBuilder)
	}

	return resultBuilder.String(), errorResult
}

func tryAppend(value string, count int, resultBuilder *strings.Builder) (bool, error) {
	if count < 0 {
		return false, ErrInvalidString
	}
	if count == 0 {
		return false, nil
	}
	escapedValue := processEscaped(value)
	appendedValue := strings.Repeat(escapedValue, count)
	resultBuilder.WriteString(appendedValue)

	return true, nil
}

func processEscaped(value string) string {
	if len(value) == 0 {
		return value
	}

	for _, item := range value {
		if string(item) != escapeSymbol {
			return value
		}
	}
	return `\`
}
