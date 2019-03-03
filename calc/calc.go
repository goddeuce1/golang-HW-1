package main

import (
	"os"
	"bufio"
	"strconv"
	"strings"
        "math"
)

type myStack struct {
	data []float64
}

func (stack *myStack) push(value float64) {
	stack.data = append(stack.data, value)
}

func (stack *myStack) pop() float64 {
	value := stack.data[len(stack.data) - 1]
	stack.data = append(stack.data[:len(stack.data) - 1], stack.data[len(stack.data):]...)
	return value
}

func Calculate(input string) float64 {
    parsedString := strings.Split(strings.TrimSuffix(input, "\n"), " ")

	if len(parsedString) < 3 {
		return math.MaxFloat32
	}

	stack := myStack{ make([]float64, 0) }

	for i := 0; i < len(parsedString); i++ {
		switch parsedString[i] {
			case "+" :
                                if len(stack.data) > 1 {
                                        stack.push(stack.pop() + stack.pop())
                                } else {
        	                        return math.MaxFloat32
                                }
			case "-" :
        			if len(stack.data) > 1 {
                                        stack.push(-stack.pop() + stack.pop())
        	                } else {
                                        return math.MaxFloat32
                                }
			case "*" :
                                if len(stack.data) > 1 {
                                        stack.push(stack.pop() * stack.pop())
                                } else {
        	                        return math.MaxFloat32
                                }
			case "/" :
        			if len(stack.data) > 1 {
                                        stack.push(( 1 / stack.pop() ) * stack.pop())
        	                } else {
                                        return math.MaxFloat32
        	                }
			default :
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
