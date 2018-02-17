package main

import (
	"fmt"
	"strings"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/miketmoore/go-dice/d6"
	"github.com/miketmoore/go-dice/dice"
	"golang.org/x/image/colornames"
)

var locale = map[string]string{
	"title":       "D6",
	"instruction": "Press enter or click to roll!",
	"youRolledAN": "You rolled a %d",
}

func run() {
	// Setup Text
	orig := pixel.V(20, 50)
	txt := text.New(orig, text.Atlas7x13)
	txt.Color = colornames.White

	// Setup GUI window
	cfg := pixelgl.WindowConfig{
		Title:  locale["title"],
		Bounds: pixel.R(0, 0, 400, 225),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(txt, locale["instruction"])

	// win.SetCursorVisible(false)
	for !win.Closed() {
		txt.Draw(win, pixel.IM.Moved(win.Bounds().Center().Sub(txt.Bounds().Center())))
		if win.JustPressed(pixelgl.KeyEnter) || win.JustPressed(pixelgl.MouseButton1) {
			win.Clear(colornames.Black)
			rolls := dice.Roll(1, 6)
			txt.Clear()
			fmt.Fprintln(txt, fmt.Sprintf(locale["youRolledAN"], rolls[0]))
			fmt.Fprintln(txt, strings.Join(d6.Drawings[rolls[0]], "\n"))
		}
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
