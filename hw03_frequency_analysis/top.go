package hw03frequencyanalysis

import (
	"sort"
	"strings"
	"unicode"
)

func Top10(value string) []string {
	// Place your code here.
	takeResult := 10
	values := getFields(strings.ToLower(value))
	values = apply(values, func(z string) string { return strings.Trim(z, "-") })
	values = filter(values, func(z string) bool { return z != "" })
	frequanceDict := toFrequenceDict(values)
	frequances := toFrequanceValues(frequanceDict)

	order(frequances)

	resultFrequeances := take(frequances, takeResult)

	result := selectValues(resultFrequeances)

	return result
}

func take(values []*frequance, count int) []*frequance {
	actualCount := count
	if actualCount > len(values) {
		actualCount = len(values)
	}
	result := values[:actualCount]
	return result
}

func order(values []*frequance) []*frequance {
	sort.Slice(values, func(left, right int) bool {
		if values[left].count == values[right].count {
			return values[left].value < values[right].value
		}
		return values[left].count > values[right].count
	})
	return values
}

func toFrequanceValues(value map[string]*frequance) []*frequance {
	result := make([]*frequance, len(value))
	index := 0
	for _, item := range value {
		result[index] = item
		index++
	}

	return result
}

func toFrequenceDict(values []string) map[string]*frequance {
	dict := make(map[string]*frequance)
	for _, item := range values {
		f, ok := dict[item]
		if !ok {
			f := newFreaquance(item)
			dict[item] = f
		} else {
			f.increaseCount()
		}
	}
	return dict
}

func getFields(value string) []string {
	return strings.FieldsFunc(value, func(c rune) bool {
		if c == '-' {
			return false
		}

		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	})
}

func selectValues(value []*frequance) []string {
	result := make([]string, len(value))
	index := 0
	for _, item := range value {
		result[index] = item.value
		index++
	}
	return result
}

func apply(value []string, selector func(item string) string) []string {
	result := make([]string, len(value))

	for _, item := range value {
		result = append(result, selector(item))
	}
	return result
}

func filter(value []string, predicate func(string) bool) []string {
	result := make([]string, 0)

	for _, item := range value {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

type frequance struct {
	value string
	count int
}

func (f *frequance) increaseCount() {
	f.count = f.count + 1
}

func newFreaquance(word string) *frequance {
	return &frequance{
		value: word,
		count: 1,
	}
}
