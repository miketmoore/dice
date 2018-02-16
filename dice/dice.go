package dice

import (
	"math/rand"
	"time"
)

var r *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func Roll(totalDice, sidesPerDice int) []int {
	rolls := []int{}
	for i := 0; i < totalDice; i++ {
		n := r.Intn(sidesPerDice)
		rolls = append(rolls, n+1)
	}
	return rolls
}
