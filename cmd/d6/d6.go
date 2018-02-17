package main

import (
	"fmt"
	"strings"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/miketmoore/go-dice/d6"
	"github.com/miketmoore/go-dice/dice"
	"github.com/nicksnyder/go-i18n/i18n"
	"golang.org/x/image/colornames"
)

var translationFile = "i18n/d6/en-US.all.json"

func run() {
	i18n.MustLoadTranslationFile(translationFile)
	T, err := i18n.Tfunc("en-US")

	// Setup Text
	orig := pixel.V(20, 50)
	txt := text.New(orig, text.Atlas7x13)
	txt.Color = colornames.White

	// Setup GUI window
	cfg := pixelgl.WindowConfig{
		Title:  T("title"),
		Bounds: pixel.R(0, 0, 400, 225),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(txt, T("instruction"))

	// win.SetCursorVisible(false)
	for !win.Closed() {
		txt.Draw(win, pixel.IM.Moved(win.Bounds().Center().Sub(txt.Bounds().Center())))
		if win.JustPressed(pixelgl.KeyEnter) || win.JustPressed(pixelgl.MouseButton1) {
			win.Clear(colornames.Black)
			rolls := dice.Roll(1, 6)
			txt.Clear()
			fmt.Fprintln(txt, fmt.Sprintf(T("youRolledAN"), rolls[0]))
			fmt.Fprintln(txt, strings.Join(d6.Drawings[rolls[0]], "\n"))
		}
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

// func handleRequest(w http.ResponseWriter, r *http.Request) {
// 	cookieLang := r.Cookie("lang")
// 	acceptLang := r.Header.Get("Accept-Language")
// 	defaultLang = "en-US" // known valid language
// 	T, err := i18n.Tfunc(cookieLang, acceptLang, defaultLang)
// 	fmt.Println(T("Hello world"))
// }
