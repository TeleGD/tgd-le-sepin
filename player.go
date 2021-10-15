package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func (g *Game) drawPlayer(screen *ebiten.Image) {
	w, h := player.Size()
	op := &ebiten.DrawImageOptions{}

	// Check user inputs
	for _, key := range g.keys {
		if key == ebiten.KeyArrowUp {
			playerPos = 0
			break
		} else if key == ebiten.KeyArrowLeft {
			playerPos = 3 * math.Pi / 2
			break
		} else if key == ebiten.KeyArrowRight {
			playerPos = math.Pi / 2
			break
		} else if key == ebiten.KeyArrowDown {
			playerPos = math.Pi
			break
		}
	}

	// Handle player image
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Rotate(playerPos)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	screen.DrawImage(player, op)
	text.Draw(screen, fmt.Sprintf("Score : %d", playerScore), text.FaceWithLineHeight(mplusBigFont, 80), 30, 30, color.White)
	text.Draw(screen, fmt.Sprintf("TPS : %0.2f", ebiten.CurrentTPS()), text.FaceWithLineHeight(mplusBigFont, 80), 30, 80, color.White)
}
