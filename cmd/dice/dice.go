package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var r *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	total := flag.Int("total", 1, "total rolls per dice")
	flag.Parse()
	diceSides := []int{4, 6, 8, 10, 12, 20, 100}
	for _, sides := range diceSides {
		fmt.Printf("d%d: ", sides)
		roll(*total, func() {
			n := r.Intn(sides + 1)
			fmt.Printf("%d ", n)
		})
		fmt.Println()
	}
}

func roll(total int, cb func()) {
	for i := 0; i < total; i++ {
		cb()
	}
}
