package main

import (
	"fmt"
	"image"
	_ "image/png"
	"math"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/miketmoore/go-dice/dice"
	"github.com/nicksnyder/go-i18n/i18n"
	"golang.org/x/image/colornames"
)

var translationFile = "i18n/d6/en-US.all.json"
var lang = "en-US"

var spritePath = "assets/diceWhite.png"

func run() {
	i18n.MustLoadTranslationFile(translationFile)
	T, err := i18n.Tfunc(lang)
	if err != nil {
		panic(err)
	}

	// Load sprite sheet graphic
	pic, err := loadPicture(spritePath)
	if err != nil {
		panic(err)
	}

	diceWidth := math.Floor(pic.Bounds().W()/3) - 21
	fmt.Printf("Third of width: %v\n", diceWidth)

	halfHeight := math.Floor(pic.Bounds().H() / 2)
	fmt.Printf("Half of height: %v\n", halfHeight)

	// Build map of dice sprite sheets
	var diceSides = map[int]*pixel.Sprite{
		5: pixel.NewSprite(pic, pixel.Rect{pixel.Vec{0, 0}, pixel.Vec{diceWidth, halfHeight}}),
		3: pixel.NewSprite(pic, pixel.Rect{pixel.Vec{diceWidth, 0}, pixel.Vec{diceWidth * 2, halfHeight}}),
		4: pixel.NewSprite(pic, pixel.Rect{pixel.Vec{diceWidth * 2, 0}, pixel.Vec{diceWidth * 3, halfHeight}}),
		6: pixel.NewSprite(pic, pixel.Rect{pixel.Vec{0, halfHeight}, pixel.Vec{diceWidth, halfHeight * 2}}),
		1: pixel.NewSprite(pic, pixel.Rect{pixel.Vec{diceWidth, halfHeight}, pixel.Vec{diceWidth * 2, halfHeight * 2}}),
		2: pixel.NewSprite(pic, pixel.Rect{pixel.Vec{diceWidth * 2, halfHeight}, pixel.Vec{diceWidth * 3, halfHeight * 2}}),
	}

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

	for !win.Closed() {
		x := win.Bounds().Center().Sub(txt.Bounds().Center())
		txt.Draw(win, pixel.IM.Moved(x))
		if win.JustPressed(pixelgl.KeyEnter) || win.JustPressed(pixelgl.MouseButton1) {
			win.Clear(colornames.Darkgrey)
			rolls := dice.Roll(1, 6)
			txt.Clear()
			diceSides[rolls[0]].Draw(win, pixel.IM.Moved(win.Bounds().Center()))
		}
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}
