package calc

import (
	"calc/queue"
	"calc/stack"
	"errors"
	pkgErr "github.com/pkg/errors"
	"math"
	"strconv"
	"unicode"
)

var IncorrectInputErr = errors.New("error: incorrect input string")
var DivisionByZero = errors.New("error: division by zero")

const (
	MUL_DIV_PRIORITY  int     = 2
	ADD_DIFF_PRIORITY int     = 1
	DEFAULT_PRIORITY  int     = 0
	EPS               float64 = 1e-07
)

func priority(symbol rune) int {
	switch symbol {
	case '+', '-':
		return ADD_DIFF_PRIORITY
	case '*', '/':
		return MUL_DIV_PRIORITY
	}

	return DEFAULT_PRIORITY
}

func toPostfix(expression string) (queue.Queue, error) {
	var st stack.Stack
	var q queue.Queue
	curNum := ""

	if len(expression) == 0 {
		return q, IncorrectInputErr
	}

	for _, char := range expression {
		if unicode.IsDigit(char) {
			curNum += string(char)
		} else {
			if len(curNum) != 0 {
				q.Push(curNum)
			}

			switch char {
			case '+', '-', '*', '/':
				if st.IsEmpty() || st.Top() == '(' || priority(char) >= priority(st.Top().(rune)) {
					st.Push(char)
				} else if priority(char) < priority(st.Top().(rune)) {
					for !st.IsEmpty() && (st.Top() != '(' || priority(char) <= priority(st.Top().(rune))) {
						q.Push(st.Top())
						st.Pop()
					}

					st.Push(char)
				}
			case '(':
				st.Push(char)
			case ')':
				for st.Top() != '(' {
					q.Push(st.Top())
					st.Pop()
				}

				st.Pop()
			default:
				return q, IncorrectInputErr
			}

			curNum = ""
		}
	}

	if len(curNum) != 0 {
		q.Push(curNum)
	}

	for !st.IsEmpty() {
		q.Push(st.Top())
		st.Pop()
	}

	return q, nil
}

func operation(first float64, op rune, second float64) (float64, error) {
	var result float64

	switch op {
	case '+':
		result = first + second
	case '-':
		result = first - second
	case '*':
		result = first * second
	case '/':
		if math.Abs(second) < EPS {
			return result, DivisionByZero
		}

		result = first / second
	default:
		return result, IncorrectInputErr
	}

	return result, nil
}

func Calculate(expression string) (float64, error) {
	result := 0.0
	q, err := toPostfix(expression)

	if err != nil {
		return result, pkgErr.Wrap(err, "postfix conversion error")
	}

	var st stack.Stack

	for !q.IsEmpty() {
		element := q.Front()
		q.Pop()

		if str, success := element.(string); success {
			var number float64

			if number, err = strconv.ParseFloat(str, 64); err != nil {
				return result, pkgErr.Wrap(err, "failed to parse float")
			}

			st.Push(number)
		} else if element == '+' || element == '-' || element == '*' || element == '/' {
			second, success := st.Top().(float64)

			if !success {
				return result, pkgErr.Wrap(IncorrectInputErr, "error of conversion to float")
			}

			st.Pop()
			first, success := st.Top().(float64)

			if !success {
				return result, pkgErr.Wrap(IncorrectInputErr, "error of conversion to float")
			}

			st.Pop()
			var tmpResult float64

			if tmpResult, err = operation(first, element.(rune), second); err != nil {
				return result, pkgErr.Wrap(err, "operation application error")
			}

			st.Push(tmpResult)
		} else {
			return result, pkgErr.Wrap(IncorrectInputErr, "error of conversion to string")
		}
	}

	result = st.Top().(float64)

	return result, nil
}
