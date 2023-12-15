package funcs

import (
	"fmt"
	"strconv"
)

func toFloat(number string) (float64, error) {
	floatNumber, err := strconv.ParseFloat(number, 64)
	if err != nil {
		fmt.Println("Erro ao converter a string para float: ", err)
		return 0, err
	}
	return floatNumber, nil
}

func ToFloat(number string) (float64, error) {
	return toFloat(number)
}
