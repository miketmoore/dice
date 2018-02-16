package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/miketmoore/go-dice/dice"
)

var r *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	total := flag.Int("total", 1, "total rolls per dice")
	sides := flag.Int("sides", 6, "total sides per dice")
	flag.Parse()

	fmt.Printf("You roll %d dice...\n", *total)

	rolls := dice.Roll(*total, *sides)
	for _, roll := range rolls {
		fmt.Printf("d%d: %d\n", *sides, roll)
	}
}
