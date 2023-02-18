package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"uniq/uniq"
)

func main() {
	var options uniq.Options
	uniq.Init(&options)
	flag.Parse()

	if len(flag.Args()) > 2 {
		fmt.Println(errors.New("error: too many args"))
		uniq.Usage()
		return
	}

	input, err := getStream(os.Stdin, 0, os.Open)

	if checkError(err) {
		return
	}

	strings, err := readLines(input)

	if checkError(err) {
		return
	}

	result, err := uniq.Uniq(strings, options)

	if checkError(err) {
		uniq.Usage()
		return
	}

	output, err := getStream(os.Stdout, 1, os.Create)

	if checkError(err) {
		return
	}

	err = writeLines(output, result)

	if checkError(err) {
		return
	}
}

func checkError(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}

	return false
}

func getStream(defaultValue *os.File, numArg int, operation func(string) (*os.File, error)) (*os.File, error) {
	var err error
	stream := defaultValue

	if filename := flag.Arg(numArg); filename != "" {
		if stream, err = operation(filename); err != nil {
			return nil, err
		}

		defer stream.Close()
	}

	return stream, nil
}

func readLines(input *os.File) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func writeLines(output *os.File, strings []string) error {
	writer := bufio.NewWriter(output)

	for _, line := range strings {
		fmt.Fprintln(writer, line)
	}

	return writer.Flush()
}
