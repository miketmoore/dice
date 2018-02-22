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
	"github.com/miketmoore/dice/degrees"
	"github.com/miketmoore/dice/dice"
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
	halfHeight := math.Floor(pic.Bounds().H() / 2)

	// Build map of dice sprite sheets
	var diceSides = map[int]*pixel.Sprite{
		5: newSprite(pic, 0, 0, diceWidth, halfHeight),
		3: newSprite(pic, diceWidth, 0, diceWidth*2, halfHeight),
		4: newSprite(pic, diceWidth*2, 0, diceWidth*3, halfHeight),
		6: newSprite(pic, 0, halfHeight, diceWidth, halfHeight*2),
		1: newSprite(pic, diceWidth, halfHeight, diceWidth*2, halfHeight*2),
		2: newSprite(pic, diceWidth*2, halfHeight, diceWidth*3, halfHeight*2),
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

	win.SetSmooth(true)

	fmt.Fprintln(txt, T("instruction"))

	rollCount := 0
	// Two states
	// State 1: Waiting for input
	// State 2: Rolling
	state := "waitingForInput"
	var roll int
	animationCount := 0
	degreeIndex := len(degrees.Degrees) - 1
	diceSideIndex := 1
	for !win.Closed() {
		txt.Draw(win, pixel.IM.Moved(win.Bounds().Center().Sub(txt.Bounds().Center())))

		switch state {
		case "waitingForInput":
			if win.JustPressed(pixelgl.KeyEnter) || win.JustPressed(pixelgl.MouseButton1) {
				fmt.Printf("Input detected...\n")
				win.Clear(colornames.Darkgrey)
				txt.Clear()
				state = "rollingAnimation"
			}
		case "rollingAnimation":
			animationCount++
			win.Clear(colornames.Darkgrey)
			mat := pixel.IM
			mat = mat.Moved(win.Bounds().Center())
			mat = mat.Rotated(win.Bounds().Center(), degrees.Degrees[degreeIndex])
			diceSides[diceSideIndex].Draw(win, mat)

			degreeIndex--
			if degreeIndex == 0 {
				degreeIndex = len(degrees.Degrees) - 1
			}

			diceSideIndex++
			if diceSideIndex == 7 {
				diceSideIndex = 1
			}

			if animationCount == 20 {
				animationCount = 0
				state = "rolling"
			}
		case "rolling":
			win.Clear(colornames.Darkgrey)
			fmt.Printf("Rolling...\n")
			// Get roll
			rolls := dice.Roll(1, 6)
			roll = rolls[0]
			fmt.Printf("Roll #%d: %d\n", rollCount, roll)

			diceSide := diceSides[roll]
			mat := pixel.IM
			mat = mat.Moved(win.Bounds().Center())
			mat = mat.Rotated(win.Bounds().Center(), degrees.Degrees[degreeIndex])
			diceSide.Draw(win, mat)

			rollCount++
			state = "waitingForInput"
		}
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

func newSprite(pic pixel.Picture, xa, ya, xb, yb float64) *pixel.Sprite {
	return pixel.NewSprite(pic, pixel.Rect{pixel.Vec{xa, ya}, pixel.Vec{xb, yb}})
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
