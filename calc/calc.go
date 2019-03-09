package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type myStack struct {
	data []float64
}

func (stack *myStack) push(value float64) {
	stack.data = append(stack.data, value)
}

func (stack *myStack) pop() string {
	if len(stack.data) > 0 {
		value := stack.data[len(stack.data)-1]
		stack.data = append(stack.data[:len(stack.data)-1], stack.data[len(stack.data):]...)
		return fmt.Sprintf("%g", value)
	}
	return "[ERROR]"
}

func Calculate(input string) string {
	parsedString := strings.Split(strings.TrimSuffix(input, "\n"), " ")

	if len(parsedString) < 3 {
		return "[ERROR]"
	}

	stack := myStack{make([]float64, 0)}

	for i := 0; i < len(parsedString); i++ {
		switch parsedString[i] {
		case "+":
			if len(stack.data) > 1 {
				firstValue, _ := strconv.ParseFloat(stack.pop(), 64)
				secondValue, _ := strconv.ParseFloat(stack.pop(), 64)
				stack.push(firstValue + secondValue)
			} else {
				return "[ERROR]"
			}
		case "-":
			if len(stack.data) > 1 {
				firstValue, _ := strconv.ParseFloat(stack.pop(), 64)
				secondValue, _ := strconv.ParseFloat(stack.pop(), 64)
				stack.push(-firstValue + secondValue)
			} else {
				return "[ERROR]"
			}
		case "*":
			if len(stack.data) > 1 {
				firstValue, _ := strconv.ParseFloat(stack.pop(), 64)
				secondValue, _ := strconv.ParseFloat(stack.pop(), 64)
				stack.push(firstValue * secondValue)
			} else {
				return "[ERROR]"
			}
		case "/":
			if len(stack.data) > 1 {
				firstValue, _ := strconv.ParseFloat(stack.pop(), 64)
				secondValue, _ := strconv.ParseFloat(stack.pop(), 64)
				stack.push((1 / firstValue) * secondValue)
			} else {
				return "[ERROR]"
			}
		default:
			value, _ := strconv.ParseFloat(parsedString[i], 64)
			stack.push(value)
		}
	}
	return stack.pop()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	slice, _ := reader.ReadString('\n')
	Calculate(slice)
}
