package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/miketmoore/go-dice/dice"
	"github.com/nicksnyder/go-i18n/i18n"
)

var r *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

var translationFile = "i18n/dice/en-US.all.json"
var lang = "en-US"

func main() {
	i18n.MustLoadTranslationFile(translationFile)
	T, err := i18n.Tfunc(lang)
	if err != nil {
		panic(err)
	}

	total := flag.Int("total", 1, T("flagDescription_totalRollsPerDice"))
	sides := flag.Int("sides", 6, T("flagDescription_totalSidesPerDice"))
	flag.Parse()

	fmt.Println(T("youRolledNDice", map[string]interface{}{"Total": *total}))

	rolls := dice.Roll(*total, *sides)
	for _, roll := range rolls {
		fmt.Printf("d%d: %d\n", *sides, roll)
	}
}
