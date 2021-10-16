package main

import (
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func (g *Game) keyGameOverMenu() {
	// Check user inputs
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		user.MenuSelect++
		p := g.audioContext.NewPlayerFromBytes(hitSound)
		p.SetVolume(0.1)
		p.Play()
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		user.MenuSelect--
		p := g.audioContext.NewPlayerFromBytes(hitSound)
		p.SetVolume(0.1)
		p.Play()
	}

	// We check that the menu is still in bounds
	if user.MenuSelect < 0 {
		user.MenuSelect = 1
	} else if user.MenuSelect > 1 {
		user.MenuSelect = 0
	}

	// Check if the user pressed enter or escape
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		user.Pause = false
		user.GameOver = false
		spawnedEnnemies = []*Ennemy{}
		g.count = 0
		user.Score = 5
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeyKPEnter) {
		switch user.MenuSelect {
		case 1:
			user.Pause = false
			user.GameOver = false
			spawnedEnnemies = []*Ennemy{}
			g.count = 0
			user.Score = 5
		case 0:
			os.Exit(0)
		}
	}
}

func (g *Game) drawGameOverMenu(screen *ebiten.Image) {
	w, h := player.Size()
	op := &ebiten.DrawImageOptions{}

	// Handle player image
	op.GeoM.Scale(0.5, 0.5)
	op.GeoM.Translate(-float64(w)/4, -float64(h)/4)
	op.GeoM.Rotate(0)
	op.GeoM.Translate(screenWidth/2-100, (screenHeight/2)-float64(user.MenuSelect*50)-10)
	screen.DrawImage(player, op)
	text.Draw(screen, "You SUCK !", text.FaceWithLineHeight(mplusBigFont, 80), screenWidth/2-100, (screenHeight/2 - 150), color.RGBA{255, 0, 0, 255})
	text.Draw(screen, "Retry", text.FaceWithLineHeight(mplusBigFont, 80), screenWidth/2-40, (screenHeight/2 - 50), color.White)
	text.Draw(screen, "Quit", text.FaceWithLineHeight(mplusBigFont, 80), screenWidth/2-40, (screenHeight / 2), color.White)
}
