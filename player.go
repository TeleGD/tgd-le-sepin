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
			user.Position = 0
			break
		} else if key == ebiten.KeyArrowLeft {
			user.Position = 3 * math.Pi / 2
			break
		} else if key == ebiten.KeyArrowRight {
			user.Position = math.Pi / 2
			break
		} else if key == ebiten.KeyArrowDown {
			user.Position = math.Pi
			break
		}
	}

	// Handle player image
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Rotate(user.Position)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	screen.DrawImage(player, op)
	text.Draw(screen, fmt.Sprintf("Score : %d", user.Score), text.FaceWithLineHeight(mplusBigFont, 80), 30, 30, color.White)
	text.Draw(screen, "Goal : Reach 250 !", text.FaceWithLineHeight(mplusBigFont, 80), 30, 80, color.White)
	if user.Score >= 250 || user.HitScore {
		user.HitScore = true
		text.Draw(screen, fmt.Sprintf("Flag : %s", "DUMMY"), text.FaceWithLineHeight(mplusBigFont, 80), 30, 130, color.White)
	}
}
