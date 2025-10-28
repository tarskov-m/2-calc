package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var operations =  map[string]func(){
	"AVG": avg,
	"SUM": sum,
	"MED": med,
}

func main() {
	// Принимает операцию (AVG - среднее, SUM - сумму, MED - медиану)
	op := readOperation()
	if opFunc, ok := operations[op]; ok {
		opFunc()
	} else {
		fmt.Println("Неизвестная операция")
	}
}

func readOperation() string {
	var op string
	for {
		// считывает операцию
		fmt.Print("Введите операцию (AVG, SUM, MED): ")
		fmt.Scanln(&op)
		if op == "AVG" || op == "SUM" || op == "MED" {
			return op
		}
		fmt.Println("Неверная операция")
	}
}

func avg() {
	numbers := numberInput()
	var sum float64
	for _, number := range numbers {
		sum += number
	}

	fmt.Printf("Среднее арифметическое: %.2f\n", sum/float64(len(numbers)))
}

func sum() {
	numbers := numberInput()
	var sum float64
	for _, number := range numbers {
		sum += number
	}

	fmt.Printf("Общая сумма: %.2f\n", sum)
}

func med() {
	numbers := numberInput()

	sort.Float64s(numbers)

	n := len(numbers)
	if n == 0 {
		fmt.Println("Медиана: нет данных")
		return
	}

	var median float64
	if n%2 == 0 {
		// Для четного количества: среднее двух центральных элементов
		median = (numbers[n/2-1] + numbers[n/2]) / 2
	} else {
		// Для нечетного количества: центральный элемент
		median = numbers[n/2]
	}

	fmt.Printf("Медиана: %.2f\n", median)
}

func numberInput() []float64 {
	// принимает неограниченное число чисел через запятую
	for {
		var numbers string
		fmt.Print("Введите числа через запятую: ")
		fmt.Scanln(&numbers)
		var result []float64
		for _, number := range strings.Split(numbers, ",") {
			if number != "" {
				if num, err := strconv.ParseFloat(number, 64); err == nil {
					result = append(result, num)
				}
			}
		}
		if len(result) > 0 {
			return result
		}
	}
}
