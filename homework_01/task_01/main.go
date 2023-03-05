package main

import (
	"flag"
	"fmt"
	"os"
	"uniq/streamOperations"
	"uniq/uniq"
)

func main() {
	var options uniq.Options
	uniq.Init(&options)
	flag.Parse()

	if len(flag.Args()) > 2 {
		fmt.Println(fmt.Errorf("error: too many args").Error())
		uniq.Usage()
		return
	}

	input, err := streamOperations.GetStream(os.Stdin, 0, os.Open)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to get input stream: %w", err).Error())
		return
	}

	strings, err := streamOperations.ReadLines(input)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to readLines: %w", err).Error())
		return
	}

	result, err := uniq.Uniq(strings, options)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to Uniq: %w", err).Error())
		uniq.Usage()
		return
	}

	output, err := streamOperations.GetStream(os.Stdout, 1, os.Create)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to get output stream: %w", err).Error())
		return
	}

	err = streamOperations.WriteLines(output, result)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to write lines: %w", err).Error())
		return
	}
}
