package main

import (
    "strconv"
    "sort"
)

// сюда вам надо писать функции, которых не хватает, чтобы проходили тесты в gotchas_test.go

func ReturnInt() int {
    return 1
}

func ReturnFloat() float32 {
    return float32(1.1)
}

func ReturnIntArray() [3]int {
    return [3]int{1, 3, 4}
}

func ReturnIntSlice() []int {
    return []int{1, 2, 3}
}

func IntSliceToString(slice []int) string {
    result := ""
    for id := range slice {
        result += strconv.Itoa(slice[id])
    }
    return result
}

func MergeSlices(floatSlice []float32, intSlice []int32) []int {
    newIntFromFloatSlice := make([]int, len(floatSlice))
    newIntFromIntSlice := make([]int, len(intSlice))

    for id := range floatSlice {
        newIntFromFloatSlice[id] = int(floatSlice[id])
    }

    for id := range intSlice {
        newIntFromIntSlice[id] = int(intSlice[id])
    }

    resultSlice := append(newIntFromFloatSlice, newIntFromIntSlice...)
    return resultSlice
}

func GetMapValuesSortedByKey(input map[int]string) []string {
    keysSlice := make([]int, 0)

    for key := range input {
        keysSlice = append(keysSlice, key)
    }

    sort.Ints(keysSlice)

    stringSlice := make([]string, 0)

    for _, value := range keysSlice {
        stringSlice = append(stringSlice, input[value])
    }

    return stringSlice
}
