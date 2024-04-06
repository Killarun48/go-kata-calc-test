package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Функция выполянет математические операции над арабскими числами
func calcInt(firstNumber int, secondNumber int, operator string) int {
	var result int
	switch operator {
	case "+":
		result = firstNumber + secondNumber
	case "-":
		result = firstNumber - secondNumber
	case "*":
		result = firstNumber * secondNumber
	case "/":
		result = firstNumber / secondNumber
	}
	return result
}

// Функция преобразует римские числа в арабсике
func romanToInt(roman string) int {
	var result int
	romanIntMap := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
	}
	for i := 0; i < len(roman); i++ {
		if i+2 <= len(roman) && romanIntMap[roman[i:i+2]] != 0 {
			result += romanIntMap[roman[i:i+2]]
			i++
		} else {
			result += romanIntMap[string(roman[i])]
		}
	}

	return result
}

// Функция преобразует арабские числа в римские
func intToRoman(num int) string {
	var result string
	intRomanMap := map[int]string{
		100: "C",
		90:  "XC",
		50:  "L",
		40:  "XL",
		10:  "X",
		9:   "IX",
		5:   "V",
		4:   "IV",
		1:   "I",
	}

	keys := make([]int, 0)
	for k := range intRomanMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for i := 8; i >= 0; i-- {
		if keys[i] <= num {
			result += intRomanMap[keys[i]]
			num -= keys[i]
			i = 8
		}
	}

	return result
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите пример:")
	text, _ := reader.ReadString('\n')

	// Регулярное выражаение для арабских чисел
	reFirstTemplate, _ := regexp.Compile(`^(\s*)([1-9]|10)(\s*)([-+*/])(\s*)([1-9]|10)(\s*)$`)
	matchedFirstTemplate := reFirstTemplate.MatchString(text)

	// Регулярное выражаение для римских чисел
	reSecondTemplate, _ := regexp.Compile(`^(\s*)(I|II|III|IV|V|VI|VII|VIII|IX|X)(\s*)([-+*/])(\s*)(I|II|III|IV|V|VI|VII|VIII|IX|X)(\s*)$`)
	matchedSecondTemplate := reSecondTemplate.MatchString(text)

	if matchedFirstTemplate {
		text = strings.ReplaceAll(text, " ", "")
		res := reFirstTemplate.FindStringSubmatch(text)

		// Получаем аргументы выражения
		firstNumber, _ := strconv.Atoi(res[2])
		secondNumber, _ := strconv.Atoi(res[6])
		operator := res[4]

		result := calcInt(firstNumber, secondNumber, operator)
		fmt.Println(result)

	} else if matchedSecondTemplate {
		text = strings.ReplaceAll(text, " ", "")
		res := reSecondTemplate.FindStringSubmatch(text)

		// Получаем аргументы выражения
		firstNumberRoman := res[2]
		firstNumber := romanToInt(firstNumberRoman) //Преобразуем римские в арабские
		secondNumberRoman := res[6]
		secondNumber := romanToInt(secondNumberRoman) //Преобразуем римские в арабские
		operator := res[4]

		result := calcInt(firstNumber, secondNumber, operator)

		if result > 0 {
			fmt.Println(intToRoman(result)) //Преобразуем результат из арабских в римские
		} else {
			panic("В римской системе нет отрицательных чисел.")
		}

	} else {
		panic("Некорректное выражение, повторите ввод.")
	}
}
