package uniq

import (
	"errors"
	"fmt"
	"strings"
)

var InterchangeableFlagsErr = errors.New("error: interchangeable flags")

func Usage() {
	fmt.Println("Usage: uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
}

func cutFields(str string, numFields int) string {
	words := strings.Split(str, " ")

	if numFields >= len(words) {
		return ""
	}

	return strings.Join(words[numFields:], " ")
}

func cutChars(str string, numChars int) string {
	if numChars >= len(str) {
		return ""
	}

	return str[numChars:]
}

func formatString(str string, options Options) string {
	if options.insensitive {
		str = strings.ToLower(str)
	}

	if options.field > 0 {
		str = cutFields(str, options.field)
	}

	if options.chars > 0 {
		str = cutChars(str, options.chars)
	}

	return str
}

func saveString(strs []string, str string, options Options, count int) []string {
	switch {
	case options.count:
		strs = append(strs, fmt.Sprintf("%d %s", count, str))
	case options.duplicated:
		if count > 1 {
			strs = append(strs, str)
		}

	case options.unique:
		if count == 1 {
			strs = append(strs, str)
		}
	case !options.count && !options.duplicated && !options.unique:
		strs = append(strs, str)
	}

	return strs
}

func Uniq(strs []string, options Options) ([]string, error) {
	result := []string{}

	if options.count && options.duplicated || options.duplicated && options.unique || options.count && options.unique {
		return result, fmt.Errorf("utility Uniq error: %w", InterchangeableFlagsErr)
	}

	if len(strs) == 0 {
		return result, nil
	}

	prev := strs[0]
	count := 1

	for i := 1; i < len(strs); i++ {
		prevCut := formatString(prev, options)
		curCut := formatString(strs[i], options)

		if prevCut == curCut {
			count++
		} else {
			result = saveString(result, prev, options, count)
			count = 1
			prev = strs[i]
		}
	}

	result = saveString(result, prev, options, count)

	return result, nil
}
