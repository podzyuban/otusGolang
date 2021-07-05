package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

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

			if repeatSymbol == `\` {
				repeatSymbol = currentSymbol
			} else {
				repeat, _ := strconv.Atoi(currentSymbol)
				tryAppend(repeatSymbol, repeat, &resultBuilder)
				repeatSymbol = ""
			}
			continue
		}

		if len(repeatSymbol) == 0 {
			repeatSymbol = currentSymbol
			continue
		}

		if repeatSymbol == `\` {
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

func tryAppend(value string, count int, resultBuilder *strings.Builder) {
	if count == 0 || count == -1 {
		return
	}
	escapedValue := processExcaped(value)
	appendedValue := strings.Repeat(escapedValue, count)
	resultBuilder.WriteString(appendedValue)
}

func processExcaped(value string) string {
	if len(value) == 0 {
		return value
	}

	for _, item := range value {
		if string(item) != `\` {
			return value
		}
	}
	return `\`
}
