package game

import (
	"fmt"

	"github.com/inancgumus/screen"
	"teste.com/funcs"
)

func Game() {
	fmt.Println("\n-------------------------------")
	fmt.Println("	JOGO DE ADIVINHA")
	fmt.Println("-------------------------------")

	for {
		sortNum := funcs.Randomize(1, 10)
		result := funcs.Input("Qual é o numero sorteado: ")
		intResult, _ := funcs.ToInt(result)

		if intResult == sortNum {
			fmt.Println("Parabéns, você acertou!")
		} else {
			fmt.Printf("Você errou, o número sorteado é %d\n", sortNum)
		}

		playAgain := funcs.Input("Deseja jogar novamente? (s/n): ")
		if playAgain != "s" {
			screen.Clear()
			break
		}
	}

	fmt.Println("Obrigado por jogar!")
	screen.Clear()
}
