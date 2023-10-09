package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Привет, я калькулятор\n" +
		"Ты можешь воспользоваться простыми операторами '*', '/', '+', -\n" +
		"Можешь вводить арабские и римские целые числа\n" +
		"Но вводи значения только через пробел")
	for {
		fmt.Println("Введите данные")
		text := input()
		splitText := checkLongText(text)
		checkOperanot(splitText)
		checkVolume(splitText, roman)

	}
}

var roman = map[string]int{"C": 100, "XC": 90, "L": 50, "XL": 40, "X": 10, "IX": 9, "VIII": 8, "VII": 7, "VI": 6, "V": 5,
	"IV": 4, "III": 3, "II": 2, "I": 1}
var forConfertToRoman = [14]int{100, 90, 50, 40, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
var operator = [4]string{"+", "-", "*", "/"}

const (
	SIZE     = "ОШИБКА - строка не является математической операцией. Воодите 2 аргумента и один математический оператор"
	SCALE    = "ОШИБКА - используется одновременно разные системы счисления"
	DIV      = "ОШИБКА - в римской системе нет отрицательных чисел"
	ZERO     = "ОШИБКА - в римской системе нет нуля"
	RANGE    = "ОШИБКА - калькулятор умеет работать только с арабскими целыми числами или римскими цифрами от 1 до 10"
	ONLYROME = "ОШИБКА - если хотите воспользоваться только римской системой, то вводите два аргумента используя римскую систему"
	OPER     = "ОШИБКА - такого оператора калькулятор не знает"
)

/**
Принимает строку от пользователя
*/
func input() (st string) {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		text := sc.Text()
		return text
	}
	return st
}

/**
Проверяет размер математического выражения
*/
func checkLongText(text string) []string {
	splitSTring := strings.Split(text, " ")
	if len(splitSTring) != 3 {
		panic(SIZE)
	}
	return splitSTring
}

/**
Проевряет корректность(условия) введённых данных
*/
func checkVolume(arr []string, romNum map[string]int) {
	firstNumber, err1 := strconv.Atoi(arr[0])
	secondNumber, err2 := strconv.Atoi(arr[2])
	if err1 != nil && err2 == nil || err1 == nil && err2 != nil {
		panic(SCALE)
	}
	if err1 == nil && err2 == nil {
		if firstNumber < 1 || firstNumber > 10 || secondNumber < 1 || secondNumber > 10 {
			panic(RANGE)
		} else {
			var res int = calcArabic(arr[1], firstNumber, secondNumber)
			fmt.Printf("Ответ: %d\n", res)
		}
	}
	if err1 != nil && err2 != nil {
		var fir, sec bool
		for k, v := range romNum {
			if k == arr[0] {
				firstNumber = v
				fir = true
			}
			if k == arr[2] {
				secondNumber = v
				sec = true
			}
		}
		if !fir || !sec {
			panic(ONLYROME)
		} else {
			var res string = calcRome(arr[1], firstNumber, secondNumber)
			fmt.Printf("Ответ: %s\n", res)
		}
	}
}

/**
Проеверяет соответсвие математического оператора
*/
func checkOperanot(arr []string) {
	var flag bool
	for i, _ := range operator {
		if operator[i] == arr[1] {
			flag = true
			break
		}
	}
	if !flag {
		panic(OPER)
	}
}

/**
Математические действия с арабскими цифрами
*/
func calcArabic(operator string, firstNum int, secondNum int) (res int) {
	switch operator {
	case "+":
		return firstNum + secondNum
	case "-":
		return firstNum - secondNum
	case "*":
		return firstNum * secondNum
	case "/":
		return firstNum / secondNum
	}
	return 0
}

/**
Математические действия с римскими цифрами
*/
func calcRome(operator string, firstNum int, secondNum int) (res string) {
	var resInt int
	switch operator {
	case "+":
		resInt = firstNum + secondNum
	case "-":
		resInt = firstNum - secondNum
	case "*":
		resInt = firstNum * secondNum
	case "/":
		resInt = firstNum / secondNum
	}
	switch {
	case resInt == 0:
		panic(ZERO)
	case resInt < 0:
		panic(DIV)
	}
	//метод, который принимает число и переводит в его римскую систему исчисления
	var re string = convertNumberToRome(resInt)
	return re
}

/**
Переводит число из арабской системы исчесления в римскую
*/
func convertNumberToRome(numb int) (res string) {
	var stringNumRes string
	for numb > 0 {
		for _, elem := range forConfertToRoman {
			for i := elem; i <= numb; {
				for index, value := range roman {
					if value == elem {
						stringNumRes += index
						numb -= elem
					}
				}
			}
		}
	}
	return stringNumRes
}
