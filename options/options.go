package options

import (
	"fmt"

	"teste.com/funcs"
)

func ChoosenOption() {

	fmt.Println("\n-------------------------------")
	fmt.Println("	MENU DE OPÇÕES")
	fmt.Println("-------------------------------")

	fmt.Println("[1] Conversor de valores")
	fmt.Println("[2] Jogo de Adivinha")
	fmt.Println("[3] Sair")

	for {
		option := funcs.Input("-> ")
		intOption, _ := funcs.ToInt(option)

		switch intOption {
		case 1:
			convert()
		case 2:
			game()
		}
	}
}
