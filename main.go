package main

import (
	"image/color"
	"log"

	"github.com/tyokoyama/hajimetego/typing"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
)

var (
	isEnter = false
)
// update proceeds the game state.
// update is called every frame (1/60 [s]).
func update(screen *ebiten.Image) error {
	// Write your game's logical update.
	
	if typing.IsContinue() {
		// Keyboard input
		typing.Input(ebiten.InputChars())

		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			isEnter = true
		} else {
			if isEnter {				
				typing.Check()
				isEnter = false
			}
		}
	}

    if ebiten.IsDrawingSkipped() {
        // When the game is running slowly, the rendering result
        // will not be adopted.
        return nil
    }

	// Write your game's rendering.
	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont := truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	mplusSmallFont := truetype.NewFace(tt, &truetype.Options{
		Size:	16,
		DPI:	dpi,
		Hinting: font.HintingFull,
	})

	for i, result := range typing.Results() {
		text.Draw(screen, result, mplusNormalFont, 100 + (i * 30), 20, color.White)
	}
	text.Draw(screen, typing.CurrentExample(), mplusNormalFont, 100, 100, color.White)
	text.Draw(screen, typing.UserString(), mplusNormalFont, 100, 150, color.White)
	text.Draw(screen, typing.CurrentTime(), mplusSmallFont, 100, 220, color.White)

    return nil
}

func main() {
    // Call ebiten.Run to start your game loop.
    if err := ebiten.Run(update, 320, 240, 2, "Hello World"); err != nil {
        log.Fatal(err)
    }
}