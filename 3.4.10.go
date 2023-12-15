package main

import (
	"fmt" // пакет используется для проверки ответа, не удаляйте его
	"os"  // пакет используется для проверки ответа, не удаляйте его
)

func main() {

	value1, value2, operation := readTask()

	switch v := value1.(type) {
	case float64:
	default:
		fmt.Printf("value=%v: %T\n", v, v)
		return
	}

	switch v := value2.(type) {
	case float64:
	default:
		fmt.Printf("value=%v: %T\n", v, v)
		return
	}

	switch op := operation.(type) {
	case string:
		switch op {
		case "+", "-", "*", "/":
			result := 0.0
			switch op {
			case "+":
				result = value1.(float64) + value2.(float64)
				fmt.Printf("%.4f\n", result)
			case "-":
				result = value1.(float64) - value2.(float64)
				fmt.Printf("%.4f\n", result)
			case "*":
				result = value1.(float64) * value2.(float64)
				fmt.Printf("%.4f\n", result)
			case "/":
				result = value1.(float64) / value2.(float64)
				fmt.Printf("%.4f\n", result)
			}

		default:
			fmt.Println("неизвестная операция")
			os.Exit(1)
		}
	default:
		fmt.Printf("value=%v: %T\n", op, op)
		return
	} // все полученные значения имеют тип пустого интерфейса
} // все полученные значения имеют тип пустого интерфейса
