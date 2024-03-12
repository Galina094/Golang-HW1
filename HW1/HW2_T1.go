package main

// На основе шаблона напишите программу, подсчитывающую сколько раз буква встречается в предложении, а также частоту встречаемости в процентах. То есть в предложении hello world результатом будет:
// h - 1 0.1
// e - 1 0.1
// l - 3 0.33
// o - 2 0.2
// w - 1 0.1
// r - 1 0.1
// d - 1 0.1

import (
	"fmt"	
	"strings"
)

func main() {
	str := "Hello world"

	str = strings.ToLower(str)
	str = strings.ReplaceAll(str, " ", "")

	arr := strings.Split(str, "")
	fmt.Println(arr)

	letterCount := make(map[string]int)
	// letterCount2 := make(map[string]map[int]float32)

	for _, v := range arr {
		letterCount[v]++
	}

	fmt.Println(letterCount)

	for letter, value := range letterCount {
		fmt.Printf("%s, %d, %0.1f\n", letter, value, (float64(value)*0.1))
	}

}
