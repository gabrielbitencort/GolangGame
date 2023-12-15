package funcs

import (
	"fmt"
	"strconv"
)

func toInt(number string) (int, error) {
	floatNumber, err := strconv.ParseFloat(number, 64)
	if err != nil {
		fmt.Println("Erro ao converter a string para float: ", err)
		return 0, err
	}
	intNumber := int(floatNumber)

	return intNumber, nil
}

func ToInt(number string) (int, error) {
	return toInt(number)
}
