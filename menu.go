package main

import (
	"image/color"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func (g *Game) drawMenu(screen *ebiten.Image) {
	w, h := player.Size()
	op := &ebiten.DrawImageOptions{}

	// Check user inputs
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) && time.Since(lastEnnemy) > 20*time.Millisecond {
		lastEnnemy = time.Now()
		user.MenuSelect++
		if user.MenuSelect > 2 {
			user.MenuSelect = 0
		}
		p := g.audioContext.NewPlayerFromBytes(hitSound)
		p.SetVolume(0.1)
		p.Play()
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) && time.Since(lastEnnemy) > 20*time.Millisecond {
		lastEnnemy = time.Now()
		user.MenuSelect--
		if user.MenuSelect < 0 {
			user.MenuSelect = 2
		}
		p := g.audioContext.NewPlayerFromBytes(hitSound)
		p.SetVolume(0.1)
		p.Play()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeyKPEnter) {
		switch user.MenuSelect {
		case 2:
			user.Menu = false
		case 0:
			os.Exit(0)
		}
	}

	// Handle player image
	op.GeoM.Scale(0.5, 0.5)
	op.GeoM.Translate(-float64(w)/4, -float64(h)/4)
	op.GeoM.Rotate(user.Position)
	op.GeoM.Translate(screenWidth/2-100, (screenHeight/2)-float64(user.MenuSelect*50)-10)
	screen.DrawImage(player, op)
	text.Draw(screen, "Play", text.FaceWithLineHeight(mplusBigFont, 80), screenWidth/2-40, (screenHeight/2 - 100), color.White)
	text.Draw(screen, "Useless Button", text.FaceWithLineHeight(mplusBigFont, 80), screenWidth/2-40, (screenHeight/2 - 50), color.White)
	text.Draw(screen, "Quit", text.FaceWithLineHeight(mplusBigFont, 80), screenWidth/2-40, (screenHeight / 2), color.White)
}
