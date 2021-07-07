package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var (
	ErrInvalidString = errors.New("invalid string")

	errInvalidParseDigit = errors.New("invalid parse digit")
	escapeSymbol         = `\`
)

func Unpack(value string) (string, error) {
	var errorResult error
	var resultBuilder strings.Builder
	var repeatSymbol string

	for _, valueItem := range value {
		currentSymbol := string(valueItem)

		if unicode.IsDigit(valueItem) {
			repeatSymbol, errorResult = processDigit(currentSymbol, repeatSymbol, &resultBuilder)
		} else {
			repeatSymbol, errorResult = processSymbol(currentSymbol, repeatSymbol, &resultBuilder)
		}

		if errorResult != nil {
			break
		}
		continue
	}

	if errorResult == nil && len(repeatSymbol) != 0 {
		tryAppend(repeatSymbol, 1, &resultBuilder)
	}

	return resultBuilder.String(), errorResult
}

func processSymbol(currentSymbol string, repeatSymbol string, resultBuilder *strings.Builder) (string, error) {
	var errorResult error
	if len(repeatSymbol) == 0 {
		return currentSymbol, nil
	}

	if repeatSymbol == escapeSymbol {
		repeatSymbol += currentSymbol
	} else {
		errorResult = tryAppend(repeatSymbol, 1, resultBuilder)
		repeatSymbol = currentSymbol
	}
	return repeatSymbol, errorResult
}

func processDigit(currentSymbol string, repeatSymbol string, resultBuilder *strings.Builder) (string, error) {
	var errorResult error
	var repeat int

	if len(repeatSymbol) == 0 {
		return "", ErrInvalidString
	}

	if repeatSymbol == escapeSymbol {
		repeatSymbol = currentSymbol
	} else {
		repeat, errorResult = strconv.Atoi(currentSymbol)

		if errorResult == nil {
			errorResult = tryAppend(repeatSymbol, repeat, resultBuilder)
		}

		repeatSymbol = ""
	}

	return repeatSymbol, errorResult
}

func tryAppend(value string, count int, resultBuilder *strings.Builder) error {
	if count < 0 {
		return errInvalidParseDigit
	}
	if count == 0 {
		return nil
	}
	escapedValue := processEscaped(value)
	appendedValue := strings.Repeat(escapedValue, count)
	resultBuilder.WriteString(appendedValue)

	return nil
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
