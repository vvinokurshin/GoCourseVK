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
		fmt.Println(scanner.Err())
	}

	var result float64
	var err error

	if result, err = calc.Calculate(expression); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
