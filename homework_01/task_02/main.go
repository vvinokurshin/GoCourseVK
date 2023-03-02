package main

import (
	"bufio"
	"fmt"
	"os"

	"calc/calc"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	expression := scanner.Text()

	if scanner.Err() != nil {
		fmt.Println(fmt.Errorf("failed to read expression: %w", scanner.Err()).Error())
		return
	}

	var result float64
	var err error

	if result, err = calc.Calculate(expression); err != nil {
		fmt.Println(fmt.Errorf("failed to calculate expression: %w", err).Error())
		return
	}

	fmt.Println(result)
}
