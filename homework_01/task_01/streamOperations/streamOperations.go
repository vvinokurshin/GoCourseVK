package streamOperations

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func GetStream(defaultValue *os.File, numArg int, operation func(string) (*os.File, error)) (*os.File, error) {
	var err error
	stream := defaultValue

	if filename := flag.Arg(numArg); filename != "" {
		if stream, err = operation(filename); err != nil {
			return nil, fmt.Errorf("failed to open/create file: %w", err)
		}

		defer stream.Close()
	}

	return stream, nil
}

func ReadLines(input *os.File) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func WriteLines(output *os.File, strings []string) error {
	writer := bufio.NewWriter(output)

	for _, line := range strings {
		fmt.Fprintln(writer, line)
	}

	return writer.Flush()
}
