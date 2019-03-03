package main

import (
    "testing"
    "reflect"
    "math"
)

func TestCalculate(t *testing.T) {
    var cases = []struct {
        expected float64
        input string
    } {
        {
            expected: 40,
            input: "23 17 +",
        },
        {
            expected: 26,
            input: "2 3 * 4 5 * +",
        },
        {
            expected: 15,
            input: "1 2 3 4 + * +",
        },
        {
            expected: 21,
            input: "1 2 + 3 4 + *",
        },
        {
            expected: 6,
            input: "8 2 5 * + 1 3 2 * + 4 - /",
        },
        {
            expected: math.Inf(1),
            input: "1 0 /",
        },
        {
            expected: math.MaxFloat32, // Тест ошибки
            input: "2 3 + *",
        },
        {
            expected: math.MaxFloat32, // Тест ошибки
            input: "1 + 2",
        },
        {
            expected: math.MaxFloat32, // Тест ошибки
            input: "123+",
        },
        {
            expected: math.MaxFloat32, // Тест ошибки
            input: "",
        },
    }
    for _, item := range cases {
        result := Calculate(item.input)
        if !reflect.DeepEqual(result, item.expected) {
            t.Error("expected", item.expected, "have", result)
        }
    }
}
