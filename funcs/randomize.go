package funcs

import (
	"math/rand"
)

func Randomize(Min int, Max int) int {
	min := Min
	max := Max
	random := rand.Intn(max-min) + min

	return random
}
