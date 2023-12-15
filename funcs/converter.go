package funcs

import (
	"fmt"
)

func Converter() {
	numStr := input("Digite um número: ")

	numInt, _ := toInt(numStr)
	numFloat, _ := toFloat(numStr)

	fmt.Println("\n-------------------------------")
	fmt.Println("    Convertendo strings...")
	fmt.Println("-------------------------------")

	fmt.Printf("String: '%s'\n", numStr)
	fmt.Printf("Int:     %d\n", numInt)
	fmt.Printf("Float:   %f\n", numFloat)
	fmt.Printf("Money:   $%.2f\n", numFloat)
}
