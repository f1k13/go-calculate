package main

import (
	"fmt"
	"os"
	"strings"
)

func romanToArabic(roman string) int {
	romanNumerals := map[rune]int{'I': 1, 'V': 5, 'X': 10}
	result := 0
	prevValue := 0

	for _, char := range roman {
		value := romanNumerals[char]

		if value > prevValue {
			result += value - 2*prevValue
		} else {
			result += value
		}

		prevValue = value
	}

	return result
}

func arabicToRoman(arabic int) string {
	if arabic <= 0 || arabic > 3999 {
		panic("Недопустимое число для преобразования в римскую запись!")
	}

	romanNumerals := []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := ""
	for _, numeral := range romanNumerals {
		for arabic >= numeral.value {
			roman += numeral.symbol
			arabic -= numeral.value
		}
	}

	return roman
}

func calculate(a int, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			panic("Деление на ноль!")
		}
		return a / b
	default:
		panic("Недопустимая операция!")
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		panic("Ожидается ввод трех аргументов: число операция число")
	}

	aStr := args[0]
	operator := args[1]
	bStr := args[2]

	var a, b int
	isRoman := false

	if strings.ContainsAny(aStr, "IVX") || strings.ContainsAny(bStr, "IVX") {
		isRoman = true
		a = romanToArabic(aStr)
		b = romanToArabic(bStr)
	} else {
		fmt.Sscan(aStr, &a)
		fmt.Sscan(bStr, &b)
	}

	if (a < 1 || a > 10) || (b < 1 || b > 10) {
		panic("Числа должны быть от 1 до 10 включительно!")
	}

	result := calculate(a, b, operator)
	if isRoman {
		fmt.Println("Результат:", arabicToRoman(result))
	} else {
		fmt.Println("Результат:", result)
	}
}
