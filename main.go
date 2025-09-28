package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Принимает операцию (AVG - среднее, SUM - сумму, MED - медиану)
	switch op := readOperation(); op {
	case "AVG":
		// считает среднее
		avg()
	case "SUM":
		// считает сумму
		sum()
	case "MED":
		// считает медиану
		med()
	default:
		break
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
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	if len(numbers)%2 == 0 {
		fmt.Printf("Медиана: %.2f\n", (numbers[len(numbers)/2-1]+numbers[len(numbers)/2])/2)
	} else {
		fmt.Printf("Медиана: %.2f\n", numbers[len(numbers)/2])
	}
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
