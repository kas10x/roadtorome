package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var romanSymbols = map[string]int{
	"C":  100,
	"XC": 90,
	"L":  50,
	"XL": 40,
	"X":  10,
	"IX": 9,
	"V":  5,
	"IV": 4,
	"I":  1,
}

func main() {

	validOperations := []string{"+", "-", "*", "/"}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите арифметическую операцию: ")

		input, _ := reader.ReadString('\n')
		trimmed := strings.TrimSpace(input)
		removedSpaces := strings.ReplaceAll(trimmed, " ", "")

		operationToPerform := ""
		for _, op := range validOperations {
			count := strings.Count(removedSpaces, op)

			if operationToPerform == "" && count == 1 {
				operationToPerform = op
			} else if count > 0 {
				panicOnInput()
			}
		}

		operands := strings.Split(removedSpaces, operationToPerform)
		_, numberType := identifyInput(operands)

		if numberType == "arabic" {
			first, _ := strconv.Atoi(operands[0])
			second, _ := strconv.Atoi(operands[1])

			if first > 10 || second > 10 {
				panicOnInput()
			}

			res := doMath([]int{first, second}, operationToPerform)
			fmt.Printf("результат: %v\n", res)
		} else if numberType == "roman" {
			first := romanToArabic(operands[0])
			second := romanToArabic(operands[1])
			if second >= first && operationToPerform == "-" || first > 10 || second > 10 {
				panicOnInput()
			}
			res := doMath([]int{first, second}, operationToPerform)
			if res < 1 {
				panicOnInput()
			}
			fmt.Printf("результат: %v\n", arabicToRoman(res))
		} else {
			panic(`что-то не так...`)
		}
	}

}

func panicOnInput() {
	panic("Некорректный ввод")
}

func arabicToRoman(inputArabic int) (outputRoman string) {

	arabicValues := make([]int, 0, len(romanSymbols))

	for _, val := range romanSymbols {
		arabicValues = append(arabicValues, val)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(arabicValues)))

	remaining := inputArabic

	for _, val := range arabicValues {
		for remaining >= val {
			remaining -= val

			for keyRoman, valueArabic := range romanSymbols {
				if valueArabic == val {
					outputRoman += keyRoman
				}
			}
		}
	}

	return outputRoman
}

func romanToArabic(inputRoman string) (outputArabic int) {

	chars := strings.Split(inputRoman, "")
	length := len(chars)
	fmt.Println(chars)

	for i := 0; i < length; i++ {
		if i+1 < len(chars) {
			checkForDouble := chars[i] + chars[i+1]
			if _, exists := romanSymbols[checkForDouble]; exists {
				outputArabic += romanSymbols[checkForDouble]
				i++
				continue
			}
		}

		symbol := chars[i]
		if _, exists := romanSymbols[symbol]; exists {
			outputArabic += romanSymbols[symbol]
		}
	}

	return outputArabic
}

func doMath(operands []int, operationToPerform string) (res int) {
	switch operationToPerform {
	case "+":
		res = operands[0] + operands[1]
	case "-":
		res = operands[0] - operands[1]
	case "*":
		res = operands[0] * operands[1]
	case "/":
		res = operands[0] / operands[1]
	}
	return res
}

func identifyInput(inputs []string) (isValid bool, numberType string) {

	if len(inputs) < 2 {
		panicOnInput()
	}

	// arabic?
	_, err1 := strconv.Atoi(inputs[0])
	_, err2 := strconv.Atoi(inputs[1])

	if err1 == nil && err2 == nil {
		return true, "arabic"
	}

	// roman?
	for i := range inputs {
		for _, c := range inputs[i] {
			_, exists := romanSymbols[string(c)]
			if !exists {
				panicOnInput()
			}
		}
	}
	return true, "roman"
}
